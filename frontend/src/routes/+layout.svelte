<script lang="ts">
  import '../app.css';
</script>

<div class="min-h-screen flex flex-col bg-slate-950 text-slate-100">
  <!-- Top Navigation Bar -->
  <header class="bg-slate-900/80 backdrop-blur border-b border-slate-800 sticky top-0 z-40">
    <div class="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
      <a href="/" class="flex items-center space-x-3 group">
        <div class="w-9 h-9 rounded-lg bg-gradient-to-tr from-indigo-600 to-violet-500 flex items-center justify-center text-white font-black text-lg shadow-lg group-hover:scale-105 transition">
          B
        </div>
        <div>
          <span class="text-lg font-bold tracking-tight text-white">Bugsinguish</span>
          <span class="text-[10px] uppercase font-mono tracking-widest text-indigo-400 block -mt-1">Active Resolution Engine</span>
        </div>
      </a>

      <nav class="flex items-center space-x-4">
        <a 
          href="/" 
          class="text-sm font-medium text-slate-300 hover:text-white px-3 py-2 rounded-lg hover:bg-slate-800 transition"
        >
          Kanban Board
        </a>
        <a 
          href="/submit" 
          class="text-sm font-semibold bg-indigo-600 hover:bg-indigo-500 text-white px-4 py-2 rounded-lg shadow-md transition"
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
  <footer class="border-t border-slate-900 bg-slate-950 py-4 text-center text-xs text-slate-600">
    Bugsinguish — Developer Tool Reconstruction (Bugzilla CloneFest) · Team Pranav · Ujwal · Pavan
  </footer>
</div>
