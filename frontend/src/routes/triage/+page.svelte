<script lang="ts">
  let testQuery = 'App crashes when signing in on mobile device';
  let isScanning = false;
  let matches = [
    { id: 'BUG-1248', title: 'Users unable to login on iOS app', score: 0.94, status: 'Duplicate Identified', repo: 'iOS App' },
    { id: 'BUG-101', title: 'Authentication token expiration silently breaks SSE', score: 0.78, status: 'Related Topic', repo: 'Web App' },
    { id: 'BUG-103', title: 'NullPointerException on User Profile fetch', score: 0.62, status: 'Low Similarity', repo: 'Backend API' }
  ];

  function runSimulatedVectorSearch() {
    isScanning = true;
    setTimeout(() => {
      isScanning = false;
    }, 600);
  }
</script>

<div class="space-y-6">
  <div>
    <h1 class="text-xl font-bold text-white tracking-tight">AI Semantic Triage & Vector Deduplication</h1>
    <p class="text-xs text-slate-400 mt-0.5">High-dimensional text embeddings in Neon Postgres (`pgvector`) detect duplicate issues regardless of phrasing.</p>
  </div>

  <!-- Interactive Vector Search Harness -->
  <div class="bg-[#0e1320] border border-white/[0.08] p-6 rounded-2xl space-y-4 shadow-xl">
    <h3 class="text-xs font-bold uppercase tracking-wider text-indigo-400">Test Vector Deduplication Matcher</h3>
    
    <div class="flex items-center space-x-3">
      <input 
        type="text" 
        bind:value={testQuery}
        placeholder="Type any bug report phrasing to test vector similarity..."
        class="flex-1 bg-[#111726] border border-white/10 rounded-xl px-4 py-2.5 text-xs text-white placeholder-slate-500 focus:outline-none focus:border-indigo-500 transition"
      />
      <button 
        on:click={runSimulatedVectorSearch}
        disabled={isScanning}
        class="bg-indigo-600 hover:bg-indigo-500 text-white text-xs font-bold px-5 py-2.5 rounded-xl shadow-lg transition"
      >
        {isScanning ? 'Scanning Embeddings...' : '⚡ Search Vector DB'}
      </button>
    </div>

    <!-- Match Results -->
    <div class="space-y-3 pt-2">
      <h4 class="text-[11px] font-bold text-slate-400 uppercase tracking-wider">Similarity Score Results (pgvector Cosine Distance)</h4>
      
      <div class="space-y-2">
        {#each matches as m}
          <div class="flex items-center justify-between p-3.5 rounded-xl bg-[#111726] border border-white/5">
            <div class="flex items-center space-x-3">
              <span class="w-8 h-8 rounded-lg bg-indigo-500/20 text-indigo-300 flex items-center justify-center font-mono font-bold text-xs">
                {Math.round(m.score * 100)}%
              </span>
              <div>
                <div class="text-xs font-bold text-white">{m.title}</div>
                <div class="text-[10px] text-slate-400 font-mono mt-0.5">#{m.id} • {m.repo}</div>
              </div>
            </div>

            <span class={`text-[10px] font-bold px-2.5 py-1 rounded-full border ${
              m.score > 0.85 ? 'bg-purple-500/10 text-purple-400 border-purple-500/20' : 'bg-slate-500/10 text-slate-400 border-slate-500/20'
            }`}>
              {m.status}
            </span>
          </div>
        {/each}
      </div>
    </div>
  </div>
</div>
