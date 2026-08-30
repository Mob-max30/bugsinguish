<script lang="ts">
  let prs = [
    {
      id: 'PR #512',
      title: 'Fix null pointer in login redirect',
      issue: '#BUG-1248',
      branch: 'fix/login-redirect',
      target: 'main',
      file: 'AuthService.java',
      diff: '--- AuthService.java\n+++ AuthService.java\n@@ -41,3 +41,5 @@\n+if (token == null) {\n+    throw new AuthException("Token missing");\n+}',
      status: 'Awaiting Merge'
    },
    {
      id: 'PR #511',
      title: 'Add zero denominator validation in divide()',
      issue: '#BUG-101',
      branch: 'fix/zero-division',
      target: 'main',
      file: 'sandbox/dummy_repo/calculator.py',
      diff: '--- calculator.py\n+++ calculator.py\n@@ -13,3 +13,5 @@\n def divide(a, b):\n+    if b == 0:\n+        raise ValueError("Cannot divide by zero")\n     return a / b',
      status: 'Ready'
    }
  ];
</script>

<div class="space-y-6">
  <div>
    <h1 class="text-xl font-bold text-white tracking-tight">AI Generated Pull Requests</h1>
    <p class="text-xs text-slate-400 mt-0.5">Automated code patches drafted by Gemini 1.5 Pro after sandbox validation.</p>
  </div>

  <div class="space-y-4">
    {#each prs as pr}
      <div class="bg-[#0e1320] border border-white/[0.08] p-6 rounded-2xl space-y-4 shadow-xl">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <span class="text-lg">🔀</span>
            <div>
              <h3 class="text-sm font-bold text-white">{pr.title}</h3>
              <div class="text-[10px] text-slate-400 font-mono mt-0.5">{pr.id} • Issue {pr.issue} • Branch <span class="text-indigo-300">{pr.branch}</span> → {pr.target}</div>
            </div>
          </div>

          <button class="bg-emerald-600 hover:bg-emerald-500 text-white text-xs font-bold px-4 py-2 rounded-xl transition shadow-lg">
            ✓ Approve & Merge PR →
          </button>
        </div>

        <div class="bg-[#090d16] border border-white/[0.06] rounded-xl overflow-hidden font-mono text-xs">
          <div class="bg-[#111726] px-4 py-2 text-[11px] text-slate-400 border-b border-white/5">{pr.file}</div>
          <pre class="p-4 text-emerald-400 overflow-x-auto"><code>{pr.diff}</code></pre>
        </div>
      </div>
    {/each}
  </div>
</div>
