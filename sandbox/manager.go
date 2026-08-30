package sandbox

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// DockerManager encapsulates the docker client and logic
type DockerManager struct {
	cli *client.Client
}

// CrashReport is the schema for the output of the sandbox
type CrashReport struct {
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	ExitCode int    `json:"exit_code"`
	Duration string `json:"duration"`
}

// Event represents an SSE event to be published to the frontend
type Event struct {
	TicketID string      `json:"ticket_id"`
	Phase    string      `json:"phase"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	Time     string      `json:"time"`
}

// EventPublisher is the interface Pavan's backend will provide us
type EventPublisher interface {
	Publish(event Event)
}

// publishHelper is a quick way to fire an event if publisher is not nil
func publishHelper(publisher EventPublisher, ticketID, phase, message string, data interface{}) {
	if publisher != nil {
		publisher.Publish(Event{
			TicketID: ticketID,
			Phase:    phase,
			Message:  message,
			Data:     data,
			Time:     time.Now().UTC().Format(time.RFC3339),
		})
	}
}

// NewDockerManager creates a new DockerManager
func NewDockerManager() (*DockerManager, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to create docker client: %w", err)
	}
	return &DockerManager{cli: cli}, nil
}

// RunDummyRepo builds an image from Dockerfile.dummy (or pulls one),
// copies the dummy repo into the container, and runs the test,
// returning a crash report.
func (dm *DockerManager) RunDummyRepo(ctx context.Context, ticketID string, publisher EventPublisher, dummyRepoPath string) (*CrashReport, error) {
	// For MVP, we will pull python:3.9-slim and mount/copy the files
	imageName := "python:3.9-slim"

	publishHelper(publisher, ticketID, "sandbox_running", "Pulling base image "+imageName+"...", nil)
	
	// Ensure the image exists
	out, err := dm.cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err == nil {
		defer out.Close()
		io.Copy(io.Discard, out) // discard pull logs
	} else {
		// Log warning but continue, it might exist locally
		fmt.Printf("Warning: ImagePull failed: %v\n", err)
	}

	start := time.Now()

	publishHelper(publisher, ticketID, "sandbox_running", "Spawning ephemeral container...", nil)

	// 1. Create a container
	resp, err := dm.cli.ContainerCreate(ctx, &container.Config{
		Image:        imageName,
		Cmd:          []string{"python", "test_calculator.py"},
		WorkingDir:   "/app",
		Tty:          false,
	}, nil, nil, nil, "")
	if err != nil {
		return nil, fmt.Errorf("failed to create container: %w", err)
	}
	containerID := resp.ID

	// Defer cleanup of the container
	defer func() {
		// Force remove container
		removeOptions := types.ContainerRemoveOptions{
			RemoveVolumes: true,
			Force:         true,
		}
		if err := dm.cli.ContainerRemove(context.Background(), containerID, removeOptions); err != nil {
			fmt.Printf("Failed to remove container %s: %v\n", containerID, err)
		}
	}()

	// 2. Copy dummy repo files into the container
	publishHelper(publisher, ticketID, "sandbox_running", "Injecting repository context...", nil)
	tarArchive, err := createTarArchive(dummyRepoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create tar archive: %w", err)
	}
	err = dm.cli.CopyToContainer(ctx, containerID, "/app", bytes.NewReader(tarArchive), types.CopyToContainerOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to copy files to container: %w", err)
	}

	// 3. Start the container
	publishHelper(publisher, ticketID, "sandbox_running", "Executing reproduction tests...", nil)
	if err := dm.cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		return nil, fmt.Errorf("failed to start container: %w", err)
	}

	// 4. Wait for it to finish
	statusCh, errCh := dm.cli.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
	var exitCode int
	select {
	case err := <-errCh:
		if err != nil {
			return nil, fmt.Errorf("error waiting for container: %w", err)
		}
	case status := <-statusCh:
		exitCode = int(status.StatusCode)
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	// 5. Grab logs
	publishHelper(publisher, ticketID, "sandbox_running", "Collecting crash logs...", nil)
	outLogs, err := dm.cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return nil, fmt.Errorf("failed to get container logs: %w", err)
	}
	defer outLogs.Close()

	var stdout, stderr bytes.Buffer
	stdcopy.StdCopy(&stdout, &stderr, outLogs)

	return &CrashReport{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: exitCode,
		Duration: time.Since(start).String(),
	}, nil
}

// createTarArchive creates a tar archive of the given directory in memory
func createTarArchive(srcPath string) ([]byte, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	err := filepath.Walk(srcPath, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fi.Mode().IsRegular() {
			return nil
		}

		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}
		
		relPath, err := filepath.Rel(srcPath, file)
		if err != nil {
			return err
		}
		header.Name = relPath

		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	if err := tw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
