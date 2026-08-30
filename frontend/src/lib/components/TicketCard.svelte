<script lang="ts">
  import type { Ticket } from '$lib/types';

  export let ticket: Ticket;
  export let onClick: (ticket: Ticket) => void = () => {};

  const severityStyles = {
    critical: 'bg-rose-500/10 text-rose-400 border-rose-500/20',
    high: 'bg-amber-500/10 text-amber-400 border-amber-500/20',
    medium: 'bg-sky-500/10 text-sky-400 border-sky-500/20',
    low: 'bg-slate-500/10 text-slate-400 border-slate-500/20'
  };
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div 
  class="bg-[#111726] border border-white/5 hover:border-indigo-500/40 rounded-lg p-3.5 cursor-pointer transition-all duration-150 hover:-translate-y-0.5 shadow-sm hover:shadow-indigo-500/10 group"
  on:click={() => onClick(ticket)}
>
  <div class="flex items-center justify-between gap-2 mb-2">
    <span class="text-[11px] font-mono text-slate-400 group-hover:text-indigo-300 transition-colors">{ticket.id}</span>
    <span class={`text-[10px] font-medium px-2 py-0.5 rounded-full border uppercase tracking-wider ${severityStyles[ticket.severity] || severityStyles.medium}`}>
      {ticket.severity}
    </span>
  </div>

  <h3 class="text-xs font-semibold text-slate-200 line-clamp-2 leading-snug mb-2 group-hover:text-white transition-colors">
    {ticket.title}
  </h3>

  <p class="text-[11px] text-slate-400 line-clamp-2 leading-relaxed mb-3">
    {ticket.description}
  </p>

  {#if ticket.diagnosis}
    <div class="bg-indigo-500/10 border border-indigo-500/20 rounded p-2 text-[11px] text-indigo-300 flex items-center space-x-1.5 font-sans">
      <span class="text-indigo-400 text-xs">⚡</span>
      <span class="truncate text-[10px] font-medium">{ticket.diagnosis.root_cause}</span>
    </div>
  {/if}
</div>
