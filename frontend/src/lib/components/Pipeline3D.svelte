<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import * as THREE from 'three';

  let canvasContainer: HTMLDivElement;
  let animationFrameId: number;
  let scene: THREE.Scene;
  let camera: THREE.PerspectiveCamera;
  let renderer: THREE.WebGLRenderer;
  let nodes: THREE.Mesh[] = [];

  let activeTooltip: { title: string; desc: string } | null = null;

  export let tickets: any[] = [];

  $: total = tickets.length;
  $: triaged = tickets.filter(t => t.status !== 'new').length;
  $: sandboxed = tickets.filter(t => t.status === 'sandbox_running' || t.status === 'diagnosed' || t.status === 'resolved').length;
  $: fixed = tickets.filter(t => t.status === 'diagnosed' || t.status === 'resolved').length;
  $: merged = tickets.filter(t => t.status === 'resolved').length;

  $: stageInfos = [
    { title: 'Step 1: Bug Reported', desc: 'A user or automated script submits a crash log. AI begins instant intake.', count: total.toString() },
    { title: 'Step 2: AI Semantic Triage', desc: 'High-dimensional vectors scan Neon Postgres to group duplicate issues.', count: triaged.toString() },
    { title: 'Step 3: Sandbox Reproduction', desc: 'An isolated Docker container runs tests on the exact codebase branch.', count: sandboxed.toString() },
    { title: 'Step 4: AI Fix Generated', desc: 'Gemini 1.5 Pro analyzes crash logs and drafts a 1-click git patch.', count: fixed.toString() },
    { title: 'Step 5: Approved & Merged', desc: 'Human lead approves the PR with 1-click. Raw logs purged for privacy.', count: merged.toString() }
  ];

  function selectNode(index: number) {
    activeTooltip = stageInfos[index];
  }

  onMount(() => {
    if (!canvasContainer) return;

    const width = canvasContainer.clientWidth;
    const height = 220;

    // 1. Scene setup
    scene = new THREE.Scene();

    // 2. Camera setup
    camera = new THREE.PerspectiveCamera(45, width / height, 0.1, 1000);
    camera.position.set(0, 2.5, 9);
    camera.lookAt(0, 0, 0);

    // 3. Renderer setup
    renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
    renderer.setSize(width, height);
    renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
    canvasContainer.appendChild(renderer.domElement);

    // 4. Lighting
    const ambientLight = new THREE.AmbientLight(0xffffff, 0.8);
    scene.add(ambientLight);

    const pointLight = new THREE.PointLight(0x6366f1, 2, 20);
    pointLight.position.set(0, 5, 5);
    scene.add(pointLight);

    // 5. 3D Wireframe Grid Floor
    const gridHelper = new THREE.GridHelper(18, 20, 0x6366f1, 0x1e1b4b);
    gridHelper.position.y = -1.2;
    scene.add(gridHelper);

    // 6. 5 Glowing 3D Pedestal Nodes
    const colors = [0xa855f7, 0x3b82f6, 0x14b8a6, 0xf59e0b, 0x10b981];
    const nodeXPositions = [-5.2, -2.6, 0, 2.6, 5.2];

    nodeXPositions.forEach((x, i) => {
      const geometry = new THREE.IcosahedronGeometry(0.5, 2);
      const material = new THREE.MeshStandardMaterial({
        color: colors[i],
        wireframe: false,
        roughness: 0.2,
        metalness: 0.8,
        emissive: colors[i],
        emissiveIntensity: 0.4
      });

      const sphere = new THREE.Mesh(geometry, material);
      sphere.position.set(x, 0, 0);
      scene.add(sphere);
      nodes.push(sphere);

      // Outer ring for node
      const ringGeo = new THREE.TorusGeometry(0.7, 0.02, 16, 100);
      const ringMat = new THREE.MeshBasicMaterial({ color: colors[i], transparent: true, opacity: 0.6 });
      const ring = new THREE.Mesh(ringGeo, ringMat);
      ring.rotation.x = Math.PI / 2;
      ring.position.set(x, -0.6, 0);
      scene.add(ring);
    });

    // 7. Connecting Beam Lines between nodes
    const points = nodeXPositions.map(x => new THREE.Vector3(x, 0, 0));
    const lineGeo = new THREE.BufferGeometry().setFromPoints(points);
    const lineMat = new THREE.LineBasicMaterial({ color: 0x6366f1, linewidth: 2, transparent: true, opacity: 0.8 });
    const connectionLine = new THREE.Line(lineGeo, lineMat);
    scene.add(connectionLine);

    // 8. Animation Loop
    let clock = new THREE.Clock();

    function animate() {
      animationFrameId = requestAnimationFrame(animate);
      const elapsedTime = clock.getElapsedTime();

      // Floating & Rotation animation for nodes
      nodes.forEach((node, i) => {
        node.position.y = Math.sin(elapsedTime * 2 + i) * 0.15;
        node.rotation.y = elapsedTime * 0.8;
        node.rotation.x = elapsedTime * 0.4;
      });

      renderer.render(scene, camera);
    }

    animate();

    const handleResize = () => {
      if (!canvasContainer) return;
      const w = canvasContainer.clientWidth;
      camera.aspect = w / height;
      camera.updateProjectionMatrix();
      renderer.setSize(w, height);
    };

    window.addEventListener('resize', handleResize);
  });

  onDestroy(() => {
    if (animationFrameId) cancelAnimationFrame(animationFrameId);
    if (renderer) renderer.dispose();
  });
</script>

<div class="space-y-3">
  <!-- 3D Canvas Container -->
  <div class="relative w-full h-[220px] rounded-xl overflow-hidden bg-[#070a12] border border-white/5" bind:this={canvasContainer}>
    <div class="absolute top-3 left-3 z-10 text-[10px] uppercase font-bold text-indigo-400 bg-indigo-950/60 border border-indigo-500/30 px-2.5 py-1 rounded-full backdrop-blur-md">
      ✨ 3D Interactive Pipeline Terrain (Three.js)
    </div>

    <!-- Node Click Selectors -->
    <div class="absolute bottom-2 inset-x-0 z-10 flex justify-between px-6 text-center">
      {#each stageInfos as info, i}
        <button 
          on:click={() => selectNode(i)}
          class="flex flex-col items-center group transition hover:scale-105"
        >
          <span class="text-[11px] font-bold text-slate-200 group-hover:text-indigo-300">{info.title.split(':')[1]}</span>
          <span class="text-[10px] font-mono font-bold text-indigo-400 bg-black/50 px-2 py-0.5 rounded border border-white/10 mt-0.5">{info.count}</span>
        </button>
      {/each}
    </div>
  </div>

  <!-- Interactive Non-Developer Tooltip Banner -->
  {#if activeTooltip}
    <div class="bg-gradient-to-r from-indigo-950/80 to-purple-950/80 border border-indigo-500/30 p-3.5 rounded-xl space-y-1 backdrop-blur-md transition">
      <div class="flex items-center justify-between text-xs font-bold text-indigo-300">
        <span>💡 {activeTooltip.title}</span>
        <button on:click={() => activeTooltip = null} class="text-slate-500 hover:text-white">✕</button>
      </div>
      <p class="text-xs text-slate-200 leading-snug">{activeTooltip.desc}</p>
    </div>
  {:else}
    <div class="text-center text-[11px] text-slate-500 font-medium py-1">
      💡 Click any 3D node above for a plain-English explanation of how AI resolves your bugs.
    </div>
  {/if}
</div>
