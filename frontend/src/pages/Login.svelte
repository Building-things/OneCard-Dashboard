<script>
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	let usr;
	let pwd;
	async function login(){
		const url = "http://127.0.0.1:5000"
		const res = await fetch(url, {
			method: "POST", 
			body: JSON.stringify({usr: usr, pwd: pwd}),
			headers: {
        		"Content-type": "application/json; charset=UTF-8"
    		}
		})
		//TODO change if(true) to if(res.ok) when testing for release
		if(true){
			//sends our user data along
			dispatch("navigate", await res.json());
		}
	}
</script>

<main>
	<form on:submit|preventDefault={login}>
		<h1>OneCard Dashboard</h1>
		<div>
			<label for="usr">NetlinkID</label> <br>
			<input name="usr" bind:value={usr}>
		</div>
		<div>
			<label for="pwd">Password</label> <br>
			<input name="pwd" type="password"  bind:value={pwd}/> 
		</div>
		<button type="submit" class="button">Login</button>
	</form>
</main>

<style>
	form{
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: space-evenly;
		background-color: var(--background-accent-colour);
		width: 80%;
		height: 60%;
		color: white;
	}
	@media only screen and (min-width: 600px){
		form{
			max-width: 450px;
		}
	}
	main{
		height: 100vh;
		width: 100vw;

		display: flex;
		align-items: center;
		justify-content: center;
	}
	h1{
		display: block;
		font-size: 1.6rem;
	}
	input{
		background-color: var(--background-main-colour);
		border: none;
		font-size: 1rem;
		padding: 5px;
		color: white;
		margin-top: 10px;
	}
	button[type="submit"]{
		width: 80%;
		height: 32px;
		
	}
</style>