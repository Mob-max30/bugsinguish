<script lang="ts">
  import type { Ticket } from '$lib/types';

  export let ticket: Ticket | null = null;
  export let isOpen = false;
  export let onClose: () => void = () => {};

  let activeTab: 'description' | 'logs' | 'ai_diagnosis' | 'diff' = 'ai_diagnosis';

  function handleOverlayClick(e: MouseEvent) {
    if (e.target === e.currentTarget) {
      onClose();
    }
  }
</script>

{#if isOpen && ticket}
  <!-- Modal Overlay -->
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <div 
    class="fixed inset-0 z-50 bg-slate-950/80 backdrop-blur-sm flex items-center justify-center p-4"
    on:click={handleOverlayClick}
  >
    <div class="bg-slate-900 border border-slate-800 rounded-xl w-full max-w-4xl max-h-[90vh] flex flex-col shadow-2xl overflow-hidden">
      <!-- Header -->
      <div class="px-6 py-4 border-b border-slate-800 flex items-center justify-between bg-slate-950/50">
        <div>
          <div class="flex items-center space-x-3">
            <span class="text-xs font-mono px-2 py-0.5 rounded bg-slate-800 text-slate-300">
              {ticket.id}
            </span>
            <span class={`text-xs font-semibold px-2 py-0.5 rounded capitalize ${
              ticket.severity === 'critical' ? 'bg-rose-950/80 text-rose-300 border border-rose-800' :
              ticket.severity === 'high' ? 'bg-amber-950/80 text-amber-300 border border-amber-800' :
              ticket.severity === 'medium' ? 'bg-blue-950/80 text-blue-300 border border-blue-800' :
              'bg-slate-800 text-slate-300'
            }`}>
              {ticket.severity} severity
            </span>
            <span class="text-xs font-mono text-slate-400">
              Branch: <code class="text-indigo-400">{ticket.repo_branch_url || 'main'}</code>
            </span>
          </div>
          <h2 class="text-xl font-bold text-slate-100 mt-2">{ticket.title}</h2>
        </div>
        <button 
          on:click={onClose}
          class="text-slate-400 hover:text-slate-100 p-2 rounded-lg hover:bg-slate-800 transition"
          aria-label="Close modal"
        >
          ✕
        </button>
      </div>

      <!-- Tabs Nav -->
      <div class="flex border-b border-slate-800 bg-slate-950/30 px-6">
        <button 
          class={`px-4 py-3 text-sm font-medium border-b-2 transition ${
            activeTab === 'ai_diagnosis' 
              ? 'border-indigo-500 text-indigo-400' 
              : 'border-transparent text-slate-400 hover:text-slate-200'
          }`}
          on:click={() => activeTab = 'ai_diagnosis'}
        >
          ⚡ AI Root Cause Diagnosis
        </button>

        <button 
          class={`px-4 py-3 text-sm font-medium border-b-2 transition ${
            activeTab === 'diff' 
              ? 'border-indigo-500 text-indigo-400' 
              : 'border-transparent text-slate-400 hover:text-slate-200'
          }`}
          on:click={() => activeTab = 'diff'}
        >
          📝 Generated Code Diff (Draft PR)
        </button>

        <button 
          class={`px-4 py-3 text-sm font-medium border-b-2 transition ${
            activeTab === 'description' 
              ? 'border-indigo-500 text-indigo-400' 
              : 'border-transparent text-slate-400 hover:text-slate-200'
          }`}
          on:click={() => activeTab = 'description'}
        >
          📄 Description & Stack Trace
        </button>

        <button 
          class={`px-4 py-3 text-sm font-medium border-b-2 transition ${
            activeTab === 'logs' 
              ? 'border-indigo-500 text-indigo-400' 
              : 'border-transparent text-slate-400 hover:text-slate-200'
          }`}
          on:click={() => activeTab = 'logs'}
        >
          🖥️ Sandbox Execution Logs
        </button>
      </div>

      <!-- Tab Content -->
      <div class="p-6 overflow-y-auto flex-1 space-y-4">
        {#if activeTab === 'ai_diagnosis'}
          {#if ticket.diagnosis}
            <div class="space-y-4">
              <div class="bg-indigo-950/30 border border-indigo-800/50 rounded-lg p-4">
                <h4 class="text-xs uppercase tracking-wider font-semibold text-indigo-400 mb-1">Root Cause Summary</h4>
                <p class="text-slate-200 text-sm font-medium">{ticket.diagnosis.root_cause}</p>
              </div>

              <div class="bg-slate-950/80 border border-slate-800 rounded-lg p-4">
                <h4 class="text-xs uppercase tracking-wider font-semibold text-slate-400 mb-2">Detailed AI Analysis (Gemini 1.5 Pro)</h4>
                <p class="text-slate-300 text-sm leading-relaxed whitespace-pre-line">{ticket.diagnosis.explanation}</p>
              </div>

              <div class="flex items-center justify-between bg-slate-950 border border-slate-800 rounded-lg px-4 py-3">
                <div class="flex items-center space-x-2">
                  <span class="text-xs text-slate-400">Target File:</span>
                  <code class="text-xs text-emerald-400 font-mono font-semibold">{ticket.diagnosis.file}</code>
                </div>
                <button class="bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-semibold px-4 py-2 rounded-lg transition shadow-lg">
                  Approve & Merge Draft PR →
                </button>
              </div>
            </div>
          {:else}
            <div class="text-center py-12 text-slate-500">
              <p class="text-sm">AI Diagnosis is in progress or pending sandbox completion...</p>
            </div>
          {/if}

        {:else if activeTab === 'diff'}
          {#if ticket.diagnosis?.diff}
            <div class="bg-slate-950 border border-slate-800 rounded-lg overflow-hidden">
              <div class="bg-slate-900 px-4 py-2 text-xs font-mono text-slate-400 border-b border-slate-800 flex justify-between">
                <span>{ticket.diagnosis.file}</span>
                <span>Unified Diff Patch</span>
              </div>
              <pre class="p-4 text-xs font-mono overflow-x-auto text-slate-200">
                <code>{ticket.diagnosis.diff}</code>
              </pre>
            </div>
          {:else}
            <div class="text-center py-12 text-slate-500">
              <p class="text-sm">No code patch generated yet.</p>
            </div>
          {/if}

        {:else if activeTab === 'description'}
          <div class="space-y-4">
            <div>
              <h4 class="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-2">User Bug Description</h4>
              <p class="text-slate-300 text-sm bg-slate-950 p-4 rounded-lg border border-slate-800 leading-relaxed">
                {ticket.description}
              </p>
            </div>

            {#if ticket.stack_trace}
              <div>
                <h4 class="text-xs font-semibold text-slate-400 uppercase tracking-wider mb-2">Submitted Stack Trace / Crash Log</h4>
                <pre class="bg-slate-950 p-4 rounded-lg border border-slate-800 text-xs font-mono text-rose-300 overflow-x-auto">
                  <code>{ticket.stack_trace}</code>
                </pre>
              </div>
            {/if}
          </div>

        {:else if activeTab === 'logs'}
          <div class="bg-slate-950 p-4 rounded-lg border border-slate-800 font-mono text-xs text-slate-300 space-y-1">
            {#if ticket.logs && ticket.logs.length > 0}
              {#each ticket.logs as log}
                <div class="py-0.5">{log}</div>
              {/each}
            {:else}
              <div class="text-slate-500">[DOCKER SANDBOX] Container sb-92a10b mounted /workspace/dummy_repo</div>
              <div class="text-slate-500">[DOCKER SANDBOX] Running `pytest test_calculator.py`...</div>
              <div class="text-rose-400">[FAILURE] ZeroDivisionError: division by zero in divide(a, b) at calc.py:14</div>
              <div class="text-slate-500">[DOCKER SANDBOX] Container destroyed. Zero-retention policy applied.</div>
            {/if}
          </div>
        {/if}
      </div>

      <!-- Footer -->
      <div class="px-6 py-4 border-t border-slate-800 bg-slate-950/50 flex justify-between items-center text-xs text-slate-400">
        <span>Created: {ticket.created_at}</span>
        <button 
          on:click={onClose}
          class="px-4 py-2 rounded-lg bg-slate-800 hover:bg-slate-700 text-slate-200 font-medium transition"
        >
          Close
        </button>
      </div>
    </div>
  </div>
{/if}
