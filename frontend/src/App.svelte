<script>
  import { onMount } from 'svelte';
  let promise;
  const apiUrl = "https://api.checq.intercube.io";
  let hostname = "";

  onMount(async () => {
    promise = fetchDomain();
  });
  
  function extractHostname(url) {
    let hostname;

    //find & remove protocol (http, ftp, etc.) and get hostname
    if (url.indexOf("//") > -1) {
        hostname = url.split('/')[2];
    }
    else {
        hostname = url.split('/')[0];
    }

    //find & remove port number
    hostname = hostname.split(':')[0];
    //find & remove "?"
    hostname = hostname.split('?')[0];

    return hostname;
  }

  async function fetchDomain() {
    if (hostname === "") return;
    hostname = extractHostname(hostname);
    const response = await fetch(`${apiUrl}/${hostname}`, { mode: "cors" });
    const json = await response.json();
    if (response.ok) {
      return json;
    } else {
      throw new Error(json);
    }
  }

  function handleClick() {
    promise = fetchDomain();
  }
</script>

<style lang="postcss">
  @import "tailwindcss/base";
  @import "tailwindcss/components";
  @import "tailwindcss/utilities";

  div.loader {
    position: relative;
    width: 2em;
    height: 2em;
    border: 3px solid #68d391;
    overflow: hidden;
    animation: spin 3s ease infinite;
  }

  div.loader::before {
    content: '';
    position: absolute;
    top: -3px;
    left: -3px;
    width: 2em;
    height: 2em;
    background-color: #48bb78;
    transform-origin: center bottom;
    transform: scaleY(1);
    animation: fill 3s linear infinite;
  }

  @keyframes spin {
    50%,
    100% {
      transform: rotate(360deg);
    }
  }

  @keyframes fill {
    25%,
    50% {
      transform: scaleY(0);
    }
    100% {
      transform: scaleY(1);
    }
  }
</style>

<div class="flex mb-4">
  <div class="container mx-auto p-4 text-center">
    <h1 class="font-sans text-2xl font-bold">🏁 ~ Checq ~ 🏁</h1>
  </div>
</div>
<div class="flex mb-4 container mx-auto ">
  <form class="w-full max-w-sm mx-auto" on:submit|preventDefault={handleClick}>
    <div class="flex items-center border-b border-b-2 border-green-400 py-2">
      <input
        class="appearance-none bg-transparent border-none w-full text-gray-700
        mr-3 py-1 px-2 leading-tight focus:outline-none"
        bind:value={hostname}
        type="text" 
        placeholder="www.google.com"
        aria-label="Full name" />
      <button
        class="flex-shrink-0 bg-green-400 hover:bg-green-500 border-green-400
        hover:border-green-500 text-sm border-4 text-white py-1 px-2 rounded focus:outline-none focus:shadow-outline focus:border-transparent"
        on:click={handleClick}
        type="button">
        Checq now
      </button>
    </div>
    <small class="text-xs">Enter a valid domainname or IP address</small>
  </form>
</div>

<div class="container mx-auto px-4 pl-2">
  {#await promise}
    <div class="flex justify-center mt-12">
      <div class="loader"></div>
    </div>
  {:then data}
    {#if data !== undefined}
      {#if data.ip !== undefined}
        <h2 class="font-sans text-xl text-green-600 font-bold mb-2">
          List of IP records:
        </h2>
        <ul class="list-none mb-6 list-inside">
          {#each data.ip as ip}
            <pre class="whitespace-normal bg-gray-200 overflow-x-auto p-3 mb-3">{ip}</pre>
          {/each}
        </ul>
      {/if}

      {#if data.ns !== undefined}
        <h2 class="font-sans text-xl text-green-600 font-bold mb-2">
          List of NS records:
        </h2>
        <ul class="list-none mb-6 list-inside">
          {#each data.ns as ns}
            <pre class="whitespace-normal bg-gray-200 overflow-x-auto p-3 mb-3">{ns}</pre>
          {/each}
        </ul>
      {/if}

      {#if data.hostname !== undefined}
        <h2 class="font-sans text-xl text-green-600 font-bold mb-2">Hostname:</h2>
        <ul class="list-none mb-6 list-inside">
          <pre class="whitespace-normal bg-gray-200 overflow-x-auto p-3 mb-3">{data.hostname}</pre>
        </ul>
      {/if}

      {#if data.txt !== undefined}
        <h2 class="font-sans text-xl text-green-600 font-bold mb-2">
          List of TXT records:
        </h2>
        <ul class="list-none mb-6 list-inside">
          {#each data.txt as txt}
            <pre class="whitespace-normal bg-gray-200 overflow-x-auto p-3 mb-3">{txt}</pre>
          {/each}
        </ul>
      {/if}
    {/if}
  {:catch error}
    <p style="color: red">{error.message}</p>
  {/await}
</div>
