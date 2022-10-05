<script>
	import Login from "./pages/Login.svelte";
	import Home from "./pages/Home.svelte";

	const pages = ["Login", "Home"]
	let active_page = pages[0]
</script>


<navbar>
	<button on:click={()=>{active_page = pages[1]}}>Home</button>
	<button on:click={()=>{active_page = pages[0]}}>Login</button>
</navbar>
<main>
	{#if active_page === pages[0]}
		<Login/>
	{:else if active_page === pages[1]}
		<Home/>
	{/if}
</main>

<style>
	
</style>