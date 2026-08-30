<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { createSseListener } from '$lib/api/sse';

  export let isConnected = false;

  let logs: Array<{ id: string; text: string; time: string; type?: string }> = [
    { id: '1', text: '[SYSTEM] Initializing Bugsinguish Realtime Stream...', time: new Date().toLocaleTimeString(), type: 'info' }
  ];

  let cleanupSse: (() => void) | null = null;
  let fallbackInterval: any = null;

  onMount(() => {
    // Attempt real SSE connection to backend
    cleanupSse = createSseListener(
      'http://localhost:8080/api/stream',
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

    // Fallback simulated stream for standalone UI testing
    const sampleMessages = [
      { text: '[1/4] Checking semantic vector similarity against Neon (pgvector)...', type: 'info' },
      { text: '[1/4] No duplicate ticket found (similarity score < 0.85 threshold).', type: 'success' },
      { text: '[2/4] Triggering Docker Sandbox manager (image: Dockerfile.dummy)...', type: 'info' },
      { text: '[2/4] Spawning isolated container [id: sb-92a10b]...', type: 'info' },
      { text: '[3/4] Running test suite against branch main...', type: 'warn' },
      { text: '[3/4] Captured crash output: ZeroDivisionError in calc.py line 14', type: 'error' },
      { text: '[4/4] Sending crash report & snippet to Gemini 1.5 Pro...', type: 'info' },
      { text: '[4/4] Root cause analysis completed. Git unified diff generated.', type: 'success' }
    ];

    let index = 0;
    fallbackInterval = setInterval(() => {
      if (!isConnected && index < sampleMessages.length) {
        const msg = sampleMessages[index];
        const time = new Date().toLocaleTimeString();
        logs = [...logs, { id: Math.random().toString(), text: msg.text, time, type: msg.type }];
        index++;
      }
    }, 2000);
  });

  onDestroy(() => {
    if (cleanupSse) cleanupSse();
    if (fallbackInterval) clearInterval(fallbackInterval);
  });
</script>

<div class="bg-slate-900 border border-slate-800 rounded-lg overflow-hidden shadow-xl font-mono text-sm">
  <div class="bg-slate-950 px-4 py-2 border-b border-slate-800 flex items-center justify-between">
    <div class="flex items-center space-x-2">
      <span class="w-3 h-3 rounded-full bg-rose-500 inline-block"></span>
      <span class="w-3 h-3 rounded-full bg-amber-500 inline-block"></span>
      <span class="w-3 h-3 rounded-full bg-emerald-500 inline-block"></span>
      <span class="text-xs text-slate-400 font-sans ml-2">Live Agent Stream (http://localhost:8080/api/stream)</span>
    </div>
    <div class="flex items-center space-x-2 text-xs">
      <span class={`w-2 h-2 rounded-full ${isConnected ? 'bg-emerald-400 animate-pulse' : 'bg-amber-500 animate-pulse'}`}></span>
      <span class="text-slate-400 font-sans">{isConnected ? 'CONNECTED' : 'SIMULATED DEMO'}</span>
    </div>
  </div>

  <div class="p-4 h-56 overflow-y-auto space-y-1.5 scrollbar-thin scrollbar-thumb-slate-700">
    {#each logs as log (log.id)}
      <div class="flex items-start space-x-3 text-xs leading-relaxed">
        <span class="text-slate-500 shrink-0 select-none">{log.time}</span>
        <span class={
          log.type === 'success' ? 'text-emerald-400' :
          log.type === 'warn' ? 'text-amber-300' :
          log.type === 'error' ? 'text-rose-400' : 'text-slate-300'
        }>
          {log.text}
        </span>
      </div>
    {/each}
  </div>
</div>
