<script lang="ts">
  import type { Ticket } from '$lib/types';
  export let ticket: Ticket | null = null;
</script>

<div class="bg-[var(--color-surface)] border border-[var(--color-border)] rounded-xl p-6 h-full flex flex-col">
  <div class="flex items-center justify-between mb-6">
    <h2 class="text-base font-semibold text-[var(--color-text-primary)]">AI Triage Overview</h2>
    <div class="px-3 py-1 bg-[var(--color-canvas)] border border-[var(--color-border)] rounded-md text-xs text-[var(--color-text-secondary)] flex items-center space-x-2">
      <span>This Week</span>
      <svg class="w-3 h-3" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" /></svg>
    </div>
  </div>

  <div class="flex-1 flex items-center justify-between">
    <!-- Stubbed SVG Donut Chart -->
    <div class="relative w-36 h-36">
      <svg viewBox="0 0 36 36" class="w-full h-full transform -rotate-90">
        <!-- Open (Yellow) -->
        <path class="text-[var(--color-status-yellow)]" stroke-dasharray="15, 100" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831" fill="none" stroke="currentColor" stroke-width="4" stroke-dashoffset="0" />
        <!-- Awaiting Review (Green) -->
        <path class="text-[var(--color-status-green)]" stroke-dasharray="18, 100" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831" fill="none" stroke="currentColor" stroke-width="4" stroke-dashoffset="-15" />
        <!-- In Progress (Blue) -->
        <path class="text-[var(--color-status-blue)]" stroke-dasharray="26, 100" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831" fill="none" stroke="currentColor" stroke-width="4" stroke-dashoffset="-33" />
        <!-- Duplicate Grouped (Purple) -->
        <path class="text-[var(--color-brand-purple)]" stroke-dasharray="41, 100" d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831" fill="none" stroke="currentColor" stroke-width="4" stroke-dashoffset="-59" />
      </svg>
      <div class="absolute inset-0 flex flex-col items-center justify-center">
        <span class="text-xl font-bold text-[var(--color-text-primary)]">1,248</span>
        <span class="text-xs text-[var(--color-text-secondary)]">Total</span>
      </div>
    </div>

    <!-- Legend -->
    <div class="flex-1 ml-8 space-y-4">
      <div class="flex items-center justify-between text-sm">
        <div class="flex items-center space-x-2">
          <span class="w-2 h-2 rounded-full bg-[var(--color-brand-purple)]"></span>
          <span class="text-[var(--color-text-secondary)]">Duplicate Grouped</span>
        </div>
        <div class="flex items-center space-x-4">
          <span class="text-[var(--color-text-primary)] font-medium">512</span>
          <span class="text-[var(--color-text-muted)] w-8 text-right">41%</span>
        </div>
      </div>
      <div class="flex items-center justify-between text-sm">
        <div class="flex items-center space-x-2">
          <span class="w-2 h-2 rounded-full bg-[var(--color-status-blue)]"></span>
          <span class="text-[var(--color-text-secondary)]">In Progress (AI)</span>
        </div>
        <div class="flex items-center space-x-4">
          <span class="text-[var(--color-text-primary)] font-medium">320</span>
          <span class="text-[var(--color-text-muted)] w-8 text-right">26%</span>
        </div>
      </div>
      <div class="flex items-center justify-between text-sm">
        <div class="flex items-center space-x-2">
          <span class="w-2 h-2 rounded-full bg-[var(--color-status-green)]"></span>
          <span class="text-[var(--color-text-secondary)]">Awaiting Review</span>
        </div>
        <div class="flex items-center space-x-4">
          <span class="text-[var(--color-text-primary)] font-medium">231</span>
          <span class="text-[var(--color-text-muted)] w-8 text-right">18%</span>
        </div>
      </div>
      <div class="flex items-center justify-between text-sm">
        <div class="flex items-center space-x-2">
          <span class="w-2 h-2 rounded-full bg-[var(--color-status-yellow)]"></span>
          <span class="text-[var(--color-text-secondary)]">Open</span>
        </div>
        <div class="flex items-center space-x-4">
          <span class="text-[var(--color-text-primary)] font-medium">185</span>
          <span class="text-[var(--color-text-muted)] w-8 text-right">15%</span>
        </div>
      </div>
    </div>
  </div>
</div>
