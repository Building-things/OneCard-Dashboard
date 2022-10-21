<script lang="ts">
    import {oneCardData, page} from "../ts/store"

    let formError = false;

    async function login(event: any){
        //if this data is populated no need to log back in
        if($oneCardData?.Balances?.StandardMealPlan){
            page.set("home");
            return;
        }
        const formData = new FormData(event.target) 
        const serverURL = "http://127.0.0.1:5000";
        const res = await fetch(serverURL, {
            method: "POST",
            body: formData,
        })
        if(res.ok){
            let data = await res.json()
            oneCardData.set(data);
            page.set("home");
        }else{
            formError = true;
        }
    }


    
</script>

<main>
    <h1>OCD</h1>
    <div class="flex-column" style="height: 80vh; justify-content:center">
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
            <div class="{formError ? "" : "hidden"} error">
                <p>Your NetlinkID or Password Is Invalid</p>
            </div>  
        </form>
    </div>
</main>

<style>
    .flex-column{
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    .hidden{
        visibility: hidden;
    }
    h1{
        display: block;
        box-sizing: border-box;
        padding: 2%;
    }
    main{
        height: 100vh;
        overflow-y: hidden;
    }
    form{
        background-color: var(--background-accent-colour);
        width: 80%;
        box-sizing: border-box;
        padding: 10%;
    }
    form > *{
        margin-bottom: 15%;
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
    label{
        font-size: 1.2rem;
    }
    @media (min-width: 600px){
        form{
            width: 600px;
        }
    }
    .error{
        box-sizing: border-box;
        padding: 10%;
        background-color: rgba(255, 0, 0, 0.5);
        border: 1px solid red;
        margin-bottom: 0;
    }
</style>