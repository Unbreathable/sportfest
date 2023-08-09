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
            return {success: false}
        }

        const json = await res.json()
        if(!json.success) {
            console.error(path + " | " + json.error)
            return {success: false, message: json.error}
        }

        return json
    } catch (e) {
        console.error(path + " | " + e)
        return {success: false}
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
        console.error("/a" + path + " | " + e)
        return {success: false}
    }
}
