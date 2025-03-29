<script>
  import { faCheck, faInfo, faExclamationTriangle, faXmark  } from '@fortawesome/free-solid-svg-icons';
  import Icon from './Icon.svelte';

  const { duration = 3000 } = $props();
  let visible = $state(false);
  let message = $state("");
  let type = $state("info");
  let timeoutId = null
  let alertTypeClass = $derived.by(() => {
    switch (type) {
      case 'info':
        return "alert-info";
      case 'success':
        return "alert-success";
      case 'warning':
        return "alert-warning";
      case 'error':
        return "alert-error";
      default:
        return "alert-info";
    }
  });
  let alertIcon = $derived.by(() => {
    switch(type) {
      case 'info': return faInfo;
      case 'success': return faCheck;
      case 'warning': return faExclamationTriangle;
      case 'error': return faXmark;
      default: return faInfo;
    }
  });

  /**
   * Displays a toast notification and automatically hides it after a set time.
   * If a toast is already visible, it will be reset.
   *
   * @param {string} toastMessage - The message to display in the toast
   * @param {string} [toastType='info'] - The type of toast, affects styling
   *                                     (possible values: 'info', 'success', 'warning', 'error')
   */
  export function showToast(toastMessage, toastType = 'info') {
    if (timeoutId !== null) {
      clearTimeout(timeoutId)
      timeoutId = null
    }
    visible = true;
    message = toastMessage
    type = toastType
    timeoutId = setTimeout(() => {
      visible = false;
      timeoutId = null;
    }, duration);
  }
  
</script>

{#if visible}
  <div class="toast">
    <div class={`alert ${alertTypeClass}`}>
      <Icon icon_definition="{alertIcon}"/>
      <span>{message}</span>
    </div>
  </div>
{/if}