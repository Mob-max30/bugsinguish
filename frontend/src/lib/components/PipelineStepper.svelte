<script lang="ts">
  import type { TicketStatus } from '$lib/types';
  
  export let currentStage: TicketStatus = 'new';
  export let onSelectStage: (stage: TicketStatus) => void = () => {};

  // We are using the exact stage array to drive the pipeline, as requested.
  const stages: { id: TicketStatus; label: string; countStub: string; colorClass: string; bgClass: string; icon: string }[] = [
    { 
      id: 'new', label: 'Reported', countStub: '1,248', colorClass: 'text-[var(--color-status-blue)]', bgClass: 'bg-[var(--color-status-blue)]',
      icon: '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />'
    },
    { 
      id: 'triaging', label: 'AI Triage', countStub: '512', colorClass: 'text-[var(--color-status-blue)]', bgClass: 'bg-[var(--color-status-blue)]',
      icon: '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />'
    },
    { 
      id: 'sandbox_running', label: 'Sandbox', countStub: '320', colorClass: 'text-[var(--color-status-cyan)]', bgClass: 'bg-[var(--color-status-cyan)]',
      icon: '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />'
    },
    { 
      id: 'diagnosed', label: 'Fix Generated', countStub: '187', colorClass: 'text-[var(--color-brand-purple)]', bgClass: 'bg-[var(--color-brand-purple)]',
      icon: '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />'
    },
    { 
      id: 'resolved', label: 'Merged', countStub: '98', colorClass: 'text-[var(--color-status-green)]', bgClass: 'bg-[var(--color-status-green)]',
      icon: '<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />'
    }
  ];

  $: currentIndex = stages.findIndex(s => s.id === currentStage);
</script>

<div class="bg-[var(--color-surface)] border border-[var(--color-border)] rounded-xl p-6 h-full flex flex-col">
  <div class="flex items-center justify-between mb-8">
    <h2 class="text-base font-semibold text-[var(--color-text-primary)]">Resolution Pipeline</h2>
    <button class="text-sm text-[var(--color-brand-purple)] font-medium hover:underline">View All</button>
  </div>

  <div class="flex items-center w-full justify-between flex-1 relative px-4">
    {#each stages as stage, i}
      {@const isCompleted = i <= currentIndex}
      {@const isActive = i === currentIndex}
      
      <div class="relative flex flex-col items-center group cursor-pointer z-10" on:click={() => onSelectStage(stage.id)}>
        <!-- Glowing Icon Node -->
        <div class={`w-14 h-14 rounded-full flex items-center justify-center transition-all duration-300 relative ${isCompleted ? stage.bgClass + '/10' : 'bg-[var(--color-border)]/30'}`}>
          
          {#if isCompleted}
            <!-- Glow effect behind -->
            <div class={`absolute inset-0 rounded-full blur-md opacity-40 ${stage.bgClass}`}></div>
          {/if}

          <!-- Inner solid circle with icon -->
          <div class={`relative z-10 w-10 h-10 rounded-full flex items-center justify-center ${isCompleted ? stage.bgClass : 'bg-[var(--color-border)]'}`}>
            <svg class="w-5 h-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              {@html stage.icon}
            </svg>
          </div>
        </div>

        <!-- Labels and Stub Counts beneath -->
        <div class="mt-4 text-center">
          <div class={`text-sm font-medium ${isCompleted ? 'text-[var(--color-text-primary)]' : 'text-[var(--color-text-muted)]'}`}>
            {stage.label}
          </div>
          <!-- STUBBED: Since PipelineStepper doesn't receive `tickets` prop, we use static stubs to match the design counts -->
          <div class={`text-lg font-bold mt-1 ${isCompleted ? 'text-[var(--color-text-primary)]' : 'text-[var(--color-text-muted)]'}`}>
            {stage.countStub}
          </div>
        </div>
      </div>

      <!-- Connector Arrows -->
      {#if i < stages.length - 1}
        <div class="flex-1 flex justify-center -mt-12 z-0">
          <svg class="w-6 h-6 text-[var(--color-brand-purple)] opacity-50" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 8l4 4m0 0l-4 4m4-4H3" />
          </svg>
        </div>
      {/if}
    {/each}
  </div>
  
  <div class="mt-6 pt-4 border-t border-[var(--color-border)] flex items-center justify-between text-sm">
    <div class="flex items-center space-x-3 w-1/2">
      <span class="text-[var(--color-text-secondary)]">Auto-resolution rate: 38.7%</span>
      <div class="flex-1 h-2 bg-[var(--color-border)] rounded-full overflow-hidden">
        <div class="h-full bg-[var(--color-status-green)] rounded-full w-[38.7%] shadow-[0_0_8px_rgba(16,185,129,0.6)]"></div>
      </div>
    </div>
    <span class="text-[var(--color-text-secondary)]">Improving steadily! 🚀</span>
  </div>
</div>
