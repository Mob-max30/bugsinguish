<script lang="ts">
  import { onMount } from 'svelte';
  import type { Ticket, TicketStatus, TicketSeverity } from '$lib/types';
  import TicketModal from '$lib/components/TicketModal.svelte';
  import { fetchTickets } from '$lib/api/client';

  let selectedTicket: Ticket | null = null;
  let isModalOpen = false;
  let searchQuery = '';
  let selectedStatusFilter: TicketStatus | 'all' = 'all';
  let selectedSeverityFilter: TicketSeverity | 'all' = 'all';

  let tickets: Ticket[] = [
    {
      id: 'BUG-1248',
      title: 'Users unable to login on iOS app',
      description: 'Authentication tokens fail to sign on iOS devices with 401 response.',
      stack_trace: 'Traceback (most recent call last):\n  File "AuthService.java", line 42, in authenticateUser\nNullPointerException: token is null',
      repo_branch_url: 'iOS App / main',
      severity: 'critical',
      status: 'diagnosed',
      diagnosis: {
        root_cause: 'Null pointer exception in AuthService.java:42',
        explanation: 'Tokens expired without refresh handler.',
        file: 'AuthService.java',
        diff: '--- AuthService.java\n+++ AuthService.java\n@@ -41,3 +41,5 @@\n+if (token == null) {\n+    throw new AuthException("Token missing");\n+}'
      },
      created_at: '2026-08-30 22:00',
      updated_at: '2026-08-30 22:05'
    },
    {
      id: 'BUG-1247',
      title: 'Data not syncing across devices',
      description: 'Background worker pool exhausts connection timeout under load.',
      stack_trace: 'pgx: connection timeout after 5000ms',
      repo_branch_url: 'Sync Service / dev',
      severity: 'high',
      status: 'sandbox_running',
      created_at: '2026-08-30 21:45',
      updated_at: '2026-08-30 21:50'
    },
    {
      id: 'BUG-1246',
      title: 'Payment gateway timeout error',
      description: 'CORS header missing on GET /checkout endpoint.',
      stack_trace: 'Access to fetch blocked by CORS policy',
      repo_branch_url: 'Web Checkout / main',
      severity: 'medium',
      status: 'triaging',
      created_at: '2026-08-30 21:00',
      updated_at: '2026-08-30 21:05'
    },
    {
      id: 'BUG-1245',
      title: 'UI glitch in dark mode contrast',
      description: 'Contrast on column card headers drops below 3.0 ratio.',
      stack_trace: '',
      repo_branch_url: 'Web App / main',
      severity: 'low',
      status: 'new',
      created_at: '2026-08-30 20:30',
      updated_at: '2026-08-30 20:30'
    },
    {
      id: 'BUG-1244',
      title: 'Crash on clicking export PDF',
      description: 'Export engine buffer overflows when rendering large reports.',
      stack_trace: 'BufferOverflowError: max length 10MB exceeded',
      repo_branch_url: 'Reports / main',
      severity: 'critical',
      status: 'resolved',
      diagnosis: {
        root_cause: 'Buffer overflow in PDF exporter',
        explanation: 'Max page limit unhandled.',
        file: 'pdf_exporter.go',
        diff: '--- pdf_exporter.go\n+++ pdf_exporter.go\n@@ -10,3 +10,5 @@\n+if len(buf) > maxLimit {\n+    return ErrLimitExceeded\n+}'
      },
      created_at: '2026-08-30 19:00',
      updated_at: '2026-08-30 19:30'
    }
  ];

  async function loadBackendTickets() {
    try {
      const data = await fetchTickets();
      if (Array.isArray(data) && data.length > 0) {
        tickets = data;
      }
    } catch {
      // Dynamic fallback
    }
  }

  onMount(() => {
    loadBackendTickets();
  });

  $: filteredTickets = tickets.filter(t => {
    const matchesSearch = t.title.toLowerCase().includes(searchQuery.toLowerCase()) || 
                          t.id.toLowerCase().includes(searchQuery.toLowerCase()) ||
                          t.description.toLowerCase().includes(searchQuery.toLowerCase());
    const matchesStatus = selectedStatusFilter === 'all' || t.status === selectedStatusFilter;
    const matchesSeverity = selectedSeverityFilter === 'all' || t.severity === selectedSeverityFilter;
    return matchesSearch && matchesStatus && matchesSeverity;
  });

  function openTicketModal(ticket: Ticket) {
    selectedTicket = ticket;
    isModalOpen = true;
  }
</script>

<div class="space-y-6">
  <!-- Page Header -->
  <div class="flex items-center justify-between">
    <div>
      <h1 class="text-xl font-bold text-white tracking-tight">All Issues & Defects</h1>
      <p class="text-xs text-slate-400 mt-0.5">Filter, search, and manage all bug reports across repositories.</p>
    </div>
    <a href="/submit" class="bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-bold px-4 py-2 rounded-xl transition">
      + Report New Issue
    </a>
  </div>

  <!-- Filters & Search Bar -->
  <div class="flex flex-wrap items-center justify-between gap-4 bg-[#0e1320] border border-white/[0.08] p-4 rounded-2xl">
    <div class="flex items-center space-x-3 flex-1 min-w-[280px]">
      <input 
        type="text" 
        bind:value={searchQuery}
        placeholder="Filter issues by title, ID, or stack trace..."
        class="w-full bg-[#111726] border border-white/10 rounded-xl px-4 py-2 text-xs text-white placeholder-slate-500 focus:outline-none focus:border-indigo-500 transition"
      />
    </div>

    <div class="flex items-center space-x-3 text-xs">
      <select 
        bind:value={selectedStatusFilter}
        class="bg-[#111726] border border-white/10 rounded-xl px-3 py-2 text-slate-200 focus:outline-none focus:border-indigo-500 transition"
      >
        <option value="all">All Statuses</option>
        <option value="new">New</option>
        <option value="triaging">Semantic Triaging</option>
        <option value="sandbox_running">Sandbox Running</option>
        <option value="diagnosed">AI Diagnosed</option>
        <option value="resolved">Resolved</option>
      </select>

      <select 
        bind:value={selectedSeverityFilter}
        class="bg-[#111726] border border-white/10 rounded-xl px-3 py-2 text-slate-200 focus:outline-none focus:border-indigo-500 transition"
      >
        <option value="all">All Severities</option>
        <option value="critical">Critical</option>
        <option value="high">High</option>
        <option value="medium">Medium</option>
        <option value="low">Low</option>
      </select>
    </div>
  </div>

  <!-- Issues List Table -->
  <div class="bg-[#0e1320] border border-white/[0.08] rounded-2xl overflow-hidden shadow-xl">
    <div class="overflow-x-auto">
      <table class="w-full text-left text-xs text-slate-300">
        <thead class="bg-[#090d16] text-[10px] uppercase font-bold text-slate-400 border-b border-white/[0.08]">
          <tr>
            <th class="py-3.5 px-4">Ticket ID</th>
            <th class="py-3.5 px-4">Title & Description</th>
            <th class="py-3.5 px-4">Severity</th>
            <th class="py-3.5 px-4">Repository / Branch</th>
            <th class="py-3.5 px-4">Status</th>
            <th class="py-3.5 px-4 text-right">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-white/[0.06]">
          {#each filteredTickets as t}
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
            <tr 
              class="hover:bg-white/[0.03] transition cursor-pointer"
              on:click={() => openTicketModal(t)}
            >
              <td class="py-4 px-4 font-mono font-bold text-indigo-400">{t.id}</td>
              <td class="py-4 px-4 max-w-md">
                <div class="font-bold text-white line-clamp-1">{t.title}</div>
                <div class="text-[11px] text-slate-400 line-clamp-1 mt-0.5">{t.description}</div>
              </td>
              <td class="py-4 px-4">
                <span class={`text-[10px] font-bold px-2 py-0.5 rounded-full border uppercase ${
                  t.severity === 'critical' ? 'bg-rose-500/10 text-rose-400 border-rose-500/20' :
                  t.severity === 'high' ? 'bg-amber-500/10 text-amber-400 border-amber-500/20' :
                  t.severity === 'medium' ? 'bg-sky-500/10 text-sky-400 border-sky-500/20' :
                  'bg-slate-500/10 text-slate-400 border-slate-500/20'
                }`}>
                  {t.severity}
                </span>
              </td>
              <td class="py-4 px-4 font-mono text-[11px] text-slate-400">{t.repo_branch_url}</td>
              <td class="py-4 px-4">
                <span class={`text-[10px] font-bold px-2 py-0.5 rounded-full border ${
                  t.status === 'diagnosed' ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20' :
                  t.status === 'sandbox_running' ? 'bg-amber-500/10 text-amber-400 border-amber-500/20' :
                  'bg-indigo-500/10 text-indigo-400 border-indigo-500/20'
                }`}>
                  {t.status}
                </span>
              </td>
              <td class="py-4 px-4 text-right">
                <button class="text-xs font-semibold text-indigo-400 hover:text-indigo-300">View Details →</button>
              </td>
            </tr>
          {/each}

          {#if filteredTickets.length === 0}
            <tr>
              <td colspan="6" class="py-12 text-center text-slate-500 text-xs">
                No issues match your current search and filter criteria.
              </td>
            </tr>
          {/if}
        </tbody>
      </table>
    </div>
  </div>

  <!-- Ticket Detail Modal -->
  <TicketModal ticket={selectedTicket} isOpen={isModalOpen} onClose={() => isModalOpen = false} />
</div>
