<script lang="ts">
    import {oneCardData, page} from "../ts/store"

    async function login(event: SubmitEvent){
        const formData = new FormData(event.target)
        const serverURL = "http://127.0.0.1:5000";
        const res = await fetch(serverURL, {
            method: "POST",
            body: formData,
        })
        if(res.ok){
            oneCardData.set(await res.json());
            console.log($oneCardData);
            page.set("home");
        }else{
            //TODO ERROR
        }
    }

    
</script>

<main>
    <h2>OCD</h2>
    <div class="flex-column" style="height: 100vh; justify-content:center">
        <form class="flex-column" on:submit|preventDefault={login}>
            <div>
                <label for="username">NetlinkID</label>
                <input name="username">
            </div>
            <div>
                <label for="password">Password</label>
                <input type="password" name="password">
            </div>
            <button class="button" type="submit">Submit</button>
        </form>
    </div>
</main>

<style>
    .flex-column{
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    main{
        height: 100vh;
        overflow-y: hidden;
    }
    form{
        background-color: var(--background-accent-colour);
        width: 90%;
        box-sizing: border-box;
        padding: 10%;
    }
    form > *{
        margin-bottom: 10%;
    }
    button{
        width: 100%;
    }
    input{
        box-sizing: border-box;
        font-size: 1rem;
        padding: 5px;
        width: 100%;
    }
    @media (min-width: 600px){
        form{
            width: 600px;
        }
    }
</style>