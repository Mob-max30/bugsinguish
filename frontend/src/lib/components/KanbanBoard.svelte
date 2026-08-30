<script lang="ts">
  import type { Ticket, TicketStatus } from '$lib/types';
  import TicketCard from './TicketCard.svelte';

  export let tickets: Ticket[] = [];
  export let onSelectTicket: (ticket: Ticket) => void = () => {};

  const columns: { id: TicketStatus; label: string; dot: string }[] = [
    { id: 'new', label: 'New Reports', dot: 'bg-slate-400' },
    { id: 'triaging', label: 'Semantic Triage', dot: 'bg-sky-400' },
    { id: 'sandbox_running', label: 'Sandbox Running', dot: 'bg-amber-400' },
    { id: 'diagnosed', label: 'AI Diagnosed (Draft PR)', dot: 'bg-indigo-400' },
    { id: 'resolved', label: 'Resolved & Closed', dot: 'bg-emerald-400' }
  ];

  function getTicketsByStatus(status: TicketStatus) {
    return tickets.filter(t => t.status === status);
  }
</script>

<div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 gap-3.5">
  {#each columns as col}
    {@const colTickets = getTicketsByStatus(col.id)}
    <div class="bg-[#0e1320] border border-white/[0.06] rounded-xl p-3 flex flex-col h-[640px]">
      <!-- Column Header -->
      <div class="flex items-center justify-between pb-3 mb-3 border-b border-white/[0.06]">
        <div class="flex items-center space-x-2">
          <span class={`w-2 h-2 rounded-full ${col.dot}`}></span>
          <h3 class="text-xs font-semibold text-slate-300 tracking-tight">{col.label}</h3>
        </div>
        <span class="text-[11px] font-mono text-slate-400 bg-white/5 border border-white/10 px-2 py-0.5 rounded-full font-medium">
          {colTickets.length}
        </span>
      </div>

      <!-- Column Ticket Stack -->
      <div class="flex-1 overflow-y-auto space-y-2.5 pr-0.5">
        {#each colTickets as ticket (ticket.id)}
          <TicketCard {ticket} onClick={onSelectTicket} />
        {/each}

        {#if colTickets.length === 0}
          <div class="h-28 border border-dashed border-white/10 rounded-lg flex items-center justify-center text-[11px] text-slate-500 font-medium">
            No tickets
          </div>
        {/if}
      </div>
    </div>
  {/each}
</div>
