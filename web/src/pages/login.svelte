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
        for(const val of formData.values()){
            if(!val){
                formError = true;
                return;
            }
        }
        const serverURL = window.location.href;
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
    <h1>ONECard Dashboard</h1>
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
    <svg id="bottomArt" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320"><path fill="var(--accent-colour)" fill-opacity="1" d="M0,192L0,224L160,224L160,128L320,128L320,96L480,96L480,96L640,96L640,224L800,224L800,96L960,96L960,0L1120,0L1120,0L1280,0L1280,128L1440,128L1440,320L1280,320L1280,320L1120,320L1120,320L960,320L960,320L800,320L800,320L640,320L640,320L480,320L480,320L320,320L320,320L160,320L160,320L0,320L0,320Z"></path></svg>
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
        width: 100%;
        text-align: center;
        box-sizing: border-box;
        padding: 2%;
        background-color: var(--accent-colour);
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
        margin-bottom: 10%;
    }
    button{
        width: 100%;
    }
    input{
        box-sizing: border-box;
        font-size: 1.2rem;
        outline: none;
        padding: 7px;
        width: 100%;
        background-color: var(--background-main-colour);
        border-bottom: 3px solid var(--background-main-colour);
        color: white;
        margin-top: 5px;
        transition: border-bottom 0.5s linear;
    }
    input:focus{
        border-bottom: 3px solid var(--accent-colour);
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
        padding: 5%;
        background-color: rgba(255, 0, 0, 0.5);
        border: 1px solid red;
        margin-bottom: 0;
    }
    #bottomArt{
        position: absolute;
        bottom: 0;
    }
</style>