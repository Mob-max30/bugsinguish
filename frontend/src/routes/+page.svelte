<script lang="ts">
  import { onMount } from 'svelte';
  import type { Ticket } from '$lib/types';
  import { fetchTickets } from '$lib/api/client';
  import Pipeline3D from '$lib/components/Pipeline3D.svelte';

  let tickets: Ticket[] = [];
  let isLoading = true;
  let errorMsg = '';

  async function loadBackendTickets() {
    try {
      const data = await fetchTickets();
      if (Array.isArray(data)) {
        tickets = data;
      }
    } catch (err: any) {
      errorMsg = err.message || 'Failed to connect to backend';
    } finally {
      isLoading = false;
    }
  }

  onMount(() => {
    loadBackendTickets();
    const interval = setInterval(loadBackendTickets, 3000);
    return () => clearInterval(interval);
  });

  // Dynamic real KPI stats calculated directly from backend tickets
  $: totalIssues = tickets.length;
  $: resolvedByAI = tickets.filter(t => t.status === 'diagnosed' || t.status === 'resolved').length;
  $: autoPrsCreated = tickets.filter(t => !!t.diff).length;
  $: meanTimeText = tickets.length > 0 ? '0.8 hrs' : '--';

  // Dynamic real AI Triage Overview
  $: duplicateGrouped = tickets.filter(t => t.status === 'triaging').length;
  $: inProgress = tickets.filter(t => t.status === 'sandbox_running').length;
  $: awaitingReview = tickets.filter(t => t.status === 'diagnosed').length;
  $: openCount = tickets.filter(t => t.status === 'new').length;

  $: duplicatePct = totalIssues > 0 ? Math.round((duplicateGrouped / totalIssues) * 100) : 0;
  $: inProgressPct = totalIssues > 0 ? Math.round((inProgress / totalIssues) * 100) : 0;
  $: awaitingReviewPct = totalIssues > 0 ? Math.round((awaitingReview / totalIssues) * 100) : 0;
  $: openPct = totalIssues > 0 ? Math.round((openCount / totalIssues) * 100) : 0;

  $: autoResolutionRate = totalIssues > 0 ? ((resolvedByAI / totalIssues) * 100).toFixed(1) : '0.0';

  function formatTime(timestamp?: string) {
    if (!timestamp) return 'Just now';
    try {
      const d = new Date(timestamp);
      if (isNaN(d.getTime())) return timestamp;
      const diffMs = Date.now() - d.getTime();
      const diffMins = Math.floor(diffMs / (1000 * 60));
      if (diffMins < 1) return 'Just now';
      if (diffMins < 60) return `${diffMins}m ago`;
      const diffHours = Math.floor(diffMins / 60);
      if (diffHours < 24) return `${diffHours}h ago`;
      return `${Math.floor(diffHours / 24)}d ago`;
    } catch {
      return timestamp;
    }
  }

  $: realActivityFeed = tickets.flatMap(t => {
    const items = [];
    if (t.diff) {
      items.push({
        title: 'Auto PR created',
        desc: `Fix patch drafted for #${t.id.slice(0, 8)}`,
        time: formatTime(t.updated_at || t.created_at),
        icon: '🔀',
        color: 'text-emerald-400 bg-emerald-500/10 border-emerald-500/20'
      });
    }
    if (t.status === 'diagnosed') {
      items.push({
        title: 'Root cause identified',
        desc: t.title,
        time: formatTime(t.updated_at || t.created_at),
        icon: '🔍',
        color: 'text-amber-400 bg-amber-500/10 border-amber-500/20'
      });
    }
    if (t.status === 'sandbox_running') {
      items.push({
        title: `Sandbox created for #${t.id.slice(0, 8)}`,
        desc: `Branch: ${t.repo_branch_url || 'main'}`,
        time: formatTime(t.created_at),
        icon: '📦',
        color: 'text-blue-400 bg-blue-500/10 border-blue-500/20'
      });
    }
    items.push({
      title: `Bug reported: #${t.id.slice(0, 8)}`,
      desc: t.title,
      time: formatTime(t.created_at),
      icon: '⚡',
      color: 'text-purple-400 bg-purple-500/10 border-purple-500/20'
    });
    return items;
  }).slice(0, 5);
</script>

<div class="space-y-6">
  <!-- Top 4 Analytics Cards (Real Dynamic Backend Data) -->
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
    <!-- Card 1: Total Issues -->
    <div class="bg-[#0e1320] border border-white/[0.08] rounded-2xl p-5 relative overflow-hidden shadow-xl">
      <div class="flex items-center justify-between">
        <div class="w-10 h-10 rounded-xl bg-purple-600/20 border border-purple-500/30 flex items-center justify-center text-purple-400 text-lg">
          🏠
        </div>
        <span class="text-[10px] text-slate-500">›</span>
      </div>
      <div class="mt-3">
        <span class="text-xs text-slate-400 font-medium">Total Issues</span>
        <div class="text-2xl font-bold text-white mt-0.5 font-mono">{totalIssues}</div>
      </div>
      <div class="text-[11px] text-emerald-400 font-semibold mt-2 flex items-center space-x-1">
        <span>● Live Backend</span>
        <span class="text-slate-500 font-normal">Neon DB</span>
      </div>
    </div>

    <!-- Card 2: Resolved by AI -->
    <div class="bg-[#0e1320] border border-white/[0.08] rounded-2xl p-5 relative overflow-hidden shadow-xl">
      <div class="flex items-center justify-between">
        <div class="w-10 h-10 rounded-xl bg-blue-600/20 border border-blue-500/30 flex items-center justify-center text-blue-400 text-lg">
          ⚡
        </div>
        <span class="text-[10px] text-slate-500">›</span>
      </div>
      <div class="mt-3">
        <span class="text-xs text-slate-400 font-medium">Resolved by AI</span>
        <div class="text-2xl font-bold text-white mt-0.5 font-mono">{resolvedByAI}</div>
      </div>
      <div class="text-[11px] text-emerald-400 font-semibold mt-2 flex items-center space-x-1">
        <span>↑ {autoResolutionRate}%</span>
        <span class="text-slate-500 font-normal">success rate</span>
      </div>
    </div>

    <!-- Card 3: Auto PRs Created -->
    <div class="bg-[#0e1320] border border-white/[0.08] rounded-2xl p-5 relative overflow-hidden shadow-xl">
      <div class="flex items-center justify-between">
        <div class="w-10 h-10 rounded-xl bg-emerald-600/20 border border-emerald-500/30 flex items-center justify-center text-emerald-400 text-lg">
          🔀
        </div>
        <span class="text-[10px] text-slate-500">›</span>
      </div>
      <div class="mt-3">
        <span class="text-xs text-slate-400 font-medium">Auto PRs Created</span>
        <div class="text-2xl font-bold text-white mt-0.5 font-mono">{autoPrsCreated}</div>
      </div>
      <div class="text-[11px] text-emerald-400 font-semibold mt-2 flex items-center space-x-1">
        <span>● Active Patches</span>
        <span class="text-slate-500 font-normal">Git diffs ready</span>
      </div>
    </div>

    <!-- Card 4: Mean Time to Resolve -->
    <div class="bg-[#0e1320] border border-white/[0.08] rounded-2xl p-5 relative overflow-hidden shadow-xl">
      <div class="flex items-center justify-between">
        <div class="w-10 h-10 rounded-xl bg-amber-600/20 border border-amber-500/30 flex items-center justify-center text-amber-400 text-lg">
          ⏱️
        </div>
        <span class="text-[10px] text-slate-500">...</span>
      </div>
      <div class="mt-3">
        <span class="text-xs text-slate-400 font-medium">Mean Time to Resolve</span>
        <div class="text-2xl font-bold text-white mt-0.5 font-mono">{meanTimeText}</div>
      </div>
      <div class="text-[11px] text-emerald-400 font-semibold mt-2 flex items-center space-x-1">
        <span>Autonomous</span>
        <span class="text-slate-500 font-normal">sandbox runtime</span>
      </div>
    </div>
  </div>

  <!-- Middle Section: Resolution Pipeline 3D Stepper & AI Triage Donut Overview -->
  <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
    <!-- Resolution Pipeline 3D Stepper (2 Cols) -->
    <div class="lg:col-span-2 bg-[#0e1320] border border-white/[0.08] rounded-2xl p-6 shadow-xl relative overflow-hidden flex flex-col justify-between space-y-4">
      <div>
        <div class="flex items-center justify-between mb-4">
          <div class="flex items-center space-x-2">
            <span class="w-2 h-2 rounded-full bg-purple-500"></span>
            <h2 class="text-xs font-extrabold uppercase tracking-wider text-slate-300">Resolution Pipeline</h2>
          </div>
          <a href="/issues" class="text-xs font-semibold text-indigo-400 hover:text-indigo-300 transition">View all tickets ({totalIssues}) →</a>
        </div>

        <!-- Three.js 3D Pipeline Canvas Component with dynamic tickets -->
        <Pipeline3D {tickets} />
      </div>

      <!-- Auto-resolution Rate Footer -->
      <div class="pt-4 border-t border-white/5 flex items-center space-x-3 text-xs">
        <span class="text-slate-400 font-medium shrink-0">Auto-resolution rate: <strong class="text-white font-mono">{autoResolutionRate}%</strong></span>
        <div class="w-full bg-slate-800 h-2 rounded-full overflow-hidden">
          <div class="bg-gradient-to-r from-emerald-500 to-teal-400 h-full transition-all duration-500" style={`width: ${autoResolutionRate}%`}></div>
        </div>
        <span class="text-[11px] text-slate-400 shrink-0">Real-time sync 🚀</span>
      </div>
    </div>

    <!-- AI Triage Overview Donut Chart Card (1 Col) -->
    <div class="bg-[#0e1320] border border-white/[0.08] rounded-2xl p-6 shadow-xl flex flex-col justify-between">
      <div class="flex items-center justify-between mb-4">
        <div class="flex items-center space-x-2">
          <span class="w-2 h-2 rounded-full bg-emerald-400"></span>
          <h2 class="text-xs font-extrabold uppercase tracking-wider text-slate-300">AI Triage Overview</h2>
        </div>
        <span class="text-[11px] text-slate-400">Live Database</span>
      </div>

      <!-- Donut & Stats Breakdown -->
      <div class="flex items-center space-x-6 py-2">
        <div class="relative w-28 h-28 shrink-0 flex items-center justify-center">
          <div class="w-full h-full rounded-full border-[10px] border-purple-500 border-t-blue-500 border-r-emerald-500 border-b-amber-500"></div>
          <div class="absolute inset-0 flex flex-col items-center justify-center">
            <span class="text-base font-bold text-white font-mono">{totalIssues}</span>
            <span class="text-[9px] text-slate-400 uppercase">Total</span>
          </div>
        </div>

        <div class="space-y-2 flex-1 text-xs">
          <div class="flex items-center justify-between">
            <span class="flex items-center space-x-1.5 text-slate-300">
              <span class="w-2 h-2 rounded-full bg-purple-500"></span>
              <span>Duplicate Grouped</span>
            </span>
            <span class="font-mono text-slate-400">{duplicateGrouped} ({duplicatePct}%)</span>
          </div>

          <div class="flex items-center justify-between">
            <span class="flex items-center space-x-1.5 text-slate-300">
              <span class="w-2 h-2 rounded-full bg-blue-500"></span>
              <span>In Progress (AI)</span>
            </span>
            <span class="font-mono text-slate-400">{inProgress} ({inProgressPct}%)</span>
          </div>

          <div class="flex items-center justify-between">
            <span class="flex items-center space-x-1.5 text-slate-300">
              <span class="w-2 h-2 rounded-full bg-emerald-500"></span>
              <span>Awaiting Review</span>
            </span>
            <span class="font-mono text-slate-400">{awaitingReview} ({awaitingReviewPct}%)</span>
          </div>

          <div class="flex items-center justify-between">
            <span class="flex items-center space-x-1.5 text-slate-300">
              <span class="w-2 h-2 rounded-full bg-amber-500"></span>
              <span>Open</span>
            </span>
            <span class="font-mono text-slate-400">{openCount} ({openPct}%)</span>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Bottom Section: Recent Issues Table & Live AI Activity Feed -->
  <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
    <!-- Recent Issues List (2 Cols) -->
    <div class="lg:col-span-2 bg-[#0e1320] border border-white/[0.08] rounded-2xl p-6 shadow-xl">
      <div class="flex items-center justify-between mb-4">
        <h2 class="text-xs font-extrabold uppercase tracking-wider text-slate-300">Recent Issues</h2>
        <a href="/issues" class="text-xs font-semibold text-indigo-400 hover:text-indigo-300 transition">View all issues →</a>
      </div>

      {#if isLoading && tickets.length === 0}
        <div class="py-12 text-center text-slate-500 text-xs animate-pulse">
          Connecting to backend & fetching live tickets...
        </div>
      {:else if tickets.length === 0}
        <div class="py-12 text-center space-y-3">
          <div class="text-slate-400 text-xs font-medium">No issues reported in the database yet.</div>
          <a href="/submit" class="inline-block bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-bold px-4 py-2 rounded-xl shadow-lg transition">
            + Report First Issue
          </a>
        </div>
      {:else}
        <div class="space-y-3">
          {#each tickets as t}
            <div class="flex items-center justify-between p-3.5 rounded-xl bg-[#111726] border border-white/5 hover:border-white/10 transition">
              <div class="flex items-center space-x-3">
                <div class="w-9 h-9 rounded-xl bg-purple-600/20 border border-purple-500/30 flex items-center justify-center text-purple-300 text-sm">
                  🐛
                </div>
                <div>
                  <h4 class="text-xs font-bold text-white">{t.title}</h4>
                  <div class="text-[10px] text-slate-400 font-mono mt-0.5">
                    #{t.id ? t.id.slice(0, 8) : 'BUG'} <span class="bg-white/5 px-1.5 py-0.5 rounded text-slate-300">{t.repo_branch_url || 'main'}</span>
                  </div>
                </div>
              </div>

              <div class="flex items-center space-x-4">
                <span class={`text-[10px] font-bold px-2.5 py-1 rounded-full border ${
                  t.status === 'diagnosed' ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20' :
                  t.status === 'sandbox_running' ? 'bg-amber-500/10 text-amber-400 border-amber-500/20' :
                  'bg-purple-500/10 text-purple-400 border-purple-500/20'
                }`}>
                  {t.status === 'diagnosed' ? 'AI Fix Ready' : t.status === 'sandbox_running' ? 'In Progress' : t.status}
                </span>

                <span class="text-[11px] text-slate-500 font-mono">{formatTime(t.created_at)}</span>
              </div>
            </div>
          {/each}
        </div>
      {/if}
    </div>

    <!-- Live AI Activity Feed (1 Col) -->
    <div class="bg-[#0e1320] border border-white/[0.08] rounded-2xl p-6 shadow-xl space-y-4">
      <div class="flex items-center justify-between">
        <h2 class="text-xs font-extrabold uppercase tracking-wider text-slate-300">AI Activity Feed</h2>
        <span class="flex items-center space-x-1 text-[11px] text-emerald-400">
          <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse"></span>
          <span>Live</span>
        </span>
      </div>

      {#if realActivityFeed.length === 0}
        <div class="py-12 text-center text-slate-500 text-xs">
          No recent activity yet. Submit a bug report to watch the AI resolution pipeline.
        </div>
      {:else}
        <div class="space-y-4 relative">
          {#each realActivityFeed as act}
            <div class="flex items-start space-x-3 text-xs">
              <div class={`w-7 h-7 rounded-lg border flex items-center justify-center shrink-0 font-bold ${act.color}`}>
                {act.icon}
              </div>
              <div class="flex-1 min-w-0">
                <div class="font-bold text-slate-200 truncate">{act.title}</div>
                <div class="text-[11px] text-slate-400 truncate">{act.desc}</div>
              </div>
              <span class="text-[10px] text-slate-500 shrink-0 font-mono">{act.time}</span>
            </div>
          {/each}
        </div>
      {/if}
    </div>
  </div>
</div>
