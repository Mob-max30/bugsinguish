<script lang="ts">
  import type { Ticket } from '$lib/types';

  export let ticket: Ticket;
  export let onClick: (ticket: Ticket) => void = () => {};
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div 
  class="bg-slate-900 border border-slate-800 hover:border-indigo-500/50 rounded-lg p-4 cursor-pointer transition shadow-md hover:shadow-indigo-500/5 group"
  on:click={() => onClick(ticket)}
>
  <div class="flex items-center justify-between mb-2">
    <span class="text-xs font-mono text-slate-400 group-hover:text-indigo-400 transition">{ticket.id}</span>
    <span class={`text-[10px] font-bold px-2 py-0.5 rounded uppercase tracking-wider ${
      ticket.severity === 'critical' ? 'bg-rose-950 text-rose-400 border border-rose-800/60' :
      ticket.severity === 'high' ? 'bg-amber-950 text-amber-400 border border-amber-800/60' :
      ticket.severity === 'medium' ? 'bg-blue-950 text-blue-400 border border-blue-800/60' :
      'bg-slate-800 text-slate-400'
    }`}>
      {ticket.severity}
    </span>
  </div>

  <h3 class="text-sm font-semibold text-slate-200 line-clamp-2 mb-2 group-hover:text-slate-100 transition">
    {ticket.title}
  </h3>

  <p class="text-xs text-slate-400 line-clamp-2 mb-3">
    {ticket.description}
  </p>

  {#if ticket.diagnosis}
    <div class="bg-indigo-950/40 border border-indigo-900/50 rounded p-2 text-[11px] text-indigo-300 flex items-center space-x-1.5">
      <span>⚡</span>
      <span class="truncate">{ticket.diagnosis.root_cause}</span>
    </div>
  {/if}
</div>
