# Bugsinguish

> **Active Resolution Engine** — An AI-Native Defect Resolution and Automated Sandbox Diagnosis Platform (CloneFest Track 2: Developer Tool Reconstruction – Bugzilla).

## 🚀 Innovation Delta vs Legacy Bugzilla

| Feature | Legacy Bugzilla | Bugsinguish |
|---|---|---|
| **Paradigm** | Passive Record-Keeping ("Digital Filing Cabinet") | **Active Resolution Engine** (Autonomous sandbox & auto-PR) |
| **UI/UX** | 1990s CGI Forms, Full Page Reloads | **SvelteKit 5 + Tailwind CSS v4**, Real-time SSE Stream |
| **Triage** | Manual human reading & duplicate tagging | **Semantic Vector Deduplication** via Neon Postgres (`pgvector`) |
| **Context** | Copy-pasted plain text stack traces | **Live Codebase Connection** via Ephemeral Docker Sandbox |
| **End Goal** | Human developer manually marks "Resolved" | **AI drafts Pull Request**, Human developer clicks "Merge" |

---

## 🛠️ System Architecture

* **Frontend**: SvelteKit 5 + Tailwind CSS v4 — Real-time SSE terminal logs, 5-stage Kanban board, 4-tab Ticket Modal.
* **Backend**: Go (Golang) + Chi Router — High-performance concurrent API handlers.
* **Database**: Neon Serverless Postgres + `pgvector` — Relational state and 768-dim vector embeddings.
* **AI Engine**: Google Gen AI SDK — Root-cause analysis (RCA) and git patch diff generation.
* **Sandbox**: Ephemeral Docker Engine SDK — Isolated temporary containers for test reproduction.

---

## 📄 Team Work Breakdown & System Plan
For the full task ownership, dependency map, and hackathon demo script, see [bugsinguish_master_readme.txt](bugsinguish_master_readme.txt).
