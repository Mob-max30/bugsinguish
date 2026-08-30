<script lang="ts">
  import { onMount } from 'svelte';
  import KanbanBoard from '$lib/components/KanbanBoard.svelte';
  import TicketModal from '$lib/components/TicketModal.svelte';
  import SseTerminal from '$lib/components/SseTerminal.svelte';
  import type { Ticket } from '$lib/types';
  import { fetchTickets } from '$lib/api/client';

  let selectedTicket: Ticket | null = null;
  let isModalOpen = false;
  let isLoading = false;
  let isBackendConnected = false;

  let tickets: Ticket[] = [
    {
      id: 'BUG-101',
      title: 'ZeroDivisionError in calculator divide operation',
      description: 'Calculator crashes immediately when user inputs 0 as denominator in division function.',
      stack_trace: 'Traceback (most recent call last):\n  File "test_calc.py", line 12, in test_divide\n    result = divide(10, 0)\n  File "calc.py", line 14, in divide\n    return a / b\nZeroDivisionError: division by zero',
      repo_branch_url: 'https://github.com/Mob-max30/bugsinguish/tree/main',
      severity: 'critical',
      status: 'diagnosed',
      diagnosis: {
        root_cause: 'Missing zero denominator validation in divide() function at calc.py line 14.',
        explanation: 'The divide function accepts integer arguments a and b without checking if b == 0 before performing division operator. This raises ZeroDivisionError unhandled.',
        file: 'sandbox/dummy_repo/calculator.py',
        diff: '--- calculator.py\n+++ calculator.py\n@@ -13,3 +13,5 @@\n def divide(a, b):\n+    if b == 0:\n+        raise ValueError("Cannot divide by zero")\n     return a / b'
      },
      created_at: '2026-08-30 18:00',
      updated_at: '2026-08-30 18:05'
    },
    {
      id: 'BUG-102',
      title: 'NullPointerException on User Profile fetch',
      description: 'User dashboard fails to render when avatar URL is undefined in Neon DB response.',
      stack_trace: 'Unhandled Exception: NullPointerException: Cannot read property "avatar" of undefined',
      repo_branch_url: 'https://github.com/Mob-max30/bugsinguish/tree/dev',
      severity: 'high',
      status: 'sandbox_running',
      created_at: '2026-08-30 18:15',
      updated_at: '2026-08-30 18:20'
    },
    {
      id: 'BUG-103',
      title: 'Authentication token expiration silently breaks SSE subscription',
      description: 'When JWT expires after 1 hour, SSE stream disconnects without triggering reconnect logic in UI.',
      stack_trace: 'EventSource failed to connect to /api/stream: 401 Unauthorized',
      repo_branch_url: 'https://github.com/Mob-max30/bugsinguish/tree/main',
      severity: 'medium',
      status: 'triaging',
      created_at: '2026-08-30 18:22',
      updated_at: '2026-08-30 18:25'
    },
    {
      id: 'BUG-104',
      title: 'Database pool connection leak during peak vector search queries',
      description: 'pgvector queries exhaust max open connections under heavy concurrent triage loads.',
      stack_trace: 'pgx: max connections reached (100/100). Connection timeout after 5000ms',
      repo_branch_url: 'https://github.com/Mob-max30/bugsinguish/tree/main',
      severity: 'high',
      status: 'new',
      created_at: '2026-08-30 18:30',
      updated_at: '2026-08-30 18:30'
    },
    {
      id: 'BUG-105',
      title: 'CORS header missing on GET /tickets/:id endpoint',
      description: 'Frontend local dev server on port 5173 receives CORS preflight failure from Chi backend.',
      stack_trace: 'Access to fetch at http://localhost:8080/tickets/BUG-101 has been blocked by CORS policy',
      repo_branch_url: 'https://github.com/Mob-max30/bugsinguish/tree/dev',
      severity: 'low',
      status: 'resolved',
      diagnosis: {
        root_cause: 'Chi router missing cors.Handler middleware wrapper on API subrouter.',
        explanation: 'Added github.com/go-chi/cors middleware to main router allowing http://localhost:5173.',
        file: 'backend/main.go',
        diff: '--- main.go\n+++ main.go\n@@ -15,2 +15,3 @@\n+    r.Use(cors.AllowAll().Handler)'
      },
      created_at: '2026-08-30 17:00',
      updated_at: '2026-08-30 17:30'
    }
  ];

  async function loadBackendTickets() {
    isLoading = true;
    try {
      const data = await fetchTickets();
      if (Array.isArray(data) && data.length > 0) {
        tickets = data;
        isBackendConnected = true;
      }
    } catch {
      isBackendConnected = false;
    } finally {
      isLoading = false;
    }
  }

  onMount(() => {
    loadBackendTickets();
  });

  function openTicketModal(ticket: Ticket) {
    selectedTicket = ticket;
    isModalOpen = true;
  }

  function closeTicketModal() {
    isModalOpen = false;
    selectedTicket = null;
  }
</script>

<div class="space-y-6">
  <!-- Top Banner -->
  <div class="flex items-center justify-between">
    <div>
      <h1 class="text-2xl font-black text-slate-100">Defect Resolution Kanban</h1>
      <p class="text-sm text-slate-400 mt-1">Autonomous bug triage, sandbox reproduction, and AI root-cause analysis pipeline.</p>
    </div>
    
    <div class="flex items-center space-x-3 text-xs bg-slate-900 px-4 py-2 rounded-lg border border-slate-800">
      <span class="text-slate-400">Total Tickets: <strong class="text-white">{tickets.length}</strong></span>
      <span class="text-slate-600">|</span>
      <span class="text-slate-400">Diagnosed: <strong class="text-indigo-400">{tickets.filter(t => t.status === 'diagnosed').length}</strong></span>
      <span class="text-slate-600">|</span>
      <span class="flex items-center space-x-1">
        <span class={`w-2 h-2 rounded-full ${isBackendConnected ? 'bg-emerald-400' : 'bg-amber-400'}`}></span>
        <span class="text-slate-400">{isBackendConnected ? 'API Live' : 'Demo Mode'}</span>
      </span>
    </div>
  </div>

  <!-- Live SSE Stream Panel -->
  <SseTerminal />

  <!-- Kanban Board -->
  <KanbanBoard {tickets} onSelectTicket={openTicketModal} />

  <!-- Ticket Details Modal -->
  <TicketModal ticket={selectedTicket} isOpen={isModalOpen} onClose={closeTicketModal} />
</div>
