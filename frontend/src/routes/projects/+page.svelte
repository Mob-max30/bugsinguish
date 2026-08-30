<script lang="ts">
  const projects = [
    { name: 'iOS App', repo: 'Mob-max30/ios-client', branch: 'main', openBugs: 12, resolved: 148, health: '98%' },
    { name: 'Sync Service', repo: 'Mob-max30/sync-backend', branch: 'dev', openBugs: 4, resolved: 89, health: '95%' },
    { name: 'Web Checkout', repo: 'Mob-max30/web-checkout', branch: 'main', openBugs: 2, resolved: 64, health: '99%' },
    { name: 'dummy_repo', repo: 'Mob-max30/bugsinguish/sandbox/dummy_repo', branch: 'main', openBugs: 1, resolved: 15, health: '100%' }
  ];
</script>

<div class="space-y-6">
  <div class="flex items-center justify-between">
    <div>
      <h1 class="text-xl font-bold text-white tracking-tight">Connected Projects & Repositories</h1>
      <p class="text-xs text-slate-400 mt-0.5">Repositories monitored by Bugsinguish for automated sandbox reproduction.</p>
    </div>
    <button class="bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-bold px-4 py-2 rounded-xl transition">+ Connect Repo</button>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
    {#each projects as p}
      <div class="bg-[#0e1320] border border-white/[0.08] p-5 rounded-2xl space-y-3 shadow-xl">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <span class="text-lg">📁</span>
            <div>
              <h3 class="text-sm font-bold text-white">{p.name}</h3>
              <div class="text-[11px] text-slate-400 font-mono">{p.repo}</div>
            </div>
          </div>
          <span class="text-xs font-bold text-emerald-400">{p.health} Health</span>
        </div>

        <div class="flex items-center justify-between text-xs text-slate-300 pt-2 border-t border-white/5 font-mono">
          <span>Branch: <strong class="text-indigo-300">{p.branch}</strong></span>
          <span>Open: <strong>{p.openBugs}</strong> | Fixed: <strong>{p.resolved}</strong></span>
        </div>
      </div>
    {/each}
  </div>
</div>
