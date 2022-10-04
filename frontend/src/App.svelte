<script>
	let data;

	let usr;
	let pwd;

	async function submit(){
		const url = "http://127.0.0.1:5000"
		const res = await fetch(url, {
			method: "POST", 
			body: JSON.stringify({usr: usr, pwd: pwd}),
			headers: {
        		"Content-type": "application/json; charset=UTF-8"
    		}
		})
		
		data = await res.text()
	}
</script>

<main bind:innerHTML={data} contenteditable="true"></main>
<p>Username:</p>
<input bind:value={usr}>
<p>Password:</p>
<input bind:value={pwd} type="password" /> <br>
<button on:click={submit}>Submit</button>


<style>
	
</style>