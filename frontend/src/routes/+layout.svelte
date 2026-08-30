<script lang="ts">
  import '../app.css';
  import { page } from '$app/stores';

  $: currentPath = $page.url.pathname;

  const navItems = [
    { label: 'Dashboard', icon: '📊', href: '/' },
    { label: 'Issues', icon: '📋', href: '/issues' },
    { label: 'AI Triage', icon: '⚡', href: '/triage', badge: 'New' },
    { label: 'Sandboxes', icon: '📦', href: '/sandboxes' },
    { label: 'Pull Requests', icon: '🔀', href: '/pull-requests' }
  ];
</script>

<div class="min-h-screen flex bg-[#070a12] text-slate-100 font-sans antialiased">
  <!-- Left Sidebar -->
  <aside class="w-64 border-r border-white/[0.08] bg-[#090d16] flex flex-col justify-between p-4 shrink-0 sticky top-0 h-screen">
    <div>
      <!-- Brand Logo Header -->
      <a href="/" class="flex items-center space-x-3 px-2 py-3 mb-6">
        <div class="w-9 h-9 rounded-xl bg-gradient-to-br from-indigo-500 via-purple-600 to-indigo-800 flex items-center justify-center text-white font-black text-lg shadow-lg shadow-indigo-500/20">
          ⚙️
        </div>
        <div>
          <div class="text-sm font-bold tracking-tight text-white flex items-center space-x-1.5">
            <span>Bugsinguish</span>
          </div>
          <span class="text-[10px] text-slate-400 font-medium">AI-Native Defect Resolution</span>
        </div>
      </a>

      <!-- Navigation Menu -->
      <nav class="space-y-1">
        {#each navItems as item}
          {@const isActive = currentPath === item.href}
          <a 
            href={item.href}
            class={`flex items-center justify-between px-3.5 py-2.5 rounded-xl text-xs font-semibold transition ${
              isActive 
                ? 'bg-gradient-to-r from-indigo-600/90 to-purple-600/90 text-white shadow-lg shadow-indigo-500/20' 
                : 'text-slate-400 hover:text-slate-100 hover:bg-white/[0.04]'
            }`}
          >
            <div class="flex items-center space-x-3">
              <span class="text-sm">{item.icon}</span>
              <span>{item.label}</span>
            </div>
            {#if item.badge}
              <span class="text-[9px] font-extrabold uppercase px-1.5 py-0.5 rounded-full bg-indigo-500/20 text-indigo-300 border border-indigo-500/30">
                {item.badge}
              </span>
            {/if}
          </a>
        {/each}
      </nav>
    </div>

    <!-- System Status Footer -->
    <div class="pt-4 border-t border-white/[0.08]">
      <div class="flex items-center space-x-2.5 px-3 py-2.5 rounded-xl bg-white/[0.03] border border-white/[0.06] text-xs">
        <span class="w-2 h-2 rounded-full bg-emerald-400 animate-pulse"></span>
        <span class="font-medium text-[11px] text-slate-300">AI Resolution Engine Online</span>
      </div>
    </div>
  </aside>

  <!-- Main Content Area -->
  <div class="flex-1 flex flex-col min-w-0">
    <!-- Top Command Header Bar -->
    <header class="h-16 border-b border-white/[0.08] bg-[#090d16]/80 backdrop-blur-md px-8 flex items-center justify-between sticky top-0 z-30">
      <div>
        <h1 class="text-base font-bold text-white flex items-center space-x-2">
          <span>Defect Resolution Engine</span>
          <span class="text-base">⚡</span>
        </h1>
        <p class="text-[11px] text-slate-400">Autonomous bug triage, ephemeral sandboxes, and AI-generated PRs.</p>
      </div>

      <div class="flex items-center space-x-4">
        <!-- Search Bar -->
        <div class="relative w-64">
          <input 
            type="text" 
            placeholder="Search issues, PRs, crash logs... ⌘K"
            class="w-full bg-[#0d121f] border border-white/[0.1] rounded-xl px-3.5 py-1.5 text-xs text-slate-200 placeholder-slate-500 focus:outline-none focus:border-indigo-500 transition"
          />
        </div>

        <!-- Notification Bell -->
        <div class="relative p-2 rounded-xl bg-[#0d121f] border border-white/[0.1] text-slate-300 text-xs">
          🔔
          <span class="absolute -top-1 -right-1 w-4 h-4 bg-indigo-500 text-[9px] font-bold text-white rounded-full flex items-center justify-center">3</span>
        </div>

        <!-- New Issue CTA -->
        <a 
          href="/submit" 
          class="bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-500 hover:to-purple-500 text-white text-xs font-bold px-4 py-2 rounded-xl shadow-lg shadow-indigo-500/25 transition flex items-center space-x-1.5"
        >
          <span>+ New Issue</span>
        </a>
      </div>
    </header>

    <!-- Page Body -->
    <main class="flex-1 p-8 overflow-y-auto space-y-6">
      <slot />
    </main>
  </div>
</div>
