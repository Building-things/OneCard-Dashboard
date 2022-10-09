<script>
	import Login from "./pages/Login.svelte";
	import Home from "./pages/Home.svelte";

	const pages = ["Login", "Home"]
	let active_page = pages[0]

	let user_data;
</script>

<main>
	{#if active_page === pages[0]}
		<Login on:navigate={(event)=>{active_page = "Home"; user_data=event.detail}}/>
	{:else if active_page === pages[1]}
		<Home user_data={user_data}/>
	{/if}
</main>

<style>
	main{
		background-color: var(--background-main-colour);
	}
</style>