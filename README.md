# Bugsinguish 🐛🔥

**Deconstructing Bugzilla into an AI-Native Defect Resolution Engine**

*While conventional bug trackers act as passive filing cabinets that store text complaints and wait weeks for human developers to investigate, Bugsinguish automates the entire pipeline. The moment a bug report arrives, our platform converts the report into high dimensional vector embeddings to group duplicate issues instantly, clones the codebase into an isolated temporary test sandbox to reproduce stack traces automatically, diagnoses root causes using Google Gemini 1.5 Pro, and generates ready to merge Pull Requests for single click approval.*

---

## 🖥️ Live Platform Screenshots & Dashboard Features

Below are screenshots from our live, working application hosted at **[https://bugsinguish-phi.vercel.app](https://bugsinguish-phi.vercel.app)**. All displayed numbers, vector similarity scores, activity logs, and status counts are **dynamically generated** from live database state and test data in the `sandbox/` folder (zero hardcoded static numbers, you can test it for yourself xD).

![Bugsinguish 3D Resolution Pipeline Dashboard](./dashboard_top.png)
*Top Section: KPI Metrics, Three.js 3D Interactive Pipeline Terrain, and AI Triage Overview Donut Chart.*

![Bugsinguish Recent Issues & Live AI Activity Feed](./dashboard_bottom.png)
*Bottom Section: Dynamic Recent Issues Table and Real-time AI Activity Feed.*

---

## 🚀 Now, you may wonder.... what exactly is Bugsinguish 🤔? 

Well, traditional bug trackers like Bugzilla are passive "digital filing cabinets." When something breaks, they store a record of the complaint and wait weeks for a human engineer to manually search through past tickets, reproduce the issue, and write a fix.

**Bugsinguish transforms bug tracking from a static filing cabinet into an "Active Resolution Engine."** 
When a bug is reported, Bugsinguish:
1. Instantly understands what the bug means using AI (catching duplicate reports even if written in completely different words).
2. Clones the affected code into an isolated, temporary test sandbox.
3. Diagnoses the exact root cause and drafts a ready-to-merge Pull Request (PR) with the fix.
4. Leaves the human engineer in control with a single click to review and merge.

## ⚠️ The Problem: Why did Bugzilla fall behind?

Bugzilla was revolutionary in 1998, but modern software development has outgrown its design philosophy.

* **Pain Point 1 — Duplicate Ticket Nightmare:** When a popular app crashes, 500 users file tickets using different words. A human admin must read and manually tag all 500 duplicates.
* **Pain Point 2 — Disconnected from Code ("Works on My Machine"):** Bugzilla tickets are just plain text — no branch, no environment, no way to reproduce without manual dev setup.
* **Pain Point 3 — The Slow Manual Bottleneck:** A ticket sits in queue for days before a developer even opens the repo.
* **Pain Point 4 — Enterprise Privacy & Storage Bloat:** Storing decades of raw code/crash logs creates compliance and security liabilities.
* **Pain Point 5 — Intimidating 1990s Interface:** A database spreadsheet with 40+ required fields, alienating everyone but specialized QA engineers.

## 🛠️ How Bugsinguish Solves Each Problem

| Legacy Bugzilla Problem | How Bugsinguish Solves It |
| :--- | :--- |
| **Exact-match search misses duplicates.** | **Semantic Deduplication:** Reports become vector "meanings." AI detects "App crashes on login" and "Cannot sign in on iOS" are the same issue and groups them automatically. |
| **Bugs lack live code context.** | **Ephemeral Docker Sandbox:** Spins up an isolated, lightweight container, pulls the exact branch, and executes tests. |
| **Developers spend hours debugging.** | **Agentic Root Cause Analysis & Auto-PR:** AI reads sandbox crash logs, pinpoints the broken line, and drafts a proposed code fix. |
| **Enterprises fear data leaks in logs.** | **Zero-Retention Privacy:** Raw code/logs discarded after resolution — only anonymized vectors kept. |
| **Clunky, slow page reloads.** | **Real-Time Reactive UI:** SvelteKit 5 streams live AI debugging logs step-by-step via SSE. |

## 🌟 How We Differ From Regular Bugzilla (Innovation Table)

| Feature | Legacy Bugzilla | Bugsinguish |
| :--- | :--- | :--- |
| **Paradigm** | Passive Record-Keeping | Active Resolution |
| **UI/UX** | 1990s CGI, Page Refresh | SvelteKit, Real-Time SSE Updates |
| **Triage** | Human reads & tags | AI vector-matches & groups instantly |
| **Context** | Copy-pasted text/logs | Live codebase connection via Sandbox |
| **End Goal** | Human marks "Resolved" | AI drafts PR, Human clicks "Merge" |

## 🏗️ System Architecture

```text
bugsinguish/
├── frontend/                        # SvelteKit 5 + Tailwind v4 + Three.js UI
│   ├── src/
│   │   ├── lib/
│   │   │   ├── api/
│   │   │   │   ├── client.ts        --> Connects to Backend REST API (POST /tickets, GET /tickets)
│   │   │   │   └── sse.ts           --> Subscribes to Backend SSE Stream (GET /api/stream)
│   │   │   └── components/
│   │   │       ├── KanbanBoard.svelte  --> Displays 5-stage defect lifecycle cards
│   │   │       ├── TicketModal.svelte  --> Renders AI diagnosis, diffs, and execution logs
│   │   │       ├── SseTerminal.svelte  --> Monospace terminal streaming live agent steps
│   │   │       └── Pipeline3D.svelte   --> Three.js 3D WebGL interactive resolution terrain
│   │   └── routes/                  --> Dynamic routes (/, /issues, /triage, /sandboxes, /submit)
│   └── package.json
│
├── backend/                         # Go (Golang) + Chi Router Server
│   ├── db/
│   │   └── neon.go                  --> Neon Postgres connection & pgvector extension migrations
│   ├── embeddings/
│   │   └── embed.go                 --> Google Gen AI SDK 768-dim text embedding generator
│   ├── handlers/
│   │   └── tickets.go               --> REST HTTP handlers & async pipeline trigger worker
│   ├── models/
│   │   └── ticket.go                --> Ticket, Diagnosis, & Status struct models
│   ├── sse/
│   │   └── stream.go                --> In-memory pub/sub broker pushing live SSE events to UI
│   ├── main.go                      --> Main Chi router entry point listening on :8080
│   └── go.mod
│
├── sandbox/                         # Ephemeral Docker Manager & AI Engine
│   ├── dummy_repo/                  --> Test repository prop with intentional ZeroDivisionError
│   ├── manager.go                   --> Native Docker Engine SDK (spawns & removes container)
│   ├── rca.go                       --> Gemini 1.5 Pro prompt harness (structured JSON & diffs)
│   └── go.mod
│
└── docker-compose.yml               # Local container orchestration for frontend & backend
```

## ⚡ Quickstart Guide: Running & Accessing the Project

### 1. Live Production Deployment
* **Live Web App:** [https://bugsinguish-phi.vercel.app](https://bugsinguish-phi.vercel.app)
* **Backend API (Render):** `https://bugsinguish-api.onrender.com` (Healthcheck: `https://bugsinguish-api.onrender.com/health`)
* **Database:** Neon Serverless Postgres with `pgvector` enabled

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
