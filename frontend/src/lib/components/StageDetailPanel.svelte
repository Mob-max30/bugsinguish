<script lang="ts">
  import type { Ticket, TicketStatus } from '$lib/types';

  export let ticket: Ticket | null = null;
  export let activeStage: TicketStatus = 'new';
  export let liveLogs: string[] = [];
  export let isSandboxRunning: boolean = false;
</script>

<div class="bg-[var(--color-surface)] border border-[var(--color-border)] rounded-xl h-full flex flex-col overflow-hidden">
  <div class="px-6 py-5 flex items-center justify-between border-b border-[var(--color-border)]">
    <div class="flex items-center space-x-2">
      <h2 class="text-base font-semibold text-[var(--color-text-primary)]">AI Activity Feed</h2>
      <span class="flex items-center space-x-1.5 px-2 py-0.5 rounded-full bg-[var(--color-status-green)]/10 text-[var(--color-status-green)] text-xs font-medium border border-[var(--color-status-green)]/20">
        <span class="w-1.5 h-1.5 rounded-full bg-[var(--color-status-green)] animate-pulse"></span>
        <span>Live</span>
      </span>
    </div>
  </div>

  <div class="flex-1 overflow-y-auto px-6 py-4 space-y-6">
    <!-- STUBBED GLOBAL EVENTS to match the screenshot -->
    <div class="flex space-x-4">
      <div class="w-8 h-8 rounded-lg bg-[var(--color-text-muted)]/10 flex items-center justify-center shrink-0 border border-[var(--color-border)]">
        <svg class="w-4 h-4 text-[var(--color-text-secondary)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path></svg>
      </div>
      <div class="flex-1 pt-1">
        <p class="text-sm font-medium text-[var(--color-text-primary)]">AI grouped 5 similar issues into one</p>
        <p class="text-xs text-[var(--color-text-secondary)] mt-1">"Login failed on mobile" + 4 others</p>
      </div>
      <div class="text-xs text-[var(--color-text-muted)] pt-1">2m ago</div>
    </div>
    
    <div class="flex space-x-4">
      <div class="w-8 h-8 rounded-lg bg-[var(--color-brand-purple)]/10 flex items-center justify-center shrink-0 border border-[var(--color-brand-purple)]/20">
        <svg class="w-4 h-4 text-[var(--color-brand-purple)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path></svg>
      </div>
      <div class="flex-1 pt-1">
        <p class="text-sm font-medium text-[var(--color-text-primary)]">Sandbox created for <span class="text-[var(--color-brand-purple)]">#BUG-1248</span></p>
        <p class="text-xs text-[var(--color-text-secondary)] mt-1">Branch: fix/login-redirect</p>
      </div>
      <div class="text-xs text-[var(--color-text-muted)] pt-1">3m ago</div>
    </div>

    <!-- DYNAMIC LIVE LOGS (if applicable) -->
    {#if ticket && isSandboxRunning && liveLogs.length > 0}
      {#each liveLogs as log, i}
        <div class="flex space-x-4">
          <div class="w-8 h-8 rounded-lg bg-[var(--color-status-yellow)]/10 flex items-center justify-center shrink-0 border border-[var(--color-status-yellow)]/20">
            <svg class="w-4 h-4 text-[var(--color-status-yellow)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path></svg>
          </div>
          <div class="flex-1 pt-1">
            <p class="text-sm font-medium text-[var(--color-text-primary)]">Sandbox Output (Live)</p>
            <p class="text-xs text-[var(--color-text-secondary)] mt-1 font-mono">{log}</p>
          </div>
          <div class="text-xs text-[var(--color-text-muted)] pt-1">just now</div>
        </div>
      {/each}
    {/if}
    
    <div class="flex space-x-4">
      <div class="w-8 h-8 rounded-lg bg-[var(--color-status-green)]/10 flex items-center justify-center shrink-0 border border-[var(--color-status-green)]/20">
        <svg class="w-4 h-4 text-[var(--color-status-green)]" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
      </div>
      <div class="flex-1 pt-1">
        <p class="text-sm font-medium text-[var(--color-text-primary)]">Tests passed in sandbox</p>
        <p class="text-xs text-[var(--color-text-secondary)] mt-1">23/23 tests passed</p>
      </div>
      <div class="text-xs text-[var(--color-text-muted)] pt-1">6m ago</div>
    </div>

  </div>
</div>
