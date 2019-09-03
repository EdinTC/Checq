<script>
  import Results from "./Results.svelte";
	let promise;
	const apiUrl = 'https://api.checq.intercube.io'
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

<div class="flex mb-4">
  <div class="container mx-auto p-4 text-center">
    <h1 class="font-sans text-2xl font-bold">🔍 Checq</h1>
  </div>
</div>
<div class="container mx-auto px-4 pl-2">
  {#await promise}
    <div class="flex justify-center mt-12">
      <div class="loader"></div>
    </div>
  {:then data}
    <Results {data}/>
  {:catch error}
    <p style="color: red">{error.message}</p>
  {/await}
</div>

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