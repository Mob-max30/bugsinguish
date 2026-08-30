# Bugsinguish 🐛🔥

> **Autonomous AI Defect Resolution Engine**  
> *CloneFest Track 2: Developer Tool Reconstruction (Bugzilla)*  
> **Team:** Pranav (Lead) · Ujwal · Pavan  
> 🌐 **Live Running Web App:** [https://bugsinguish-phi.vercel.app](https://bugsinguish-phi.vercel.app)  
> ⚡ **Live Production API:** [https://bugsinguish-api.onrender.com](https://bugsinguish-api.onrender.com)  

---

## 📂 Monorepo File Structure & Interconnected Pipeline

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

---

## 🌐 Live Experience and Interactive Demo

Experience the future of defect tracking live in your browser:

👉 **[Launch Live Bugsinguish Application](https://bugsinguish-phi.vercel.app)**

Test autonomous vector deduplication, interactive 3D pipeline visualization, and single-click AI pull request generation directly on our live platform.

![Bugsinguish 3D Resolution Pipeline Dashboard](docs/screenshots/dashboard_top.png)
*Top Section: KPI Metrics, Three.js 3D Interactive Pipeline Terrain, and AI Triage Overview Donut Chart.*

![Bugsinguish Recent Issues & Live AI Activity Feed](docs/screenshots/dashboard_bottom.png)
*Bottom Section: Dynamic Recent Issues Table and Real-time AI Activity Feed.*

---

## 🚀 Welcome to Bugsinguish

Traditional bug trackers like Bugzilla operate as passive digital filing cabinets. When software breaks, legacy systems store a text record of the problem and wait weeks for human developers to manually read tickets, recreate complex environments, and write code fixes.

Bugsinguish turns bug tracking into an active resolution pipeline. The moment a bug report arrives, Bugsinguish:

1. Understands the semantic meaning of the issue using high dimensional vector embeddings to catch duplicate reports even when written in completely different words.
2. Clones the target codebase branch into an isolated temporary test sandbox container to reproduce the stack trace automatically.
3. Diagnoses the underlying root cause using Google Gemini 1.5 Pro and drafts a ready to merge Pull Request code fix.
4. Empowers engineers with a single click to review, approve, and merge automated fixes into production.

---

## ⚠️ The Legacy Problem

Software development has outgrown the manual workflow of legacy bug tracking tools.

* **Duplicate Ticket Overload:** When an app crashes, hundreds of users file tickets using different phrasing. Human administrators spend hours reading and manually linking duplicate reports.
* **Disconnected Code Context:** Plain text stack traces lack live environment context, leading to frustrating setup bottlenecks.
* **Slow Manual Debugging:** Tickets sit in queues for days before a developer opens the target codebase.
* **Privacy and Storage Liabilities:** Storing decades of raw error logs creates enterprise compliance risks.
* **Outdated User Interfaces:** Complex multi field forms alienate team members and slow down development velocity.

---

## 🛠️ How Bugsinguish Transforms Defect Resolution

| Legacy Problem | Bugsinguish Resolution |
| :--- | :--- |
| **Exact match search misses duplicates** | **Semantic Deduplication:** Converts reports into vector meanings via Neon pgvector to group duplicates instantly. |
| **Bugs lack live code context** | **Ephemeral Docker Sandbox:** Spins up lightweight containers on demand to execute test suites automatically. |
| **Hours spent manually debugging** | **Agentic Root Cause Analysis:** Gemini analyzes crash logs and drafts unified git patches ready for review. |
| **Privacy risks from log storage** | **Zero Retention Privacy:** Ephemeral test environments and raw logs are purged immediately post resolution. |
| **Page reloads and slow forms** | **Real Time Reactive UI:** SvelteKit 5 streams live debugging events step by step via Server Sent Events. |

---

## 🌟 Key Platform Innovations

| Feature | Legacy Bugzilla | Bugsinguish |
| :--- | :--- | :--- |
| **Paradigm** | Passive Record Keeping | Active Resolution Engine |
| **User Experience** | Legacy Forms and Page Reloads | SvelteKit 5, Real Time Log Streaming |
| **Triage** | Manual Human Tagging | Instant Vector Similarity Matching |
| **Code Context** | Static Plain Text Logs | Live Codebase Connection via Sandbox |
| **Resolution Target** | Human Manually Marks Resolved | AI Drafts Pull Request, Human Merges with 1 Click |

---

## ⚡ Quickstart Guide: Running & Accessing the Project

### 1. Live Production Deployment
* **Live Running Web App:** [https://bugsinguish-phi.vercel.app](https://bugsinguish-phi.vercel.app)
* **Backend API (Render):** [https://bugsinguish-api.onrender.com](https://bugsinguish-api.onrender.com) (Healthcheck: `https://bugsinguish-api.onrender.com/health`)
* **Database:** Neon Serverless Postgres with pgvector enabled

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
The Go server automatically connects to Neon, enables pgvector, migrates the database schema, and starts on port 8080.

#### C. Run via Docker Compose
```bash
docker-compose up --build
```
