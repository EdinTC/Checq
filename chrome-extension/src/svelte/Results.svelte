<script>
    import ListItem from "./ListItem.svelte";
    export let data;
</script>

{#if data !== undefined}
    {#if data.ip !== undefined}
        <h2 class="font-sans text-xl text-green-600 font-bold mb-2">
            List of IP records:
        </h2>
        <ul class="list-none mb-6 list-inside">
            {#each data.ip as ip}
                <ListItem>{ip}</ListItem>
            {/each}
        </ul>
    {/if}

    {#if data.ns !== undefined}
        <h2 class="font-sans text-xl text-green-600 font-bold mb-2">
            List of NS records:
        </h2>
        <ul class="list-none mb-6 list-inside">
            {#each data.ns as ns}
                <ListItem>{ns}</ListItem>
            {/each}
        </ul>
    {/if}

    {#if data.hostname !== undefined}
        <h2 class="font-sans text-xl text-green-600 font-bold mb-2">Hostname:</h2>
        <ul class="list-none mb-6 list-inside">
            <ListItem>{data.hostname}</ListItem>
        </ul>
    {/if}

    {#if data.txt !== undefined}
        <h2 class="font-sans text-xl text-green-600 font-bold mb-2">
            List of TXT records:
        </h2>
        <ul class="list-none mb-6 list-inside">
            {#each data.txt as txt}
                <ListItem>{txt}</ListItem>
            {/each}
        </ul>
    {/if}
{/if}