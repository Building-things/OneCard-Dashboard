<script>
	import Login from "./pages/Login.svelte";
	import Home from "./pages/Home.svelte";

	const pages = ["Login", "Home"]
	let active_page = pages[0]
</script>

<main>
	{#if active_page === pages[0]}
		<Login/>
	{:else if active_page === pages[1]}
		<Home/>
	{/if}
</main>

<style>
	main{
		background-color: var(--background-main-colour);
	}
</style>