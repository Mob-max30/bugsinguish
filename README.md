# Bugsinguish 🐛🔥

**Deconstructing Bugzilla into an AI-Native Defect Resolution Engine**

*Track 2: Developer Tool Reconstruction – Bugzilla (CloneFest)*

**Team:** Pranav (Lead) · Ujwal · Pavan  
**Repo:** [https://github.com/Mob-max30/bugsinguish](https://github.com/Mob-max30/bugsinguish)

---

## 🚀 What is Bugsinguish? (30-Second Elevator Pitch)

Traditional bug trackers like Bugzilla are passive "digital filing cabinets." When something breaks, they store a record of the complaint and wait weeks for a human engineer to manually search through past tickets, reproduce the issue, and write a fix.

**Bugsinguish transforms bug tracking from a static filing cabinet into an "Active Resolution Engine."** When a bug is reported, Bugsinguish:
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

* **Frontend:** SvelteKit 5 + Tailwind CSS v4 + Bits UI — SSE for live terminal-style logs, no page reloads.
* **Backend/API:** Go (Golang) + Chi Router — handles concurrent triage requests with high performance.
* **Database:** Neon Serverless Postgres + pgvector — relational ticket state and semantic embeddings side-by-side.
* **AI/Agent Layer:** Google Gen AI SDK — structured JSON outputs for predictable code patching and embeddings.
* **Sandbox/Infra:** Native Docker Engine API (Go SDK) — short-lived containers destroyed immediately after use.

## 🗺️ The Step-by-Step User Journey

### PHASE 1 — Intake and Instant Triage (The Gateway)
1. User (or automated script) submits a bug description, crash log, and repo branch URL.
2. Go backend calls the Google Gen AI SDK to convert the text into a vector embedding.
3. Backend queries Neon Postgres (pgvector) for similar embeddings.
4. If a duplicate is found → user is alerted, logs appended to the existing ticket. If not → a new ticket is created with a Ticket ID.

### PHASE 2 — Autonomous Reproduction (The Sandbox)
1. Backend spins up a lightweight Docker container.
2. Container clones the specified repo branch and sets up the environment.
3. Container executes a test to reproduce the stack trace.
4. Crash logs are extracted, and the container is immediately destroyed (Zero-Retention).

### PHASE 3 — AI Diagnosis and Patch Generation (The RCA)
1. Backend sends the bug report, code context, and crash logs to Google Gemini.
2. Gemini outputs a strict JSON schema with a human-readable root cause analysis.
3. Gemini generates a git patch/code diff to fix the bug.
4. UI updates in real-time via SSE, showing the Draft PR forming.

### PHASE 4 — Human-in-the-Loop Review (The Resolution)
1. Developer reviews the AI's diagnosis and code diff on the dashboard.
2. Developer clicks "Approve & Merge" — platform opens/merges the PR via the GitHub API.
3. Ticket is marked RESOLVED, and raw code data is purged.
