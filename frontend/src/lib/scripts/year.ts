import { writable } from "svelte/store";
import { postRqAuthorized } from "./requests";

export const yearLoading = writable(true);
export const yearError = writable("");
export const currentYear = writable({
    id: -1,
    name: "",
});

export async function loadCurrentYear(id: any, currentId: number) {
    yearLoading.set(true);
        
    let yearId = 0;
    try {
        yearId = parseInt(id);
    } catch(e) {
        yearError.set("Dieser Jahrgang konnte nicht geladen werden!");
        return
    }

    if(currentId == yearId) {
        yearLoading.set(false);
        return;
    }

    const res = await postRqAuthorized("/years/get", JSON.stringify({
        "id": yearId
    }));
    if(!res.success) {
        yearError.set(res.message);
        return;
    }

    yearLoading.set(false)
    currentYear.set(res.obj);
}