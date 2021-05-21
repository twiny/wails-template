<script>
    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();

    export let name = '';
</script>

<div class="inline-block rounded-lg">
    <button
        on:click="{() => {
            dispatch('click');
        }}"
        class="hover:shadow-lg"
    >
        <div>
            <slot />
            <p class="mt-2 text-center text-sm font-semibold text-gray-100">
                {name}
            </p>
        </div>
    </button>
</div>
