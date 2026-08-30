<script lang="ts">
  const metrics = [
    { label: 'Auto-Resolution Rate', value: '38.7%', change: '+4.2%', color: 'text-emerald-400' },
    { label: 'Mean Time to Resolve', value: '2.4 hrs', change: '-42%', color: 'text-emerald-400' },
    { label: 'Duplicate Rate Reduced', value: '64.1%', change: '+12%', color: 'text-indigo-400' },
    { label: 'Developer Hours Saved', value: '142 hrs', change: '+38 hrs', color: 'text-purple-400' }
  ];
</script>

<div class="space-y-6">
  <div>
    <h1 class="text-xl font-bold text-white tracking-tight">Defect Insights & Analytics</h1>
    <p class="text-xs text-slate-400 mt-0.5">Performance telemetry, auto-resolution rates, and triage metrics.</p>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
    {#each metrics as m}
      <div class="bg-[#0e1320] border border-white/[0.08] p-5 rounded-2xl space-y-2 shadow-xl">
        <span class="text-xs text-slate-400 font-medium">{m.label}</span>
        <div class="text-2xl font-bold text-white">{m.value}</div>
        <div class={`text-xs font-semibold ${m.color}`}>{m.change} <span class="text-slate-500 font-normal">this month</span></div>
      </div>
    {/each}
  </div>
</div>
