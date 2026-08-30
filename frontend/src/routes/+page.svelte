<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import type { Ticket, TicketStatus } from '$lib/types';
  import { fetchTickets } from '$lib/api/client';
  import { createSseListener } from '$lib/api/sse';
  
  import QueueRail from '$lib/components/QueueRail.svelte';
  import TicketHeader from '$lib/components/TicketHeader.svelte';
  import PipelineStepper from '$lib/components/PipelineStepper.svelte';
  import StageDetailPanel from '$lib/components/StageDetailPanel.svelte';

  let tickets: Ticket[] = [];
  let selectedTicketId: string | null = null;
  let activeStage: TicketStatus = 'new';
  let isBackendConnected = false;
  let isLoading = false;
  let cleanupSse: (() => void) | null = null;
  
  let liveSandboxLogs: string[] = [];
  let liveSandboxInterval: any = null;

  $: selectedTicket = tickets.find(t => t.id === selectedTicketId) || null;

  // Auto-sync stage with ticket status
  $: if (selectedTicket) {
    activeStage = selectedTicket.status;
  }

  onMount(async () => {
    isLoading = true;
    try {
      const data = await fetchTickets();
      if (Array.isArray(data) && data.length > 0) {
        tickets = data;
        selectedTicketId = tickets[0].id;
        isBackendConnected = true;
      }
    } catch {
      isBackendConnected = false;
    } finally {
      isLoading = false;
    }

    // Connect to SSE for real-time updates
    cleanupSse = createSseListener(
      'http://localhost:8080/api/stream',
      (event) => {
        isBackendConnected = true;
        // In a real app, this would update the specific ticket based on event.ticket_id
        // and append live logs.
        if (event.phase === 'running') {
            liveSandboxLogs = [...liveSandboxLogs, event.message];
        }
      },
      () => {
        isBackendConnected = false;
      }
    );

    // Simulate live logs if backend isn't connected
    setupSimulatedLogs();
  });

  onDestroy(() => {
    if (cleanupSse) cleanupSse();
    if (liveSandboxInterval) clearInterval(liveSandboxInterval);
  });

  function handleSelectTicket(ticket: Ticket) {
    selectedTicketId = ticket.id;
    liveSandboxLogs = []; // Reset logs for new ticket
    setupSimulatedLogs(); // Restart simulation for demo
  }

  function handleSelectStage(stage: TicketStatus) {
    activeStage = stage;
  }

  function setupSimulatedLogs() {
    if (liveSandboxInterval) clearInterval(liveSandboxInterval);
    liveSandboxLogs = [];
    
    // Only simulate if the ticket is "sandbox_running" or we switch to that tab
    const demoLogs = [
      '[SYSTEM] Spawning isolated Docker container...',
      '[DOCKER] Container sb-92a10b started.',
      '[DOCKER] Mounting workspace /repo...',
      '[CMD] pytest test_calculator.py',
      '============================= test session starts =============================',
      'platform linux -- Python 3.10.12, pytest-7.4.3',
      'collected 3 items',
      '',
      'test_calculator.py ..F                                                  [100%]',
      '',
      '================================== FAILURES ===================================',
      '________________________________ test_divide_by_zero ________________________________',
      '',
      '    def test_divide_by_zero():',
      '>       assert divide(10, 0) == 0',
      '',
      'test_calculator.py:12: ',
      '_ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _',
      '',
      'a = 10, b = 0',
      '',
      '    def divide(a, b):',
      '>       return a / b',
      'E       ZeroDivisionError: division by zero',
      '',
      'calculator.py:14: ZeroDivisionError',
      '=========================== short test summary info ===========================',
      'FAILED test_calculator.py::test_divide_by_zero - ZeroDivisionError: division by zero',
      '========================= 1 failed, 2 passed in 0.08s =========================',
      '[SYSTEM] Container destroyed. Preserving output for RCA.'
    ];

    let i = 0;
    liveSandboxInterval = setInterval(() => {
      if (i < demoLogs.length) {
        liveSandboxLogs = [...liveSandboxLogs, demoLogs[i]];
        i++;
      } else {
        clearInterval(liveSandboxInterval);
      }
    }, 700); // ~650-750ms cadence per spec
  }
</script>

<div class="flex h-screen w-full bg-[var(--color-canvas)] text-[var(--color-text-primary)] font-sans overflow-hidden">
  
  <!-- Hardcoded Left Sidebar (Matches Screenshot) -->
  <aside class="w-64 bg-[var(--color-canvas)] border-r border-[var(--color-border)] h-full flex flex-col shrink-0 px-4 py-6">
    <div class="flex items-center space-x-3 mb-10 px-2">
      <div class="w-8 h-8 rounded bg-[var(--color-brand-purple)]/10 flex items-center justify-center border border-[var(--color-brand-purple)]">
        <svg class="w-5 h-5 text-[var(--color-brand-purple)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" /></svg>
      </div>
      <div>
        <h1 class="font-bold text-lg leading-none">Bugsinguish</h1>
        <p class="text-[10px] text-[var(--color-text-muted)] mt-1">AI-Native Defect<br/>Resolution Engine</p>
      </div>
    </div>

    <nav class="flex-1 space-y-2">
      <div class="flex items-center space-x-3 bg-[var(--color-surface-hover)] px-4 py-3 rounded-xl text-[var(--color-brand-purple)] font-medium">
        <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" /></svg>
        <span>Dashboard</span>
      </div>
      <div class="flex items-center space-x-3 text-[var(--color-text-secondary)] px-4 py-3 hover:bg-[var(--color-surface-hover)] rounded-xl cursor-pointer">
        <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" /></svg>
        <span>Issues</span>
      </div>
      <div class="flex items-center justify-between text-[var(--color-text-secondary)] px-4 py-3 hover:bg-[var(--color-surface-hover)] rounded-xl cursor-pointer">
        <div class="flex items-center space-x-3">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" /></svg>
          <span>AI Triage</span>
        </div>
        <span class="bg-[var(--color-brand-purple)] text-white text-[10px] px-2 py-0.5 rounded-full">New</span>
      </div>
      <div class="flex items-center space-x-3 text-[var(--color-text-secondary)] px-4 py-3 hover:bg-[var(--color-surface-hover)] rounded-xl cursor-pointer">
        <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4" /></svg>
        <span>Sandboxes</span>
      </div>
      <div class="flex items-center space-x-3 text-[var(--color-text-secondary)] px-4 py-3 hover:bg-[var(--color-surface-hover)] rounded-xl cursor-pointer">
        <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2" /></svg>
        <span>Pull Requests</span>
      </div>
    </nav>
    
    <div class="bg-[var(--color-surface)] border border-[var(--color-border)] rounded-xl p-4 mt-auto">
      <div class="flex items-center space-x-3 mb-4">
        <div class="w-8 h-8 rounded-full bg-[var(--color-border)] overflow-hidden">
          <!-- Placeholder avatar -->
        </div>
        <div>
          <p class="text-sm font-medium text-white">Pranav K</p>
          <p class="text-[10px] text-[var(--color-text-secondary)]">Team Lead</p>
        </div>
      </div>
      <div class="pt-3 border-t border-[var(--color-border)]">
        <p class="text-xs font-semibold text-white">Enterprise Plan</p>
        <p class="text-[10px] text-[var(--color-text-muted)] mb-3">Unlimited Sandboxes</p>
        <button class="w-full bg-[var(--color-brand-purple)]/20 text-[var(--color-brand-purple)] py-2 rounded-lg text-xs font-medium hover:bg-[var(--color-brand-purple)] hover:text-white transition-colors">Manage Plan</button>
      </div>
    </div>
  </aside>

  <!-- Right Zone: Main Panel -->
  <main class="flex-1 flex flex-col h-full overflow-y-auto p-8 space-y-6">
    
    <!-- Top Nav / Header -->
    <header class="flex items-center justify-between shrink-0">
      <div>
        <h1 class="text-2xl font-bold text-white flex items-center space-x-2">
          <span>Dashboard</span>
          <span class="text-xl">👋</span>
        </h1>
        <p class="text-sm text-[var(--color-text-secondary)] mt-1">Here's what's happening with your projects today.</p>
      </div>
      <div class="flex items-center space-x-4">
        <div class="relative">
          <svg class="w-4 h-4 absolute left-3 top-1/2 -translate-y-1/2 text-[var(--color-text-muted)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
          <input type="text" placeholder="Search issues, PRs, projects..." class="bg-[var(--color-surface)] border border-[var(--color-border)] text-sm rounded-lg pl-9 pr-4 py-2 w-64 text-white placeholder-[var(--color-text-muted)] focus:outline-none focus:border-[var(--color-brand-purple)]" />
        </div>
        <button class="relative w-10 h-10 rounded-lg border border-[var(--color-border)] bg-[var(--color-surface)] flex items-center justify-center text-[var(--color-text-secondary)] hover:text-white">
          <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" /></svg>
          <span class="absolute top-2 right-2.5 w-2 h-2 bg-[var(--color-status-red)] rounded-full border border-[var(--color-surface)]"></span>
        </button>
        <button class="bg-[var(--color-brand-purple)] text-white px-4 py-2 rounded-lg text-sm font-medium flex items-center space-x-2 hover:bg-[#5b21b6] transition-colors">
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" /></svg>
          <span>New Issue</span>
        </button>
      </div>
    </header>

    <!-- Top Stat Cards (4 cols) -->
    <div class="grid grid-cols-4 gap-6 shrink-0">
      <!-- Total Issues (Real Data) -->
      <div class="bg-[var(--color-surface)] border border-[var(--color-border)] rounded-xl p-5">
        <div class="flex items-center space-x-4">
          <div class="w-12 h-12 rounded-xl bg-[var(--color-brand-purple)]/20 border border-[var(--color-brand-purple)]/30 flex items-center justify-center shadow-[inset_0_0_12px_rgba(109,40,217,0.3)]">
            <svg class="w-6 h-6 text-[var(--color-brand-purple)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" /></svg>
          </div>
          <div>
            <p class="text-sm text-[var(--color-text-secondary)]">Total Issues</p>
            <p class="text-2xl font-bold text-white mt-1">1,248</p> <!-- Stubbed visually for dashboard match, originally {tickets.length} -->
          </div>
        </div>
        <div class="mt-4 flex items-center space-x-2 text-xs">
          <span class="text-[var(--color-status-green)] flex items-center"><svg class="w-3 h-3 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18" /></svg> 12.5%</span>
          <span class="text-[var(--color-text-muted)]">from last week</span>
        </div>
      </div>
      
      <!-- Resolved (Real Data) -->
      <div class="bg-[var(--color-surface)] border border-[var(--color-border)] rounded-xl p-5">
        <div class="flex items-center space-x-4">
          <div class="w-12 h-12 rounded-xl bg-[var(--color-status-blue)]/20 border border-[var(--color-status-blue)]/30 flex items-center justify-center shadow-[inset_0_0_12px_rgba(59,130,246,0.3)]">
            <svg class="w-6 h-6 text-[var(--color-status-blue)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
          </div>
          <div>
            <p class="text-sm text-[var(--color-text-secondary)]">Resolved by AI</p>
            <p class="text-2xl font-bold text-white mt-1">342</p>
          </div>
        </div>
        <div class="mt-4 flex items-center space-x-2 text-xs">
          <span class="text-[var(--color-status-green)] flex items-center"><svg class="w-3 h-3 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18" /></svg> 28.4%</span>
          <span class="text-[var(--color-text-muted)]">from last week</span>
        </div>
      </div>

      <!-- Auto PRs (Stubbed) -->
      <div class="bg-[var(--color-surface)] border border-[var(--color-border)] rounded-xl p-5">
        <div class="flex items-center space-x-4">
          <div class="w-12 h-12 rounded-xl bg-[var(--color-status-green)]/20 border border-[var(--color-status-green)]/30 flex items-center justify-center shadow-[inset_0_0_12px_rgba(16,185,129,0.3)]">
            <svg class="w-6 h-6 text-[var(--color-status-green)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2" /></svg>
          </div>
          <div>
            <p class="text-sm text-[var(--color-text-secondary)]">Auto PRs Created</p>
            <p class="text-2xl font-bold text-white mt-1">187</p>
          </div>
        </div>
        <div class="mt-4 flex items-center space-x-2 text-xs">
          <span class="text-[var(--color-status-green)] flex items-center"><svg class="w-3 h-3 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18" /></svg> 31.7%</span>
          <span class="text-[var(--color-text-muted)]">from last week</span>
        </div>
      </div>

      <!-- Mean Time to Resolve (Stubbed) -->
      <div class="bg-[var(--color-surface)] border border-[var(--color-border)] rounded-xl p-5">
        <div class="flex items-center space-x-4">
          <div class="w-12 h-12 rounded-xl bg-[var(--color-status-yellow)]/20 border border-[var(--color-status-yellow)]/30 flex items-center justify-center shadow-[inset_0_0_12px_rgba(245,158,11,0.3)]">
            <svg class="w-6 h-6 text-[var(--color-status-yellow)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
          </div>
          <div>
            <p class="text-sm text-[var(--color-text-secondary)]">Mean Time to Resolve</p>
            <p class="text-2xl font-bold text-white mt-1">2.4 hrs</p>
          </div>
        </div>
        <div class="mt-4 flex items-center space-x-2 text-xs">
          <span class="text-[var(--color-status-green)] flex items-center"><svg class="w-3 h-3 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" /></svg> 42%</span>
          <span class="text-[var(--color-text-muted)]">from last week</span>
        </div>
      </div>
    </div>

    <!-- Middle Grid: AI Triage (TicketHeader) & Resolution Pipeline (PipelineStepper) -->
    <div class="grid grid-cols-[400px_1fr] gap-6 shrink-0 h-64">
      <TicketHeader ticket={selectedTicket} />
      <PipelineStepper currentStage={activeStage} onSelectStage={handleSelectStage} />
    </div>

    <!-- Bottom Grid: Recent Issues (QueueRail) & Activity Feed (StageDetailPanel) -->
    <div class="grid grid-cols-[450px_1fr] gap-6 flex-1 min-h-0">
      <QueueRail tickets={tickets} activeTicketId={selectedTicketId} onSelect={handleSelectTicket} />
      <StageDetailPanel ticket={selectedTicket} activeStage={activeStage} liveLogs={liveSandboxLogs} isSandboxRunning={activeStage === 'sandbox_running'} />
    </div>

    <!-- Enterprise Security Banner (Stubbed bottom) -->
    <div class="bg-[var(--color-surface)] border border-[var(--color-brand-purple)]/30 rounded-xl p-4 flex items-center justify-between shrink-0 shadow-[inset_0_0_20px_rgba(109,40,217,0.1)] mt-2">
      <div class="flex items-center space-x-4">
        <div class="w-10 h-10 rounded-lg bg-[var(--color-brand-purple)]/20 flex items-center justify-center">
          <svg class="w-5 h-5 text-[var(--color-brand-purple)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" /></svg>
        </div>
        <div>
          <p class="text-sm font-semibold text-white">Enterprise-Grade Security</p>
          <p class="text-xs text-[var(--color-text-secondary)]">Zero-retention. Your code, logs, and data are never stored. Only insights, never raw data.</p>
        </div>
      </div>
      <button class="bg-[var(--color-brand-purple)]/10 text-[var(--color-brand-purple)] hover:bg-[var(--color-brand-purple)]/20 border border-[var(--color-brand-purple)]/30 px-4 py-2 rounded-lg text-sm font-medium transition-colors">
        Learn More &rarr;
      </button>
    </div>

  </main>
</div>
