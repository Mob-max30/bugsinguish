<script lang="ts">
  import type { Ticket } from '$lib/types';

  export let ticket: Ticket;
  export let onClick: (ticket: Ticket) => void = () => {};

  const severityLabels = {
    critical: { label: 'Urgent Fix Required', style: 'bg-rose-500/10 text-rose-400 border-rose-500/20' },
    high: { label: 'High Priority', style: 'bg-amber-500/10 text-amber-400 border-amber-500/20' },
    medium: { label: 'Medium Impact', style: 'bg-sky-500/10 text-sky-400 border-sky-500/20' },
    low: { label: 'Minor Issue', style: 'bg-slate-500/10 text-slate-400 border-slate-500/20' }
  };
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div 
  class="bg-[#111726] border border-white/[0.08] hover:border-indigo-500/50 rounded-xl p-4 cursor-pointer transition-all duration-200 hover:-translate-y-0.5 shadow-sm hover:shadow-indigo-500/10 group"
  on:click={() => onClick(ticket)}
>
  <!-- Card Top: ID & Severity -->
  <div class="flex items-center justify-between gap-2 mb-2.5">
    <span class="text-[11px] font-mono font-medium text-slate-400 group-hover:text-indigo-300 transition-colors">
      Ticket {ticket.id}
    </span>
    <span class={`text-[10px] font-semibold px-2.5 py-0.5 rounded-full border ${severityLabels[ticket.severity]?.style || severityLabels.medium.style}`}>
      {severityLabels[ticket.severity]?.label || 'Medium Impact'}
    </span>
  </div>

  <!-- Title -->
  <h3 class="text-xs font-bold text-slate-100 line-clamp-2 leading-snug mb-2 group-hover:text-indigo-200 transition-colors">
    {ticket.title}
  </h3>

  <!-- Description -->
  <p class="text-[11px] text-slate-400 line-clamp-2 leading-relaxed mb-3">
    {ticket.description}
  </p>

  <!-- AI Diagnosis Preview or Friendly Action -->
  {#if ticket.diagnosis}
    <div class="bg-indigo-950/40 border border-indigo-500/30 rounded-lg p-2.5 mb-3 space-y-1">
      <div class="flex items-center space-x-1.5 text-indigo-300 text-[10px] font-bold uppercase tracking-wider">
        <span>✨ AI Root Cause Identified</span>
      </div>
      <p class="text-[11px] text-slate-200 line-clamp-2 font-medium leading-tight">
        {ticket.diagnosis.root_cause}
      </p>
    </div>

    <button class="w-full bg-emerald-600 hover:bg-emerald-500 text-white text-xs font-bold py-1.5 px-3 rounded-lg shadow-md transition flex items-center justify-center space-x-1">
      <span>✓ Review & Approve 1-Click Fix</span>
    </button>
  {:else}
    <div class="text-[10px] text-slate-500 flex items-center justify-between pt-1 border-t border-white/5">
      <span>Click card to view details & logs</span>
      <span>→</span>
    </div>
  {/if}
</div>
