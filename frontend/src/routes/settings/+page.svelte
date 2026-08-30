<script lang="ts">
  import { API_BASE_URL } from '$lib/api/client';
  let apiBaseUrl = API_BASE_URL;
  let enableZeroRetention = true;
  let sseReconnectInterval = 1000;
</script>

<div class="space-y-6 max-w-3xl">
  <div>
    <h1 class="text-xl font-bold text-white tracking-tight">Platform & Environment Settings</h1>
    <p class="text-xs text-slate-400 mt-0.5">Configure API endpoints, zero-retention privacy policies, and live stream connection parameters.</p>
  </div>

  <div class="bg-[#0e1320] border border-white/[0.08] p-6 rounded-2xl space-y-5 shadow-xl">
    <div>
      <label for="backend-url-setting" class="block text-xs font-bold text-slate-300 uppercase tracking-wider mb-2">Backend REST API Base URL</label>
      <input 
        id="backend-url-setting"
        type="text" 
        bind:value={apiBaseUrl}
        class="w-full bg-[#111726] border border-white/10 rounded-xl px-4 py-2 text-xs text-white focus:outline-none focus:border-indigo-500 transition font-mono"
      />
    </div>

    <div class="flex items-center justify-between pt-2 border-t border-white/5">
      <div>
        <div class="text-xs font-bold text-white">Zero-Retention Privacy Policy</div>
        <div class="text-[11px] text-slate-400">Automatically destroy ephemeral containers & purge raw code logs after PR generation.</div>
      </div>
      <input type="checkbox" bind:checked={enableZeroRetention} class="w-4 h-4 accent-indigo-600 rounded cursor-pointer" />
    </div>

    <div class="pt-4 border-t border-white/5 flex justify-end">
      <button class="bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-bold px-5 py-2.5 rounded-xl transition shadow-lg">Save Settings</button>
    </div>
  </div>
</div>
