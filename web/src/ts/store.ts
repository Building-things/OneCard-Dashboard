import { writable, type Writable } from "svelte/store";

//save active page to local storage so that we dont go back to login screen on refresh
//obviously non-authenticated users can access any page through modifying local storage
//but all they would see is empty data anyway so this is not a security risk.
const storedPage = localStorage.getItem("activePage");
export const page = writable(storedPage);
page.subscribe(page => {
    localStorage.setItem("activePage", page ? page : "login");
})

const storedOneCardData = localStorage.getItem("oneCardData");
export const oneCardData:Writable<OneCardData> = writable(JSON.parse(storedOneCardData));
oneCardData.subscribe(data =>{
    localStorage.setItem("oneCardData", JSON.stringify(data));
})


//user data from server
interface Transaction{
    location: String,
    amount: Number,
}
interface OneCardData{
    balances: {
        standardMealPlan: Number,
        plusMealPlan: Number,
        flex: Number,
    }
    transactions: {
        recent: Array<Transaction>
        all: Array<Transaction>
    }
}