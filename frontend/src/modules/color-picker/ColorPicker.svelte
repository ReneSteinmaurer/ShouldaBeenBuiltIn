<script>
  import { faEyeDropper } from '@fortawesome/free-solid-svg-icons';
  import Icon from '../../shared/Icon.svelte';
  import Alert from '../../shared/Alert.svelte';
  import { ClipboardSetText } from '../../../wailsjs/runtime/runtime.js';

  let colorCode = $state('#ffffff');
  const colorCodeHistory = $state([]);
  let colorInputRef;
  let toastRef;

  function toggleColorPicker() {
    colorInputRef.click();
  }

  function handleColorChange(event) {
    colorCode = event.target.value;
    colorCodeHistory.push(colorCode);
  }

  function copyColorCodeToClipboardAndDisplayToast(colorCode) {
    const ok = ClipboardSetText(colorCode)
    if (!ok) {
      toastRef.showToast(`An unexpected error occurred while saving to clipboard`, 'error');
      return
    }
    toastRef.showToast(`Color code '${colorCode}' was copied to clipboard!`, 'success');
  }

</script>

<div class="flex flex-col items-center gap-4 p-4 rounded-lg bg-base-200">
  <button class="btn btn-primary btn-md gap-2" on:click={toggleColorPicker}>
    <Icon icon_definition="{faEyeDropper}" />
    <span>Pick Color</span>
  </button>

  <div class="flex items-center gap-3">
    <div class="w-10 h-10 rounded-full border-4 border-base-content/10 shadow-inner"
         style="background-color: {colorCode}"></div>
    <div on:click={() => copyColorCodeToClipboardAndDisplayToast(colorCode)}
         class="badge badge-lg hover:cursor-pointer font-mono bg-base-300 text-base-content">{colorCode}</div>
  </div>

  {#if colorCodeHistory.length > 0}
    <div class="flex-1 mt-8">
      <h3 class="flex justify-center items-center text-lg font-medium text-base-content/80 mb-3">Color History</h3>
      <div class="grid grid-cols-5 gap-4">
        {#each colorCodeHistory as code}
          <div on:click={() => copyColorCodeToClipboardAndDisplayToast(code)} class="flex flex-col items-center group">
            <div
              class="w-12 h-12 rounded-xl border-2 border-base-content/10 shadow-md hover:shadow-lg transition-all cursor-pointer group-hover:scale-105"
              style="background-color: {code}"></div>
            <div class="mt-1 px-2 py-0.5 font-mono text-xs opacity-80 group-hover:opacity-100">{code}</div>
          </div>
        {/each}
      </div>
    </div>
  {/if}


  <input class="hidden" bind:this={colorInputRef} type="color" value="{colorCode}" on:change={handleColorChange}>
  <Alert bind:this={toastRef}></Alert>
</div>
