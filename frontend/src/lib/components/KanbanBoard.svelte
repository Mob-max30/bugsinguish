<script lang="ts">
  import type { Ticket, TicketStatus } from '$lib/types';
  import TicketCard from './TicketCard.svelte';

  export let tickets: Ticket[] = [];
  export let onSelectTicket: (ticket: Ticket) => void = () => {};

  const columns: { id: TicketStatus; label: string; color: string }[] = [
    { id: 'new', label: 'New Reports', color: 'border-slate-700' },
    { id: 'triaging', label: 'Semantic Triaging', color: 'border-blue-500' },
    { id: 'sandbox_running', label: 'Sandbox Running', color: 'border-amber-500' },
    { id: 'diagnosed', label: 'AI Diagnosed (Draft PR)', color: 'border-indigo-500' },
    { id: 'resolved', label: 'Resolved & Closed', color: 'border-emerald-500' }
  ];

  function getTicketsByStatus(status: TicketStatus) {
    return tickets.filter(t => t.status === status);
  }
</script>

<div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 gap-4">
  {#each columns as col}
    {@const colTickets = getTicketsByStatus(col.id)}
    <div class="bg-slate-950/60 border border-slate-800/80 rounded-xl p-3 flex flex-col h-[680px]">
      <!-- Column Header -->
      <div class={`flex items-center justify-between pb-3 mb-3 border-b-2 ${col.color}`}>
        <h3 class="text-xs font-bold uppercase tracking-wider text-slate-300">{col.label}</h3>
        <span class="text-xs font-mono bg-slate-800 text-slate-400 px-2 py-0.5 rounded-full font-semibold">
          {colTickets.length}
        </span>
      </div>

      <!-- Column Ticket Stack -->
      <div class="flex-1 overflow-y-auto space-y-3 pr-1 scrollbar-thin scrollbar-thumb-slate-800">
        {#each colTickets as ticket (ticket.id)}
          <TicketCard {ticket} onClick={onSelectTicket} />
        {/each}

        {#if colTickets.length === 0}
          <div class="h-24 border border-dashed border-slate-800/60 rounded-lg flex items-center justify-center text-xs text-slate-600 font-medium">
            Empty Stage
          </div>
        {/if}
      </div>
    </div>
  {/each}
</div>
