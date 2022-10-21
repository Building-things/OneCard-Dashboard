import { writable, type Writable } from "svelte/store";

//save active page to local storage so that we dont go back to login screen on refresh
//obviously non-authenticated users can access any page through modifying local storage
//but all they would see is empty data anyway so this is not a security risk.
const storedPage = localStorage.getItem("activePage");
export const page = writable(storedPage === "login" || storedPage === "home" ? storedPage : "login");
page.subscribe(page => {
    localStorage.setItem("activePage", page);
})

const storedOneCardData = localStorage.getItem("oneCardData");
let storedData: OneCardData = {Balances: null, Transactions: null}
try{
    storedData = JSON.parse(storedOneCardData)
}catch(err){}
export const oneCardData:Writable<OneCardData> = writable(storedData);
oneCardData.subscribe(data =>{
    localStorage.setItem("oneCardData", JSON.stringify(data));
})


//user data from server
export interface Transaction{
    Location: string,
    Amount: Number,
    Date: string,
}
interface OneCardData{
    Balances: {
        StandardMealPlan: Number,
        PlusMealPlan: Number,
        Flex: Number,
    } | null
    Transactions: {
        Recent: Array<Transaction>
        All: Array<Transaction>
    } | null
}