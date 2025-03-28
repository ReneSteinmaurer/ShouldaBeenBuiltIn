<script>
  import Icon from './Icon.svelte';

  let { isSelected = false, tabIndex, tabName, icon, selectedTab = $bindable() } = $props();
  let selectedClass = $derived.by(() => {
    return isSelected ? 'tab-active border-gray-700 border rounded-xl' : '';
  });

</script>

<a
  on:click={() => selectedTab = tabIndex}
  role="tab"
  class={`tab ${selectedClass} hover:bg-primary/60 rounded-xl font-medium`}
>
        <span class="flex items-center gap-4">
          <Icon icon_definition="{icon}"></Icon>
          {tabName}
        </span>
</a>