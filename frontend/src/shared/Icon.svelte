<script>
    import {icon} from '@fortawesome/fontawesome-svg-core'
    export let icon_definition

    $: renderedIcon = icon(icon_definition)
    $: html = renderedIcon ? renderedIcon.html[0] : ''
</script>

<span class="fa-icon">
    {@html html}
</span>

<style>
    .fa-icon {
        display: inline-block;
    }
</style>