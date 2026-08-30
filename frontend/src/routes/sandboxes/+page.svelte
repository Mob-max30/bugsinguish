<script lang="ts">
  let activeSandboxes = [
    { id: 'sb-92a10b', ticket: '#BUG-1248', repo: 'dummy_repo', image: 'python:3.9-slim', status: 'Running Test Suite', uptime: '14s', cpu: '12%', mem: '142 MB' },
    { id: 'sb-41b9c2', ticket: '#BUG-1247', repo: 'sync_service', image: 'golang:1.24-alpine', status: 'Capturing Stderr Logs', uptime: '45s', cpu: '24%', mem: '210 MB' }
  ];
</script>

<div class="space-y-6">
  <div>
    <h1 class="text-xl font-bold text-white tracking-tight">Ephemeral Docker Sandbox Manager</h1>
    <p class="text-xs text-slate-400 mt-0.5">Short-lived containers spawned on-demand via native Docker Engine SDK for bug reproduction.</p>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
    {#each activeSandboxes as sb}
      <div class="bg-[#0e1320] border border-white/[0.08] p-5 rounded-2xl space-y-4 shadow-xl">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-2">
            <span class="w-2.5 h-2.5 rounded-full bg-emerald-400 animate-pulse"></span>
            <span class="font-mono font-bold text-xs text-white">{sb.id}</span>
          </div>
          <span class="text-[10px] font-mono bg-white/5 border border-white/10 px-2 py-0.5 rounded text-indigo-300">{sb.ticket}</span>
        </div>

        <div class="space-y-1.5 text-xs text-slate-300 font-mono">
          <div class="flex justify-between"><span>Base Image:</span><span class="text-slate-400">{sb.image}</span></div>
          <div class="flex justify-between"><span>Repository Context:</span><span class="text-slate-400">{sb.repo}</span></div>
          <div class="flex justify-between"><span>CPU / Memory:</span><span class="text-amber-400">{sb.cpu} / {sb.mem}</span></div>
          <div class="flex justify-between"><span>Container Uptime:</span><span class="text-slate-400">{sb.uptime}</span></div>
        </div>

        <div class="pt-2 flex justify-end">
          <button class="bg-rose-500/10 hover:bg-rose-500/20 text-rose-300 border border-rose-500/30 text-xs font-bold px-3 py-1.5 rounded-xl transition">
            Terminate Container
          </button>
        </div>
      </div>
    {/each}
  </div>
</div>
