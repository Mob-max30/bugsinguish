<script lang="ts">
  const integrations = [
    { name: 'Google Gemini API', desc: 'Powers 768-dim text embeddings and Gemini 1.5 Pro root cause analysis.', status: 'Connected', icon: '✨' },
    { name: 'Neon Serverless Postgres', desc: 'pgvector similarity search and relational ticket state storage.', status: 'Connected', icon: '🐘' },
    { name: 'Docker Engine SDK', desc: 'Spawns short-lived isolated test reproduction containers.', status: 'Connected', icon: '🐳' },
    { name: 'GitHub REST & GraphQL API', desc: 'Opens draft PRs, inspects branch files, and merges approved patches.', status: 'Connected', icon: '🐙' }
  ];
</script>

<div class="space-y-6">
  <div>
    <h1 class="text-xl font-bold text-white tracking-tight">Active Platform Integrations</h1>
    <p class="text-xs text-slate-400 mt-0.5">Connected services powering the AI-native defect resolution pipeline.</p>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
    {#each integrations as i}
      <div class="bg-[#0e1320] border border-white/[0.08] p-5 rounded-2xl space-y-3 shadow-xl">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <span class="text-xl">{i.icon}</span>
            <h3 class="text-sm font-bold text-white">{i.name}</h3>
          </div>
          <span class="text-[10px] font-bold px-2.5 py-0.5 rounded-full bg-emerald-500/10 text-emerald-400 border border-emerald-500/20">{i.status}</span>
        </div>
        <p class="text-xs text-slate-400 leading-relaxed">{i.desc}</p>
      </div>
    {/each}
  </div>
</div>
