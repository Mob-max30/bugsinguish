# Bugsinguish — Active Resolution Engine

> **Deconstructing Bugzilla into an AI-Native Defect Resolution Engine**  
> *CloneFest Track 2: Developer Tool Reconstruction – Bugzilla*  
> **Repository:** [https://github.com/Mob-max30/bugsinguish](https://github.com/Mob-max30/bugsinguish)  
> **Team:** Pranav (Lead) · Ujwal · Pavan  

---

## 🖼️ Live Platform Interface

Below are screenshots from our live, fully-functional application running on `http://localhost:5173`. All displayed metrics, status breakdowns, vector similarity matches, and activity logs are **dynamically calculated** from active Neon database state and test suites located in `sandbox/dummy_repo` (zero hardcoded static numbers).

![Bugsinguish 3D Resolution Pipeline Dashboard](docs/screenshots/dashboard_top.png)
*Top Section: KPI Metrics, Three.js 3D Interactive Pipeline Terrain, and AI Triage Overview Donut Chart.*

![Bugsinguish Recent Issues & Live AI Activity Feed](docs/screenshots/dashboard_bottom.png)
*Bottom Section: Dynamic Recent Issues Table and Real-time AI Activity Feed.*

---

## 💡 Executive Summary

Traditional bug trackers like Bugzilla operate as passive filing cabinets, storing text complaints and waiting weeks for human engineers to search past tickets, set up local environments, and manually fix issues.

Bugsinguish transforms bug tracking from passive record-keeping into an **Active Resolution Engine**:

1. **Instant Semantic Triage**: Converts reports into high-dimensional vector embeddings, grouping duplicate reports even when written in completely different phrasing.
2. **Autonomous Sandbox Reproduction**: Spawns short-lived, isolated Docker containers to execute test suites against the target repository branch.
3. **Agentic Root Cause Analysis & Auto-PR**: Gemini 1.5 Pro analyzes crash logs and generates unified git patch diffs.
4. **Human-in-the-Loop Review**: Developers approve and merge ready-to-use fixes with a single click.

---

## 📊 Innovation Delta vs Legacy Bugzilla

| Feature Area | Legacy Bugzilla | Bugsinguish Active Resolution Engine |
|---|---|---|
| **Core Paradigm** | Passive Record-Keeping | **Active Resolution Engine** (Autonomous sandbox & auto-PR) |
| **Interface & UX** | 1990s CGI Forms, Page Reloads | **SvelteKit 5 + Tailwind CSS v4**, Real-time SSE Log Stream |
| **Duplicate Triage** | Manual human tagging & exact match | **Semantic Vector Deduplication** via Neon Postgres (`pgvector`) |
| **Code Context** | Copy-pasted plain text stack traces | **Live Codebase Connection** via Ephemeral Docker Engine SDK |
| **Resolution Target** | Human manually marks "Resolved" | **AI drafts Pull Request**, Human developer clicks "Approve & Merge" |
| **Data Privacy** | Decades of bloat & liability | **Zero-Retention Privacy** (ephemeral containers purged post-execution) |

---

## 🛠️ System Architecture

```
                  ┌────────────────────────────────────────┐
                  │   SvelteKit 5 + Tailwind v4 Frontend   │
                  │  Real-time SSE Stream & 3D Pipeline UI  │
                  └───────────────────┬────────────────────┘
                                      │
                                      ▼
                  ┌────────────────────────────────────────┐
                  │       Go (Golang) + Chi Router API      │
                  │    Concurrent HTTP & Event Streaming   │
                  └────────┬──────────────────┬────────────┘
                           │                  │
       ┌───────────────────┴───┐          ┌───┴───────────────────┐
       │  Neon Postgres DB     │          │  Docker Engine SDK    │
       │  pgvector Embeddings  │          │  Ephemeral Sandboxes  │
       └───────────────────────┘          └───────────┬───────────┘
                                                      │
                                                      ▼
                                          ┌───────────────────────┐
                                          │ Google Gemini 1.5 Pro │
                                          │ Structured RCA & Diff │
                                          └───────────────────────┘
```

* **Frontend Layer**: SvelteKit 5, Tailwind CSS v4, Three.js 3D Pipeline visualizer, and Server-Sent Events (SSE) listener.
* **Backend Layer**: Go (Golang) with Chi router handling high-concurrency requests, vector embedding generation, and live event broadcasting.
* **Database Layer**: Neon Serverless Postgres with `pgvector` extension for relational state and 768-dimensional cosine vector similarity search.
* **Sandbox & AI Layer**: Native Docker Engine API Go SDK for container lifecycle management and Google Gen AI SDK for structured JSON root-cause analysis.

---

## ⚡ Quickstart Guide: Running & Accessing the Project

### 1. Live Production Deployment
* **Backend API (Render)**: `https://bugsinguish-api.onrender.com` (Healthcheck: `https://bugsinguish-api.onrender.com/health`)
* **Database**: Neon Serverless Postgres with `pgvector` enabled

### 2. Local Environment Setup

#### Prerequisites
* Node.js v20+ and Go v1.22+
* Docker Desktop (optional, for sandbox container execution)

#### A. Run the Frontend (SvelteKit 5 UI)
```bash
cd frontend
npm install
npm run dev
```
Open **`http://localhost:5173`** in your browser.

#### B. Run the Backend (Go Server)
1. Create `backend/.env`:
```env
GEMINI_API_KEY=your_gemini_api_key
DATABASE_URL=postgres://user:password@ep-something.neon.tech/neondb?sslmode=require
```
2. Start the server:
```bash
cd backend
go run main.go
```
The Go server automatically connects to Neon, enables `pgvector`, migrates the database schema, and starts on port `8080`.

#### C. Run via Docker Compose
```bash
docker-compose up --build
```

---

## 📄 Documentation & Team Process

For complete task breakdown, team dependency maps, and hackathon golden-path demo script, refer to [bugsinguish_master_readme.txt](bugsinguish_master_readme.txt).
