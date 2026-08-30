<script lang="ts">
  import type { TicketSeverity } from '$lib/types';

  let title = '';
  let description = '';
  let stackTrace = '';
  let repoBranchUrl = 'https://github.com/Mob-max30/bugsinguish';
  let severity: TicketSeverity = 'medium';

  let errors: { [key: string]: string } = {};
  let submittedPayload: any = null;
  let isSubmitting = false;

  function validate() {
    errors = {};
    if (!title.trim()) errors.title = 'Title is required.';
    if (!description.trim()) errors.description = 'Description is required.';
    if (!repoBranchUrl.trim()) errors.repoBranchUrl = 'Repository branch URL is required.';
    return Object.keys(errors).length === 0;
  }

  function handleSubmit() {
    if (!validate()) return;

    isSubmitting = true;
    const payload = {
      title,
      description,
      stack_trace: stackTrace,
      repo_branch_url: repoBranchUrl,
      severity,
      submitted_at: new Date().toISOString()
    };

    console.log('[BUGSINGUISH INTAKE] Submitting bug report:', payload);

    setTimeout(() => {
      isSubmitting = false;
      submittedPayload = payload;
    }, 600);
  }

  function resetForm() {
    title = '';
    description = '';
    stackTrace = '';
    severity = 'medium';
    submittedPayload = null;
  }
</script>

<div class="max-w-3xl mx-auto space-y-6 py-4">
  <div>
    <a href="/" class="text-xs font-semibold text-indigo-400 hover:text-indigo-300 transition">← Back to Kanban Board</a>
    <h1 class="text-2xl font-black text-slate-100 mt-2">Submit New Bug Report</h1>
    <p class="text-sm text-slate-400 mt-1">Provide crash logs and code context for autonomous sandbox reproduction & AI root-cause analysis.</p>
  </div>

  {#if submittedPayload}
    <div class="bg-emerald-950/60 border border-emerald-800 rounded-xl p-6 space-y-4">
      <div class="flex items-center space-x-3 text-emerald-400">
        <span class="text-xl">✓</span>
        <h3 class="text-lg font-bold">Bug Report Submitted Successfully</h3>
      </div>
      <p class="text-xs text-slate-300">
        Ticket has been queued for semantic deduplication check against Neon (pgvector). Check the live SSE stream on the dashboard.
      </p>

      <pre class="bg-slate-950 p-4 rounded-lg text-xs font-mono text-slate-300 overflow-x-auto border border-slate-900">
        <code>{JSON.stringify(submittedPayload, null, 2)}</code>
      </pre>

      <div class="flex space-x-3">
        <a href="/" class="bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-semibold px-4 py-2 rounded-lg transition">
          View on Kanban Board →
        </a>
        <button on:click={resetForm} class="bg-slate-800 hover:bg-slate-700 text-slate-300 text-xs font-semibold px-4 py-2 rounded-lg transition">
          Submit Another Report
        </button>
      </div>
    </div>
  {:else}
    <form on:submit|preventDefault={handleSubmit} class="bg-slate-900 border border-slate-800 rounded-xl p-6 space-y-5 shadow-xl">
      <!-- Title -->
      <div>
        <label for="title" class="block text-xs font-semibold uppercase tracking-wider text-slate-300 mb-1.5">
          Bug Title <span class="text-rose-400">*</span>
        </label>
        <input 
          id="title"
          type="text" 
          bind:value={title}
          placeholder="e.g. ZeroDivisionError when denominator is zero in divide()"
          class="w-full bg-slate-950 border border-slate-800 rounded-lg px-4 py-2.5 text-sm text-slate-100 focus:outline-none focus:border-indigo-500 transition"
        />
        {#if errors.title}
          <p class="text-xs text-rose-400 mt-1">{errors.title}</p>
        {/if}
      </div>

      <!-- Severity & Branch URL -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label for="severity" class="block text-xs font-semibold uppercase tracking-wider text-slate-300 mb-1.5">
            Severity Level
          </label>
          <select 
            id="severity"
            bind:value={severity}
            class="w-full bg-slate-950 border border-slate-800 rounded-lg px-4 py-2.5 text-sm text-slate-100 focus:outline-none focus:border-indigo-500 transition"
          >
            <option value="low">Low</option>
            <option value="medium">Medium</option>
            <option value="high">High</option>
            <option value="critical">Critical</option>
          </select>
        </div>

        <div>
          <label for="branch" class="block text-xs font-semibold uppercase tracking-wider text-slate-300 mb-1.5">
            Target Repo / Branch URL <span class="text-rose-400">*</span>
          </label>
          <input 
            id="branch"
            type="text" 
            bind:value={repoBranchUrl}
            placeholder="https://github.com/org/repo/tree/main"
            class="w-full bg-slate-950 border border-slate-800 rounded-lg px-4 py-2.5 text-sm text-slate-100 focus:outline-none focus:border-indigo-500 transition"
          />
          {#if errors.repoBranchUrl}
            <p class="text-xs text-rose-400 mt-1">{errors.repoBranchUrl}</p>
          {/if}
        </div>
      </div>

      <!-- Description -->
      <div>
        <label for="desc" class="block text-xs font-semibold uppercase tracking-wider text-slate-300 mb-1.5">
          Detailed Description <span class="text-rose-400">*</span>
        </label>
        <textarea 
          id="desc"
          bind:value={description}
          rows="3"
          placeholder="Describe what happened, expected vs actual behavior..."
          class="w-full bg-slate-950 border border-slate-800 rounded-lg px-4 py-2.5 text-sm text-slate-100 focus:outline-none focus:border-indigo-500 transition"
        ></textarea>
        {#if errors.description}
          <p class="text-xs text-rose-400 mt-1">{errors.description}</p>
        {/if}
      </div>

      <!-- Stack Trace / Error Log -->
      <div>
        <label for="stack" class="block text-xs font-semibold uppercase tracking-wider text-slate-300 mb-1.5">
          Stack Trace / Error Log (Optional)
        </label>
        <textarea 
          id="stack"
          bind:value={stackTrace}
          rows="5"
          placeholder="Paste console traceback or stderr logs here..."
          class="w-full bg-slate-950 border border-slate-800 rounded-lg px-4 py-2.5 text-xs font-mono text-rose-300 focus:outline-none focus:border-indigo-500 transition"
        ></textarea>
      </div>

      <!-- Submit Button -->
      <div class="pt-2 flex justify-end">
        <button 
          type="submit" 
          disabled={isSubmitting}
          class="bg-indigo-600 hover:bg-indigo-500 disabled:opacity-50 text-white font-semibold text-sm px-6 py-2.5 rounded-lg shadow-lg transition"
        >
          {isSubmitting ? 'Submitting...' : 'Submit Bug for AI Diagnosis →'}
        </button>
      </div>
    </form>
  {/if}
</div>
