<script>
  import {
    CaptureArea,
    CaptureSnippet,
    MakeWindowTransparent,
    MaximizeWindowToBounds,
    ResetWindowSizeToDefault,
    UndoMakeWindowTransparent,
  } from '../../../wailsjs/go/api/App.js';
  import Alert from '../../shared/Alert.svelte';
  import Icon from '../../shared/Icon.svelte';
  import { faImages, faInfoCircle, faScissors } from '@fortawesome/free-solid-svg-icons';
  import { WindowMinimise, WindowUnminimise } from '../../../wailsjs/runtime/runtime.js';
  import SnippSelectionCanvas from './SnippSelectionCanvas.svelte';

  let images = $state([]);
  let screenshotsLoading = $state(false);
  let showSnippingSelection = $state(false);
  let toastRef;

  async function handleScreenshotAllScreens() {
    WindowMinimise();
    screenshotsLoading = true;
    const imagesRes = await CaptureArea();
    WindowUnminimise();
    screenshotsLoading = false;
    images = imagesRes;
    toastRef.showToast('Screenshots were taken', 'success');
  }

  function handleSnippingTool() {
    showSnippingSelection = true;
    MaximizeWindowToBounds();
    MakeWindowTransparent();
  }

  function openImageInNewTab(index) {
    const base64String = images[index].replace(/^data:image\/png;base64,/, '');

    const byteCharacters = atob(base64String);
    const byteNumbers = new Array(byteCharacters.length);
    for (let i = 0; i < byteCharacters.length; i++) {
      byteNumbers[i] = byteCharacters.charCodeAt(i);
    }
    const byteArray = new Uint8Array(byteNumbers);

    const blob = new Blob([byteArray], { type: 'image/png' });
    const blobUrl = URL.createObjectURL(blob);

    window.open(blobUrl);
  }

  async function handleSelectionEnd(selectionBox) {
    try {
      await UndoMakeWindowTransparent();
      await ResetWindowSizeToDefault();
      showSnippingSelection = false;

      if (!selectionBox) {
        toastRef.showToast(`Selection box could not be retrieved!`, 'error');
        return;
      }
      const image = await CaptureSnippet(selectionBox.x, selectionBox.y, selectionBox.width, selectionBox.height);
      toastRef.showToast(`Snippet was created successfully`, 'success');
      images.push(image)
    } catch (e) {
      toastRef.showToast(`unexpected error: ${e.message || "Failed to process screenshot"}`, 'error');
    }
  }

</script>

<div>
  <div class="flex justify-center items-center">
    <button onclick={handleScreenshotAllScreens} class="btn gap-2 btn-primary btn-md">
      {#if screenshotsLoading}
        <span class="loading loading-spinner"></span>
        Taking Screenshots...
      {:else}
        <Icon icon_definition={faImages} />
        Take Screenshot
      {/if}
    </button>
  </div>
  <div class="text-center mt-2 text-sm text-base-content/70 flex items-center justify-center gap-2">
    <Icon icon_definition={faInfoCircle} />
    <span>The window will be closed during this operation</span>
  </div>

  <div class="flex justify-center items-center mt-8">
    <button onclick={handleSnippingTool} class="btn gap-2 btn-primary btn-md">
      <Icon icon_definition={faScissors} />
      Snipping Tool
    </button>
  </div>

  {#if showSnippingSelection}
    <div class="fixed inset-0 z-50 bg-black bg-opacity-50">
      <SnippSelectionCanvas onSelectionEnd={handleSelectionEnd} />
    </div>
  {/if}

  {#if images.length > 0}
    <div class="grid mt-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {#each images as image, i}
        <div class="card bg-base-100 shadow-md">
          <figure class="px-4 pt-4">
            <img src={image} alt={`Screenshot ${i+1}`} class="rounded-lg object-cover w-full" />
          </figure>
          <div class="card-body items-center text-center pt-4 pb-4">
            <h3 class="card-title text-sm">Bildschirm {i + 1}</h3>
            <div class="card-actions">
              <button
                onclick={() => openImageInNewTab(i)}
                class="btn btn-soft btn-md btn-primary"
              >
                Open in new window
              </button>
            </div>
          </div>
        </div>
      {/each}
    </div>
  {:else if !screenshotsLoading}
    <div class="text-center py-12 text-base-content/70">
      <p>No screenshots available. Click on "Take Screenshot" to get started.</p>
    </div>
  {/if}
  <Alert bind:this={toastRef}></Alert>
</div>