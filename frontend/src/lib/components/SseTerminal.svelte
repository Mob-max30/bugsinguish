<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { createSseListener } from '$lib/api/sse';
  import { API_BASE_URL } from '$lib/api/client';

  export let isConnected = false;

  let logs: Array<{ id: string; text: string; time: string; type?: string }> = [
    { id: '1', text: '[SYSTEM] Bugsinguish Realtime Stream Active', time: new Date().toLocaleTimeString(), type: 'info' }
  ];

  let cleanupSse: (() => void) | null = null;
  let fallbackInterval: any = null;

  onMount(() => {
    cleanupSse = createSseListener(
      `${API_BASE_URL}/api/stream`,
      (event) => {
        isConnected = true;
        const text = `[${event.phase || 'EVENT'}] ${event.message || JSON.stringify(event)}`;
        const type = event.phase === 'diagnosed' ? 'success' : event.phase === 'error' ? 'error' : 'info';
        logs = [...logs, { id: Math.random().toString(), text, time: new Date().toLocaleTimeString(), type }];
      },
      () => {
        isConnected = false;
      }
    );

    const sampleMessages = [
      { text: '[1/4] Semantic Deduplication: Checking vector similarity against Neon (pgvector)...', type: 'info' },
      { text: '[1/4] Triage Result: No duplicate ticket found (similarity score < 0.85 threshold).', type: 'success' },
      { text: '[2/4] Ephemeral Sandbox: Spawning Docker container [id: sb-92a10b]...', type: 'info' },
      { text: '[3/4] Test Reproduction: Running calculator.py test suite against branch main...', type: 'warn' },
      { text: '[3/4] Captured Exception: ZeroDivisionError in calculator.py line 14', type: 'error' },
      { text: '[4/4] AI Engine: Sending crash report & repository context to Gemini 1.5 Pro...', type: 'info' },
      { text: '[4/4] RCA Diagnosis Complete: Unified git patch generated. Draft PR created.', type: 'success' }
    ];

    let index = 0;
    fallbackInterval = setInterval(() => {
      if (!isConnected && index < sampleMessages.length) {
        const msg = sampleMessages[index];
        const time = new Date().toLocaleTimeString();
        logs = [...logs, { id: Math.random().toString(), text: msg.text, time, type: msg.type }];
        index++;
      }
    }, 2200);
  });

  onDestroy(() => {
    if (cleanupSse) cleanupSse();
    if (fallbackInterval) clearInterval(fallbackInterval);
  });
</script>

<div class="bg-[#0e1320] border border-white/[0.08] rounded-xl overflow-hidden shadow-2xl font-mono text-xs">
  <!-- Terminal Header -->
  <div class="bg-[#090d16] px-4 py-2.5 border-b border-white/[0.06] flex items-center justify-between">
    <div class="flex items-center space-x-2">
      <span class="w-2.5 h-2.5 rounded-full bg-rose-500/80 inline-block"></span>
      <span class="w-2.5 h-2.5 rounded-full bg-amber-500/80 inline-block"></span>
      <span class="w-2.5 h-2.5 rounded-full bg-emerald-500/80 inline-block"></span>
      <span class="text-[11px] text-slate-400 font-sans ml-2 font-medium">Agent Log Stream (/api/stream)</span>
    </div>
    <div class="flex items-center space-x-2 text-[10px]">
      <span class={`w-2 h-2 rounded-full ${isConnected ? 'bg-emerald-400 animate-pulse' : 'bg-indigo-400 animate-pulse'}`}></span>
      <span class="text-slate-400 font-sans font-medium uppercase tracking-wider">{isConnected ? 'LIVE BACKEND' : 'SIMULATED DEMO'}</span>
    </div>
  </div>

  <!-- Terminal Body -->
  <div class="p-3.5 h-48 overflow-y-auto space-y-1.5 leading-relaxed font-mono">
    {#each logs as log (log.id)}
      <div class="flex items-start space-x-3 text-[11px]">
        <span class="text-slate-500 shrink-0 select-none">{log.time}</span>
        <span class={
          log.type === 'success' ? 'text-emerald-400 font-medium' :
          log.type === 'warn' ? 'text-amber-300' :
          log.type === 'error' ? 'text-rose-400' : 'text-slate-300'
        }>
          {log.text}
        </span>
      </div>
    {/each}
  </div>
</div>
