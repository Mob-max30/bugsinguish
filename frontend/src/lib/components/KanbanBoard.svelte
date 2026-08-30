<script lang="ts">
  import type { Ticket, TicketStatus } from '$lib/types';
  import TicketCard from './TicketCard.svelte';

  export let tickets: Ticket[] = [];
  export let onSelectTicket: (ticket: Ticket) => void = () => {};

  const columns: { id: TicketStatus; step: string; label: string; desc: string; dot: string }[] = [
    { id: 'new', step: 'Step 1', label: 'New Bug Reports', desc: 'Submitted by team', dot: 'bg-slate-400' },
    { id: 'triaging', step: 'Step 2', label: 'AI Duplicate Check', desc: 'Scanning past tickets', dot: 'bg-sky-400' },
    { id: 'sandbox_running', step: 'Step 3', label: 'Auto-Test Sandbox', desc: 'Reproducing crash live', dot: 'bg-amber-400' },
    { id: 'diagnosed', step: 'Step 4', label: 'AI Solution Ready', desc: '1-Click Fix prepared', dot: 'bg-indigo-400' },
    { id: 'resolved', step: 'Step 5', label: 'Fixed & Closed', desc: 'Merged & verified', dot: 'bg-emerald-400' }
  ];

  function getTicketsByStatus(status: TicketStatus) {
    return tickets.filter(t => t.status === status);
  }
</script>

<div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 gap-3.5">
  {#each columns as col}
    {@const colTickets = getTicketsByStatus(col.id)}
    <div class="bg-[#0e1320] border border-white/[0.08] rounded-xl p-3 flex flex-col h-[660px]">
      <!-- Column Header with Plain-English Stepper -->
      <div class="pb-3 mb-3 border-b border-white/[0.08]">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-2">
            <span class={`w-2 h-2 rounded-full ${col.dot}`}></span>
            <span class="text-[10px] uppercase tracking-wider font-bold text-slate-400">{col.step}</span>
          </div>
          <span class="text-[11px] font-mono text-slate-300 bg-white/5 border border-white/10 px-2 py-0.5 rounded-full font-bold">
            {colTickets.length}
          </span>
        </div>
        <h3 class="text-xs font-bold text-slate-100 mt-1">{col.label}</h3>
        <p class="text-[10px] text-slate-400">{col.desc}</p>
      </div>

      <!-- Column Ticket Stack -->
      <div class="flex-1 overflow-y-auto space-y-3 pr-0.5">
        {#each colTickets as ticket (ticket.id)}
          <TicketCard {ticket} onClick={onSelectTicket} />
        {/each}

        {#if colTickets.length === 0}
          <div class="h-28 border border-dashed border-white/10 rounded-xl flex items-center justify-center text-[11px] text-slate-500 font-medium">
            No tickets in this stage
          </div>
        {/if}
      </div>
    </div>
  {/each}
</div>
