<script lang="ts">
  import '../app.css';
</script>

<div class="min-h-screen flex flex-col bg-[#0b0f19] text-slate-100 font-sans antialiased">
  <!-- Top Navigation Bar -->
  <header class="bg-[#0e1320]/90 backdrop-blur-md border-b border-white/[0.08] sticky top-0 z-40">
    <div class="max-w-7xl mx-auto px-6 h-14 flex items-center justify-between">
      <a href="/" class="flex items-center space-x-3 group">
        <div class="w-7 h-7 rounded-lg bg-indigo-600 flex items-center justify-center text-white font-black text-sm shadow-md shadow-indigo-500/20 group-hover:bg-indigo-500 transition">
          B
        </div>
        <div class="flex items-center space-x-2">
          <span class="text-sm font-bold tracking-tight text-white">Bugsinguish</span>
          <span class="text-[10px] font-mono text-indigo-400 bg-indigo-500/10 border border-indigo-500/20 px-1.5 py-0.5 rounded">v0.1</span>
        </div>
      </a>

      <nav class="flex items-center space-x-3">
        <a 
          href="/" 
          class="text-xs font-medium text-slate-300 hover:text-white px-3 py-1.5 rounded-md hover:bg-white/5 transition"
        >
          Kanban Board
        </a>
        <a 
          href="/submit" 
          class="text-xs font-semibold bg-indigo-600 hover:bg-indigo-500 text-white px-3.5 py-1.5 rounded-md shadow-sm transition"
        >
          + Report Bug
        </a>
      </nav>
    </div>
  </header>

  <!-- Main Content -->
  <main class="flex-1 max-w-7xl w-full mx-auto p-6">
    <slot />
  </main>

  <!-- Footer -->
  <footer class="border-t border-white/[0.06] bg-[#0b0f19] py-4 text-center text-[11px] text-slate-400">
    Bugsinguish — AI-Native Defect Resolution Engine · Team Pranav · Ujwal · Pavan
  </footer>
</div>
