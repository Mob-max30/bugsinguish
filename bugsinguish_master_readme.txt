================================================================================
                            BUGSINGUISH — MASTER README
     Deconstructing Bugzilla into an AI-Native Defect Resolution Engine
        Track 2: Developer Tool Reconstruction – Bugzilla (CloneFest)
                REPO: https://github.com/Mob-max30/bugsinguish
                Team: Pranav (Lead) · Ujwal · Pavan
================================================================================
This is the single source of truth for the project — pitch, problem,
architecture, user journey, demo script, AND who's building what. Every
teammate should be able to open only this file and know the whole picture.
================================================================================


SECTION A — THE PITCH
================================================================================

A.1 WHAT IS BUGSINGUISH? (30-SECOND ELEVATOR PITCH)
--------------------------------------------------------------------------------
Traditional bug trackers like Bugzilla are passive "digital filing cabinets."
When something breaks, they store a record of the complaint and wait weeks for
a human engineer to manually search through past tickets, reproduce the issue,
and write a fix.

Bugsinguish transforms bug tracking from a static filing cabinet into an
"Active Resolution Engine." When a bug is reported, Bugsinguish:
  1. Instantly understands what the bug means using AI (catching duplicate
     reports even if written in completely different words).
  2. Clones the affected code into an isolated, temporary test sandbox.
  3. Diagnoses the exact root cause and drafts a ready-to-merge Pull Request
     (PR) with the fix.
  4. Leaves the human engineer in control with a single click to review and
     merge.

A.2 THE PROBLEM: WHY DID BUGZILLA FALL BEHIND?
--------------------------------------------------------------------------------
Bugzilla was revolutionary in 1998, but modern software development has
outgrown its design philosophy.

  * Pain Point 1 — Duplicate Ticket Nightmare: When a popular app crashes,
    500 users file tickets using different words. A human admin must read
    and manually tag all 500 duplicates.
  * Pain Point 2 — Disconnected from Code ("Works on My Machine"): Bugzilla
    tickets are just plain text — no branch, no environment, no way to
    reproduce without manual dev setup.
  * Pain Point 3 — The Slow Manual Bottleneck: A ticket sits in queue for
    days before a developer even opens the repo.
  * Pain Point 4 — Enterprise Privacy & Storage Bloat: Storing decades of
    raw code/crash logs creates compliance and security liabilities.
  * Pain Point 5 — Intimidating 1990s Interface: A database spreadsheet with
    40+ required fields, alienating everyone but specialized QA engineers.

A.3 HOW BUGSINGUISH SOLVES EACH PROBLEM
--------------------------------------------------------------------------------
+---------------------------+--------------------------------------------------+
| Legacy Bugzilla Problem   | How Bugsinguish Solves It                        |
+---------------------------+--------------------------------------------------+
| Exact-match search misses | Semantic Deduplication: reports become vector    |
| duplicates.               | "meanings." AI detects "App crashes on login"    |
|                           | and "Cannot sign in on iOS" are the same issue   |
|                           | and groups them automatically.                   |
+---------------------------+--------------------------------------------------+
| Bugs lack live code       | Ephemeral Docker Sandbox: spins up an isolated,  |
| context.                  | lightweight container, pulls the exact branch,   |
|                           | and executes tests.                              |
+---------------------------+--------------------------------------------------+
| Developers spend hours    | Agentic Root Cause Analysis & Auto-PR: AI reads  |
| debugging syntax/logic.   | sandbox crash logs, pinpoints the broken line,   |
|                           | and drafts a proposed code fix.                  |
+---------------------------+--------------------------------------------------+
| Enterprises fear data     | Zero-Retention Privacy: raw code/logs discarded  |
| leaks in bug logs.        | after resolution — only anonymized vectors kept. |
+---------------------------+--------------------------------------------------+
| Clunky, slow page         | Real-Time Reactive UI: SvelteKit 5 streams live  |
| reloads.                  | AI debugging logs step-by-step via SSE.          |
+---------------------------+--------------------------------------------------+

A.4 HOW WE DIFFER FROM REGULAR BUGZILLA (INNOVATION TABLE)
--------------------------------------------------------------------------------
| Feature  | Legacy Bugzilla         | Bugsinguish                         |
|----------|-------------------------|--------------------------------------|
| Paradigm | Passive Record-Keeping  | Active Resolution                    |
| UI/UX    | 1990s CGI, Page Refresh | SvelteKit, Real-Time SSE Updates     |
| Triage   | Human reads & tags      | AI vector-matches & groups instantly |
| Context  | Copy-pasted text/logs   | Live codebase connection via Sandbox |
| End Goal | Human marks "Resolved"  | AI drafts PR, Human clicks "Merge"   |


SECTION B — PRODUCT & ARCHITECTURE
================================================================================

B.1 MINIMUM VIABLE PRODUCT (MVP) FOR THE HACKATHON
--------------------------------------------------------------------------------
Do NOT build a full CI/CD pipeline. The MVP must perfectly demonstrate a
"Golden Path" using a pre-configured dummy repository (e.g., a simple
calculator app with a known bug).
  * MVP Step 1: User submits a bug with a stack trace via the UI.
  * MVP Step 2: System checks Neon (pgvector) for duplicates. If none,
    creates a new ticket.
  * MVP Step 3: UI shows real-time SSE logs: "Cloning repo...", "Spawning
    container...".
  * MVP Step 4: Backend spins up a Docker container, sends context to
    Gemini, and returns a JSON root cause analysis.
  * MVP Step 5: The UI Kanban board updates with the AI's diagnosis and a
    generated code diff (Draft PR).

B.2 SYSTEM ARCHITECTURE
--------------------------------------------------------------------------------
  * Frontend: SvelteKit 5 + Tailwind CSS v4 + Bits UI — SSE for live
    terminal-style logs, no page reloads.
  * Backend/API: Go (Golang) + Chi Router — handles concurrent triage
    requests with high performance.
  * Database: Neon Serverless Postgres + pgvector — relational ticket state
    and semantic embeddings side-by-side.
  * AI/Agent Layer: Google Gen AI SDK — structured JSON outputs for
    predictable code patching and embeddings.
  * Sandbox/Infra: Native Docker Engine API (Go SDK) — short-lived
    containers destroyed immediately after use.

B.3 THE STEP-BY-STEP USER JOURNEY (FULL DETAIL)
--------------------------------------------------------------------------------
PHASE 1 — Intake and Instant Triage (The Gateway)
  1.1  User (or automated script) submits a bug description, crash log, and
       repo branch URL.
  1.2  Go backend calls the Google Gen AI SDK to convert the text into a
       vector embedding.
  1.3  Backend queries Neon Postgres (pgvector) for similar embeddings.
  1.4  If a duplicate is found → user is alerted, logs appended to the
       existing ticket. If not → a new ticket is created with a Ticket ID.

PHASE 2 — Autonomous Reproduction (The Sandbox)
  2.1  Backend spins up a lightweight Docker container.
  2.2  Container clones the specified repo branch and sets up the
       environment.
  2.3  Container executes a test to reproduce the stack trace.
  2.4  Crash logs are extracted, and the container is immediately destroyed
       (Zero-Retention).

PHASE 3 — AI Diagnosis and Patch Generation (The RCA)
  3.1  Backend sends the bug report, code context, and crash logs to
       Google Gemini.
  3.2  Gemini outputs a strict JSON schema with a human-readable root
       cause analysis.
  3.3  Gemini generates a git patch/code diff to fix the bug.
  3.4  UI updates in real-time via SSE, showing the Draft PR forming.

PHASE 4 — Human-in-the-Loop Review (The Resolution)
  4.1  Developer reviews the AI's diagnosis and code diff on the dashboard.
  4.2  Developer clicks "Approve & Merge" — platform opens/merges the PR
       via the GitHub API.
  4.3  Ticket is marked RESOLVED, and raw code data is purged.

B.4 THE HACKATHON DEMO SCRIPT (THE GOLDEN PATH)
--------------------------------------------------------------------------------
Demo runs against a sample repo containing one intentional bug (e.g., an
unhandled zero-division error in a calculator app):
  1. Submit the bug via the frontend.
  2. Watch the live SSE logs stream:
       "[1/4] Checking semantic duplicates... None found."
       "[2/4] Spinning up ephemeral Docker container..."
       "[3/4] Reproducing failure on branch 'main'..."
       "[4/4] Root cause found. Generating patch..."
  3. Inspect the resulting Kanban card showing the root cause explanation
     and the generated git diff, ready to be merged.


SECTION C — RUBRIC / TRACK COMPLIANCE CHECKLIST
================================================================================
  [ ] Not a Bugzilla clone/fork — own UX, own architecture, own workflow model
  [ ] Reference repo (bugzilla/bugzilla) cited as "studied," not copied
  [ ] Core bug-tracking lifecycle present: report → triage → diagnose →
      resolve → close (mandatory backbone regardless of AI features on top)
  [ ] Collaboration features: multiple users, ticket ownership/assignment,
      status visibility (Kanban satisfies this)
  [ ] Reporting/analytics: at minimum, a ticket-count-by-status view —
      cheap to add, don't skip it, it's an explicit rubric category
  [ ] Performance: SSE streaming + ephemeral containers double as your
      "performance" story — make it demoable, not just claimed
  [ ] Documentation: this file + inline code comments
  [ ] Innovation delta vs Bugzilla clearly stated (Section A.4 table)


SECTION D — FILE STRUCTURE (MONOREPO)
================================================================================
bugsinguish/
├── .github/
│   └── workflows/
│       └── dev-build.yml       (Basic CI to check Go/Svelte builds)
├── frontend/                   (Pranav's Domain)
│   ├── src/
│   │   ├── lib/
│   │   │   ├── components/     (Kanban Board, Ticket Modal)
│   │   │   └── api/            (SSE client, Fetch wrappers)
│   │   ├── routes/
│   │   │   ├── +page.svelte    (Dashboard/Kanban)
│   │   │   └── submit/         (Bug Intake Form)
│   │   └── app.css             (Tailwind config)
│   ├── package.json
│   └── svelte.config.js
├── backend/                    (Pavan's Domain)
│   ├── cmd/
│   │   └── server/
│   │       └── main.go         (Entry point, Chi router setup)
│   ├── internal/
│   │   ├── api/
│   │   │   └── handlers.go     (HTTP endpoints, SSE streaming)
│   │   ├── db/
│   │   │   └── neon.go         (Postgres + pgvector queries)
│   │   ├── ai/
│   │   │   └── gemini.go       (Google SDK wrapper, prompts)
│   │   └── models/
│   │       └── ticket.go       (Go structs for JSON payloads)
│   ├── go.mod
│   └── go.sum
├── sandbox/                    (Ujwal's Domain)
│   ├── manager.go              (Go Docker SDK integration)
│   ├── Dockerfile.dummy        (Environment for the test app)
│   └── dummy_repo/             (The fake app with an intentional bug)
├── docker-compose.yml          (For local dev: runs frontend + backend)
├── README.md
└── bugsinguish_master_readme.txt   (this file)


SECTION E — GIT WORKFLOW
================================================================================
  main                    → production / final demo only
  dev                     → integration branch, merge feature branches
                             here first, test, then fast-forward to main
  pranav                  → Pranav's frontend work
  pavan                   → Pavan's backend/db work
  ujwal                   → Ujwal's sandbox/AI-diagnosis work

Rule: nobody merges their own feature branch straight to main. Everything
lands in dev first. Pranav (lead) owns the dev→main merge before the demo.


SECTION F — WHO'S BUILDING WHAT (TASK OWNERSHIP)
================================================================================
Work is split so Pranav and Ujwal — who have Antigravity IDE + Gemini Pro —
carry the AI-agent-heavy modules (multi-file features, realtime UI,
non-deterministic AI-pipeline logic). Pavan's tasks are scoped tightly
enough to build via plain free-tier ChatGPT/Claude chat prompting, no IDE
agent needed — mostly backend CRUD, schema setup, and infra config.

F.1 DEPENDENCY MAP (WHO BLOCKS WHOM)
--------------------------------------------------------------------------------
  Ticket data model (Pavan)
        │
        ├──> Backend API endpoints (Pavan) ──> Frontend API client (Pranav)
        │                                              │
        │                                              └──> Kanban UI (Pranav)
        │
        ├──> Embedding + pgvector dedup (Pavan) ──> Dedup check in intake
        │                                            flow (Pranav, frontend
        │                                            side only — logic is
        │                                            Pavan's)
        │
        └──> SSE event contract (Pavan defines payload shape) ──> SSE
             listener component (Pranav) + SSE emitter calls from sandbox
             (Ujwal)

  Docker sandbox manager (Ujwal) ──> Crash log format (Ujwal defines) ──>
        Gemini RCA prompt consumes that format (Ujwal) ──> RCA JSON result
        posted back to backend (Pavan exposes the endpoint, Ujwal calls it)

  GitHub PR API integration (Ujwal, stretch goal) ──> depends on RCA
        output existing first

PRACTICAL IMPLICATION: On Day 1, all three can work fully in parallel if
Pavan front-loads the ticket schema (30–45 min task) and posts it to the
team channel before writing any other backend code — that one artifact
unblocks both Pranav and Ujwal immediately.

F.2 PRANAV — Team Lead · Frontend & Realtime UX · Integration
    Tool: Antigravity IDE + Gemini Pro | Branch: pranav
--------------------------------------------------------------------------------
OWNERSHIP: Everything the judge sees and clicks. Also final integration
owner — merges dev→main before demo and resolves cross-branch conflicts.

INDEPENDENT (start Day 1, no blockers):
  1. SvelteKit 5 + Tailwind v4 project scaffold, routing, base layout/theme.
  2. Bug Intake Form (/submit) — title, description, stack trace/error log,
     repo branch URL, severity. Client-side validation only (mock submit).
  3. Kanban board shell with dummy data — columns: New → Triaging →
     Sandbox Running → Diagnosed → Resolved.
  4. Ticket Modal component (tabs: Description / Logs / AI Diagnosis / Diff).
  5. SSE listener skeleton connecting to '/api/stream', live terminal-style
     log panel, mocked event source for now.

DEPENDENT (needs Pavan/Ujwal artifacts first):
  6. Wire intake form to Pavan's real POST /tickets endpoint.
  7. Wire Kanban board to Pavan's GET /tickets endpoint.
  8. Wire SSE listener to the real event stream (Pavan's endpoint + Ujwal's
     emitted events).
  9. Render AI diagnosis + diff view once Ujwal's RCA JSON schema lands.
  10. Final integration pass: merge all branches into dev, smoke-test the
      golden path end-to-end.

F.3 UJWAL — Infrastructure, Agentic Sandbox & AI Diagnosis Pipeline
    Tool: Antigravity IDE + Gemini Pro | Branch: ujwal
--------------------------------------------------------------------------------
OWNERSHIP: The "AI-native" heart of the pitch — what actually differentiates
this from a Bugzilla clone.

INDEPENDENT (start Day 1, no blockers):
  1. sandbox/dummy_repo — tiny intentionally-buggy app (e.g. calculator
     with unhandled divide-by-zero). Keep it small and reliable — this is
     the golden-path demo prop.
  2. sandbox/Dockerfile.dummy — minimal image to run dummy_repo's tests.
  3. sandbox/manager.go — Go Docker Engine SDK: pull image → start
     container → copy dummy_repo in → run test → capture stdout/stderr →
     kill + remove container. Standalone-testable, no dependency on
     Pavan/Pranav.
  4. Crash log parser/formatter — define and document the exact JSON shape
     handed to the Gemini RCA step. Post it to the team channel early.
  5. Gemini RCA prompt design — {bug report, crash log, relevant code} →
     strict JSON {root_cause, explanation, file, diff}. Test standalone
     before wiring to the sandbox.
  6. (Stretch) GitHub API integration to open a real draft PR from the diff.

F.4 PAVAN — Backend Core, Database & API
    Tool: Free-tier ChatGPT/Claude (plain chat prompting) | Branch: pavan
--------------------------------------------------------------------------------
OWNERSHIP: The backend spine — data model, CRUD endpoints, embeddings/
dedup storage, and the SSE transport layer. Every task here is scoped
tightly enough for plain chat prompting — describe the file, paste back
the code, test, move to the next file.

INDEPENDENT (start Day 1, no blockers):
  1. Go server bootstrap — main.go, Chi router, health-check route, CORS
     for local frontend dev.
  2. models/ticket.go — Ticket struct (id, title, description, stack
     trace, repo branch URL, severity, status, embedding vector, diagnosis
     JSON, diff text, created_at, updated_at). POST THIS SCHEMA TO THE TEAM
     CHANNEL FIRST — it unblocks both Pranav and Ujwal.
  3. db/neon.go — Postgres connection + pgvector extension + tickets table
     migration.
  4. CRUD endpoints: POST /tickets, GET /tickets, GET /tickets/:id,
     PATCH /tickets/:id — pure persistence, no AI logic yet.
  5. Embedding generation call — wrap the Google Gen AI SDK's embedding
     endpoint (single API call, not agentic reasoning).
  6. pgvector similarity query — find tickets within threshold X of a new
     embedding, wired into POST /tickets for dedup-check.
  7. SSE endpoint (GET /api/stream) — Go SSE handler + in-memory pub/sub so
     other parts of the backend (or Ujwal's sandbox, once wired) can
     publish progress events. Post the event JSON shape to the team
     channel once built.
  8. docker-compose.yml for local dev (frontend + backend together).

DEPENDENT:
  9. Wire the endpoint that receives Ujwal's sandbox/RCA output, writes it
     onto the ticket record, and publishes an SSE "diagnosed" event —
     needs Ujwal's CrashReport/RCA JSON shape.


SECTION G — PROMPTING GUIDE (BY PERSON)
================================================================================
General rule: run these ONE AT A TIME, review each output, don't paste them
all in a single mega-prompt — agentic tools drift and lose the intended
file structure if you ask for too much at once.

G.1 PRANAV (Antigravity + Gemini Pro)
--------------------------------------------------------------------------------
  - "Read bugsinguish_master_readme.txt. I'm building the frontend for
    Pranav's track. Scaffold a SvelteKit 5 project with Tailwind v4, a
    layout with a top nav, and empty routes for '/' (Kanban) and '/submit'
    (intake form)."
  - "Build the bug intake form at src/routes/submit/+page.svelte with
    fields: title, description, stack trace (textarea), repo branch URL,
    severity (select). Client-side validation, no backend call yet — log
    the payload to console on submit."
  - "Build a Kanban board component with 5 columns matching this ticket
    lifecycle: New, Triaging, Sandbox Running, Diagnosed, Resolved. Use
    dummy ticket data (array of 6 mock tickets) for now. Each card shows
    title, severity badge, and a short id."
  - "Build a Ticket Modal component that opens when a Kanban card is
    clicked, with tabs for Description, Logs, AI Diagnosis, and Diff.
    Use dummy content for the Diagnosis and Diff tabs."
  - "Build an SSE listener component that connects to '/api/stream' and
    renders incoming messages as a live scrolling terminal-style log. For
    now, mock the stream locally with a setInterval that emits sample log
    lines every 800ms so I can test the UI without a backend."
  - (Later) "Replace the mock ticket data and mock SSE source in [file]
    with real calls to [Pavan's endpoint URLs + payload shape, pasted in]."

G.2 UJWAL (Antigravity + Gemini Pro)
--------------------------------------------------------------------------------
  - "Read bugsinguish_master_readme.txt. I'm building the sandbox +
    diagnosis pipeline for Ujwal's track. Create sandbox/dummy_repo: a
    tiny Python or Node calculator app with a divide function that has an
    unhandled ZeroDivisionError bug, plus a test file that triggers it."
  - "Write sandbox/Dockerfile.dummy that builds a minimal image capable of
    running dummy_repo's test suite."
  - "Write a Go package sandbox/manager.go using
    'github.com/docker/docker/client' with a function that: pulls a base
    image, starts a container, copies in the dummy_repo test app, runs the
    test command, captures stdout/stderr, and then kills and removes the
    container. Return the captured logs as a struct."
  - "Define a Go struct and JSON schema for a 'CrashReport' — fields for
    stdout, stderr, exit code, failing file, and timestamp — that will be
    passed to an AI root-cause-analysis step."
  - "Write a Gemini prompt (Go, using the Google Gen AI SDK, structured
    JSON output) that takes a bug description, a CrashReport, and a
    snippet of source code, and returns strict JSON: root_cause (string),
    explanation (string), file (string), diff (string, unified diff
    format). Include a test harness with a sample dummy_repo crash so I
    can validate the output shape before wiring it into the pipeline."

G.3 PAVAN (Free-tier ChatGPT/Claude, plain chat)
--------------------------------------------------------------------------------
  - "I'm building a Go backend using the Chi router for a bug-tracking
    app. Write main.go that sets up a Chi router, a health-check route at
    GET /health, and CORS middleware allowing requests from
    http://localhost:5173."
  - "Write a Go struct called Ticket for a bug tracker with these fields:
    ID (uuid), Title, Description, StackTrace, RepoBranchURL, Severity,
    Status (enum: new, triaging, sandbox_running, diagnosed, resolved),
    Embedding ([]float32), Diagnosis (JSON: root_cause, explanation, file,
    diff), CreatedAt, UpdatedAt. Include JSON tags."
  - "Write a Go file that connects to a Neon Postgres database using
    pgx, enables the pgvector extension, and creates a migration for a
    tickets table matching this struct: [paste struct]."
  - "Write Chi HTTP handlers for: POST /tickets (create), GET /tickets
    (list all), GET /tickets/:id (get one), PATCH /tickets/:id (update
    status). Use the Ticket struct and Postgres connection from before."
  - "Write a Go function that calls the Google Gen AI SDK to generate a
    text embedding for a given string, returning []float32."
  - "Write a Postgres query using pgvector's <-> operator to find the
    closest existing ticket embedding to a new one, returning matches
    under a similarity threshold. Wire this into the POST /tickets handler
    so it checks for duplicates before inserting a new row."
  - "Write a basic Server-Sent Events handler in Go using Chi at GET
    /api/stream, with a simple in-memory pub/sub so other parts of the
    app can call a Publish(event) function and have it pushed to all
    connected clients as SSE messages. Show me the JSON shape you're using
    for events."
  - "Write a docker-compose.yml that runs a SvelteKit frontend on port
    5173 and this Go backend on port 8080 for local development."


SECTION H — TIMELINE
================================================================================
  Hour 0–1   : Pavan posts Ticket schema + SSE event shape drafts to the
               team channel (even before finishing the code) so Pranav and
               Ujwal can start against a stable contract.
  Hour 0–6   : All three work independent sections in parallel (Section F
               "INDEPENDENT" items).
  Hour 6–10  : First integration checkpoint — Pranav wires real API calls
               where contracts are stable; Ujwal wires sandbox → backend
               handoff; Pavan finishes SSE publish plumbing.
  Hour 10–14 : Second integration checkpoint — full golden path (submit →
               dedup check → sandbox → RCA → SSE stream → Kanban update)
               run end-to-end on the dummy_repo bug.
  Hour 14+   : Bug-fixing, polish, analytics view, demo rehearsal.


SECTION I — DEFINITION OF DONE (BEFORE DEMO DAY)
================================================================================
  [ ] Golden path runs live, start to finish, without manual intervention
  [ ] Kanban board reflects real ticket state, not dummy data
  [ ] SSE log stream visibly updates during sandbox + AI diagnosis
  [ ] At least one basic analytics view (ticket count by status, minimum)
  [ ] Innovation comparison table (Section A.4) present in the repo README
  [ ] This master file included in the repo as documentation evidence of
      team process (rubric gives credit for documentation)
  [ ] dev merged into main, tagged for demo
================================================================================
