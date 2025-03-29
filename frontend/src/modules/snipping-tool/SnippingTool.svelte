<script>
  import {CaptureArea} from '../../../wailsjs/go/api/App.js';
  import Alert from '../../shared/Alert.svelte';
  import Icon from '../../shared/Icon.svelte';
  import { faImages } from '@fortawesome/free-solid-svg-icons';
  import {
    WindowHide,
    WindowMaximise,
    WindowMinimise,
    WindowShow,
    WindowUnminimise,
  } from '../../../wailsjs/runtime/runtime.js';

  let images = $state([]);
  let loading = $state(false);
  let toastRef

  async function handleScreenshot() {
    WindowMinimise()
    loading = true;
    const imagesRes = await CaptureArea();
    WindowUnminimise()
    loading = false
    images = imagesRes;
    toastRef.showToast('Screenshots were taken', 'success');
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
</script>

<div>
  <div class="flex justify-center items-center">
    <button on:click={handleScreenshot} class="btn gap-2 btn-primary btn-md">
      {#if loading}
        <span class="loading loading-spinner"></span>
        Taking Screenshots...
        {:else}
        <Icon icon_definition={faImages} />
        Screenshot
      {/if}
    </button>
  </div>
  {#if images.length > 0}
    <div class="grid mt-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {#each images as image, i}
        <div class="card bg-base-100 shadow-md">
          <figure class="px-4 pt-4">
            <img src={image} alt={`Screenshot ${i+1}`} class="rounded-lg object-cover w-full" />
          </figure>
          <div class="card-body items-center text-center pt-4 pb-4">
            <h3 class="card-title text-sm">Bildschirm {i+1}</h3>
            <div class="card-actions">
              <button
                on:click={() => openImageInNewTab(i)}
                class="btn btn-soft btn-md btn-primary"
              >
                Open in new window
              </button>
            </div>
          </div>
        </div>
      {/each}
    </div>
  {:else if !loading}
    <div class="text-center py-12 text-base-content/70">
      <p>No screenshots available. Click on "Take Screenshot" to get started.</p>
    </div>
  {/if}
  <Alert bind:this={toastRef}></Alert>
</div>