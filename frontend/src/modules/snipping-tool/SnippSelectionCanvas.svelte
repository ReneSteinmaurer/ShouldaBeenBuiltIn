<script>
  import { onMount } from 'svelte';
  import { GetWindowSize } from '../../../wailsjs/go/api/App.js';

  const {
    onSelectionEnd = () => {},
  } = $props();

  let canvas = $state();
  let ctx = $state();
  let isSelecting = $state(false);
  let startX = $state(0);
  let startY = $state(0);
  let selectionBox = { x: 0, y: 0, width: 0, height: 0 };

  onMount(async() => {
    canvas = document.getElementById('selection-canvas');
    ctx = canvas.getContext('2d');
    const windowSize = await GetWindowSize()
    resizeCanvas(windowSize.width, windowSize.height);
  });

  function resizeCanvas(width, height) {
    canvas.width = width;
    canvas.height = height;
    console.log(width, height);
    if (isSelecting) {
      drawSelection();
    }
  }

  function clearCanvas() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.fillStyle = 'rgba(0,0,0,0.3)';
    ctx.fillRect(0, 0, canvas.width, canvas.height);
  }

  function startSelection(e) {
    isSelecting = true;
    const rect = canvas.getBoundingClientRect();
    startX = (e.clientX - rect.left) * (canvas.width / rect.width);
    startY = (e.clientY - rect.top) * (canvas.height / rect.height);
    clearCanvas();
  }

  function updateSelection(e) {
      if (!isSelecting) return;

      const rect = canvas.getBoundingClientRect();

      const currentX = (e.clientX - rect.left) * (canvas.width / rect.width);
      const currentY = (e.clientY - rect.top) * (canvas.height / rect.height);

      const width = Math.abs(currentX - startX);
      const height = Math.abs(currentY - startY);
      const x = Math.min(startX, currentX);
      const y = Math.min(startY, currentY);

      selectionBox = { x, y, width, height };
      drawSelection();
  }

  function endSelection() {
    if (!isSelecting) return;
    isSelecting = false;
    onSelectionEnd(selectionBox)
  }

  function drawSelection() {
    clearCanvas();

    ctx.clearRect(selectionBox.x, selectionBox.y, selectionBox.width, selectionBox.height);

    ctx.strokeStyle = 'white';
    ctx.setLineDash([5, 5]);
    ctx.strokeRect(selectionBox.x, selectionBox.y, selectionBox.width, selectionBox.height);

    ctx.fillStyle = 'white';
    ctx.font = '14px Arial';
    ctx.fillText(`${selectionBox.width} x ${selectionBox.height}`,
      selectionBox.x + selectionBox.width + 5,
      selectionBox.y + selectionBox.height + 20);
  }
</script>

<div class="z-1000 w-full h-full">
  <canvas id="selection-canvas"
          onmousedown={startSelection}
          onmousemove={updateSelection}
          onmouseup={endSelection}
          style="width: 100%; height: 100%; cursor: crosshair;">
  </canvas>
</div>