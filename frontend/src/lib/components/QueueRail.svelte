<script lang="ts">
  import type { Ticket } from '$lib/types';

  export let tickets: Ticket[] = [];
  export let activeTicketId: string | null = null;
  export let onSelect: (ticket: Ticket) => void = () => {};

  const getStatusColor = (status: string) => {
    switch(status) {
      case 'resolved': return 'text-[var(--color-status-green)] border-[var(--color-status-green)]';
      case 'diagnosed': return 'text-[var(--color-brand-purple)] border-[var(--color-brand-purple)]';
      case 'sandbox_running': return 'text-[var(--color-status-yellow)] border-[var(--color-status-yellow)]';
      default: return 'text-[var(--color-status-blue)] border-[var(--color-status-blue)]';
    }
  };

  const getStatusBg = (status: string) => {
    switch(status) {
      case 'resolved': return 'bg-[var(--color-status-green)]/10';
      case 'diagnosed': return 'bg-[var(--color-brand-purple)]/10';
      case 'sandbox_running': return 'bg-[var(--color-status-yellow)]/10';
      default: return 'bg-[var(--color-status-blue)]/10';
    }
  };
</script>

<div class="bg-[var(--color-surface)] border border-[var(--color-border)] rounded-xl h-full flex flex-col">
  <div class="px-6 py-5 flex items-center justify-between border-b border-[var(--color-border)]">
    <h2 class="text-base font-semibold text-[var(--color-text-primary)]">Recent Issues</h2>
    <button class="text-sm text-[var(--color-brand-purple)] font-medium hover:underline">View All Issues &rarr;</button>
  </div>

  <div class="flex-1 overflow-y-auto px-4 py-2 space-y-1">
    {#each tickets as ticket (ticket.id)}
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <!-- svelte-ignore a11y-no-static-element-interactions -->
      <div 
        class={`flex items-center justify-between p-3 rounded-lg cursor-pointer transition-colors ${
          activeTicketId === ticket.id ? 'bg-[var(--color-surface-hover)]' : 'hover:bg-[var(--color-canvas)]'
        }`}
        on:click={() => onSelect(ticket)}
      >
        <div class="flex items-center space-x-4">
          <!-- Icon Tile -->
          <div class={`w-10 h-10 rounded-lg flex items-center justify-center shadow-inner ${getStatusBg(ticket.status)} ${getStatusColor(ticket.status)} border border-current/20`}>
            <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
          </div>
          
          <!-- Text -->
          <div>
            <h3 class="text-sm font-medium text-[var(--color-text-primary)] line-clamp-1">{ticket.title}</h3>
            <div class="flex items-center space-x-2 text-xs text-[var(--color-text-secondary)] mt-0.5">
              <span>#{ticket.id}</span>
              <span class="text-[var(--color-border)]">|</span>
              <span class="truncate max-w-[120px]">{ticket.severity}</span>
            </div>
          </div>
        </div>

        <div class="flex items-center space-x-4">
          <!-- Status Badge -->
          <div class={`px-2.5 py-1 text-xs font-medium rounded-full border ${getStatusColor(ticket.status)} ${getStatusBg(ticket.status)}`}>
            {ticket.status.replace('_', ' ')}
          </div>
          
          <!-- Timestamp -->
          <div class="text-xs text-[var(--color-text-muted)] w-16 text-right">
            2m ago
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>
