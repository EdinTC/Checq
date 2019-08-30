<script>
	let promise;
	const apiUrl = 'https://api.checq.intercube.io:1337'
	let hostname = '';

	async function fetchDomain () {
		if( hostname === '') return;
		const response = await fetch(`${apiUrl}/${hostname}`, { mode: 'cors'});
		const json = await response.json();
		if (response.ok) {
			return json;
		} else {
			throw new Error(json);
		}
	}

	chrome.tabs.query({'active': true, 'currentWindow': true}, (tabs) => {
    if(tabs === undefined) return;
		var url = new URL(tabs[0].url);
		hostname = url.hostname;
		promise = fetchDomain();
	});
</script>

<style>

pre {
  margin-right: .50rem; 
}
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
<div class="container mx-auto px-4 pl-2">
  {#await promise}
    <div class="flex justify-center">
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
            <pre class="whitespace-normal bg-gray-200 overflow-x-auto p-3 mb-3 text-xs">{txt}</pre>
          {/each}
        </ul>
      {/if}
    {/if}
  {:catch error}
    <p style="color: red">{error.message}</p>
  {/await}
</div>