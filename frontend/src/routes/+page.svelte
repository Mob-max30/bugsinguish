<script lang="ts">
  import { onMount } from 'svelte';
  import type { Ticket } from '$lib/types';
  import { fetchTickets } from '$lib/api/client';
  import Pipeline3D from '$lib/components/Pipeline3D.svelte';

  let tickets: Ticket[] = [
    {
      id: 'BUG-1248',
      title: 'Users unable to login on iOS app',
      description: 'Authentication tokens fail to sign on iOS devices with 401 response.',
      stack_trace: '',
      repo_branch_url: 'iOS App',
      severity: 'critical',
      status: 'diagnosed',
      diagnosis: {
        root_cause: 'Null pointer exception in AuthService.java:42',
        explanation: 'Tokens expired without refresh handler.',
        file: 'AuthService.java',
        diff: ''
      },
      created_at: '2m ago',
      updated_at: '2m ago'
    },
    {
      id: 'BUG-1247',
      title: 'Data not syncing across devices',
      description: 'Background worker pool exhausts connection timeout under load.',
      stack_trace: '',
      repo_branch_url: 'Sync Service',
      severity: 'high',
      status: 'sandbox_running',
      created_at: '15m ago',
      updated_at: '15m ago'
    },
    {
      id: 'BUG-1246',
      title: 'Payment gateway timeout error',
      description: 'CORS header missing on GET /checkout endpoint.',
      stack_trace: '',
      repo_branch_url: 'Web Checkout',
      severity: 'medium',
      status: 'triaging',
      created_at: '1h ago',
      updated_at: '1h ago'
    },
    {
      id: 'BUG-1245',
      title: 'UI glitch in dark mode',
      description: 'Contrast on column card headers drops below 3.0 ratio.',
      stack_trace: '',
      repo_branch_url: 'Web App',
      severity: 'low',
      status: 'new',
      created_at: '2h ago',
      updated_at: '2h ago'
    },
    {
      id: 'BUG-1244',
      title: 'Crash on clicking export PDF',
      description: 'Export engine buffer overflows when rendering large reports.',
      stack_trace: '',
      repo_branch_url: 'Reports',
      severity: 'critical',
      status: 'diagnosed',
      diagnosis: {
        root_cause: 'Buffer overflow in PDF exporter',
        explanation: 'Max page limit unhandled.',
        file: 'pdf_exporter.go',
        diff: ''
      },
      created_at: '3h ago',
      updated_at: '3h ago'
    }
  ];

  async function loadBackendTickets() {
    try {
      const data = await fetchTickets();
      if (Array.isArray(data) && data.length > 0) {
        tickets = data;
      }
    } catch {
      // Keep rich interactive dashboard state
    }
  }

  onMount(() => {
    loadBackendTickets();
  });

  const activityFeed = [
    { title: 'AI grouped 5 similar issues into one', desc: '"Login failed on mobile" + 4 others', time: '2m ago', icon: '⚡', color: 'text-purple-400 bg-purple-500/10 border-purple-500/20' },
    { title: 'Sandbox created for #BUG-1248', desc: 'Branch: fix/login-redirect', time: '3m ago', icon: '📦', color: 'text-blue-400 bg-blue-500/10 border-blue-500/20' },
    { title: 'Root cause identified', desc: 'Null pointer exception in AuthService.java:42', time: '4m ago', icon: '🔍', color: 'text-amber-400 bg-amber-500/10 border-amber-500/20' },
    { title: 'Auto PR created', desc: 'PR #512: Fix null pointer in login redirect', time: '5m ago', icon: '🔀', color: 'text-emerald-400 bg-emerald-500/10 border-emerald-500/20' },
    { title: 'Tests passed in sandbox', desc: '23/23 tests passed', time: '6m ago', icon: '✓', color: 'text-teal-400 bg-teal-500/10 border-teal-500/20' }
  ];
</script>

<div class="space-y-6">
  <!-- Top 4 Analytics Cards -->
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
        <div class="text-2xl font-bold text-white mt-0.5">1,248</div>
      </div>
      <div class="text-[11px] text-emerald-400 font-semibold mt-2 flex items-center space-x-1">
        <span>↑ 12.5%</span>
        <span class="text-slate-500 font-normal">from last week</span>
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
        <div class="text-2xl font-bold text-white mt-0.5">342</div>
      </div>
      <div class="text-[11px] text-emerald-400 font-semibold mt-2 flex items-center space-x-1">
        <span>↑ 28.4%</span>
        <span class="text-slate-500 font-normal">from last week</span>
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
        <div class="text-2xl font-bold text-white mt-0.5">187</div>
      </div>
      <div class="text-[11px] text-emerald-400 font-semibold mt-2 flex items-center space-x-1">
        <span>↑ 31.7%</span>
        <span class="text-slate-500 font-normal">from last week</span>
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
        <div class="text-2xl font-bold text-white mt-0.5">2.4 hrs</div>
      </div>
      <div class="text-[11px] text-emerald-400 font-semibold mt-2 flex items-center space-x-1">
        <span>↓ 42%</span>
        <span class="text-slate-500 font-normal">from last week</span>
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
          <button class="text-xs font-semibold text-indigo-400 hover:text-indigo-300 transition">View full pipeline →</button>
        </div>

        <!-- Three.js 3D Pipeline Canvas Component -->
        <Pipeline3D />
      </div>

      <!-- Auto-resolution Rate Footer -->
      <div class="pt-4 border-t border-white/5 flex items-center space-x-3 text-xs">
        <span class="text-slate-400 font-medium shrink-0">Auto-resolution rate: <strong class="text-white">38.7%</strong></span>
        <div class="w-full bg-slate-800 h-2 rounded-full overflow-hidden">
          <div class="bg-gradient-to-r from-emerald-500 to-teal-400 h-full w-[38.7%]"></div>
        </div>
        <span class="text-[11px] text-slate-400 shrink-0">Improving steadily! 🚀</span>
      </div>
    </div>

    <!-- AI Triage Overview Donut Chart Card (1 Col) -->
    <div class="bg-[#0e1320] border border-white/[0.08] rounded-2xl p-6 shadow-xl flex flex-col justify-between">
      <div class="flex items-center justify-between mb-4">
        <div class="flex items-center space-x-2">
          <span class="w-2 h-2 rounded-full bg-emerald-400"></span>
          <h2 class="text-xs font-extrabold uppercase tracking-wider text-slate-300">AI Triage Overview</h2>
        </div>
        <span class="text-[11px] text-slate-400">This Week ⌄</span>
      </div>

      <!-- Donut & Stats Breakdown -->
      <div class="flex items-center space-x-6 py-2">
        <div class="relative w-28 h-28 shrink-0 flex items-center justify-center">
          <div class="w-full h-full rounded-full border-[10px] border-purple-500 border-t-blue-500 border-r-emerald-500 border-b-amber-500"></div>
          <div class="absolute inset-0 flex flex-col items-center justify-center">
            <span class="text-base font-bold text-white">1,248</span>
            <span class="text-[9px] text-slate-400 uppercase">Total</span>
          </div>
        </div>

        <div class="space-y-2 flex-1 text-xs">
          <div class="flex items-center justify-between">
            <span class="flex items-center space-x-1.5 text-slate-300">
              <span class="w-2 h-2 rounded-full bg-purple-500"></span>
              <span>Duplicate Grouped</span>
            </span>
            <span class="font-mono text-slate-400">512 (41%)</span>
          </div>

          <div class="flex items-center justify-between">
            <span class="flex items-center space-x-1.5 text-slate-300">
              <span class="w-2 h-2 rounded-full bg-blue-500"></span>
              <span>In Progress (AI)</span>
            </span>
            <span class="font-mono text-slate-400">320 (26%)</span>
          </div>

          <div class="flex items-center justify-between">
            <span class="flex items-center space-x-1.5 text-slate-300">
              <span class="w-2 h-2 rounded-full bg-emerald-500"></span>
              <span>Awaiting Review</span>
            </span>
            <span class="font-mono text-slate-400">231 (18%)</span>
          </div>

          <div class="flex items-center justify-between">
            <span class="flex items-center space-x-1.5 text-slate-300">
              <span class="w-2 h-2 rounded-full bg-amber-500"></span>
              <span>Open</span>
            </span>
            <span class="font-mono text-slate-400">185 (15%)</span>
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

      <div class="space-y-3">
        {#each tickets as t}
          <div class="flex items-center justify-between p-3.5 rounded-xl bg-[#111726] border border-white/5 hover:border-white/10 transition">
            <div class="flex items-center space-x-3">
              <div class="w-9 h-9 rounded-xl bg-purple-600/20 border border-purple-500/30 flex items-center justify-center text-purple-300 text-sm">
                📱
              </div>
              <div>
                <h4 class="text-xs font-bold text-white">{t.title}</h4>
                <div class="text-[10px] text-slate-400 font-mono mt-0.5">
                  #{t.id} <span class="bg-white/5 px-1.5 py-0.5 rounded text-slate-300">{t.repo_branch_url || 'iOS App'}</span>
                </div>
              </div>
            </div>

            <div class="flex items-center space-x-4">
              <span class={`text-[10px] font-bold px-2.5 py-1 rounded-full border ${
                t.status === 'diagnosed' ? 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20' :
                t.status === 'sandbox_running' ? 'bg-amber-500/10 text-amber-400 border-amber-500/20' :
                'bg-purple-500/10 text-purple-400 border-purple-500/20'
              }`}>
                {t.status === 'diagnosed' ? 'AI Fix Ready' : t.status === 'sandbox_running' ? 'In Progress' : 'Duplicate'}
              </span>

              <span class="text-[11px] text-slate-500 font-mono">{t.created_at}</span>
            </div>
          </div>
        {/each}
      </div>
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

      <div class="space-y-4 relative">
        {#each activityFeed as act}
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
    </div>
  </div>
</div>
