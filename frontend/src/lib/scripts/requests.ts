import { goto } from "$app/navigation"

export const BASE_PATH = "http://localhost:3000"
export const AUTH_BASE_PATH = "http://localhost:3000/arq"

export async function postRq(path: string, body: string) {

    try {
        const res = await fetch(BASE_PATH + path, {
            method: "POST",
            body: body,
            headers: {
                "Content-Type": "application/json"
            },
        })

        if(res.status != 200) {
            console.error(res.status + " " + path)
            return {success: false, message: "Ein unerwarteter Fehler ist aufgetreten. Bitte versuche es später erneut."}
        }

        const json = await res.json()
        if(!json.success) {
            console.error(path + " | " + json.error)
            return {success: false, message: json.error}
        }

        return json
    } catch (e) {
        console.error(path + " | " + e)
        return {success: false, message: "Ein unerwarteter Fehler ist aufgetreten. Bitte versuche es später erneut."}
    }
}

export async function postRqAuthorized(path: string, body: string) {

    try {
        const res = await fetch(AUTH_BASE_PATH + path, {
            method: "POST",
            body: body,
            headers: {
                "Authorization": "Bearer " + localStorage.getItem("token")!,
                "Content-Type": "application/json"
            },
        })

        if(res.status == 401) {
            console.error(res.status + " " + path)
            goto("/login");
            return {success: false, message: "Unauthorized"}
        }

        if(res.status != 200) {
            console.error(res.status + " " + path)
            return {success: false}
        }

        const json = await res.json()
        if(!json.success) {
            console.error(AUTH_BASE_PATH + path + " | " + json.error)
            return {success: false, message: json.error}
        }

        return json
    } catch (e) {
        console.error(AUTH_BASE_PATH + path + " | " + e)
        return {success: false, message: "Ein unerwarteter Fehler ist aufgetreten. Bitte versuche es später erneut."}
    }
}
