<script>
  import ColorPicker from './modules/color-picker/ColorPicker.svelte';
  import Tab from './shared/Tab.svelte';
  import { faEyeDropper, faScissors } from '@fortawesome/free-solid-svg-icons';
  import SnippingTool from './modules/snipping-tool/SnippingTool.svelte';

  let selectedTab = $state(0);

  function isTabActive(tabIndex) {
    return selectedTab === tabIndex;
  }
</script>

<main class="p-4 ax-h-[100vh]">
  <div class="flex flex-col gap-4">
    <!-- Tabs mit besserer Gestaltung -->
    <div role="tablist" class="tabs bg-[#0f172a] sticky top-0 z-10 p-2 gap-4 tabs-lifted tabs-lg">
      <Tab isSelected={isTabActive(0)} icon={faEyeDropper} tabName="Color Picker" tabIndex={0}
           bind:selectedTab={selectedTab} />
      <Tab isSelected={isTabActive(1)} icon={faScissors} tabName="Snipping Tool" tabIndex={1}
           bind:selectedTab={selectedTab} />
    </div>

    <div class="bg-base-200 w-full rounded-box p-6 shadow-md min-h-[60vh]">
      <div class={selectedTab === 0 ? "visible" : "hidden"}>
        <div class="flex justify-center items-center py-4">
          <ColorPicker />
        </div>
      </div>
      <div class={selectedTab === 1 ? "visible" : "hidden"}>
        <div class="flex justify-center items-center py-4">
          <SnippingTool />
        </div>
      </div>
    </div>
  </div>
</main>
