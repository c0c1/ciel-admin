import useSWR from "swr";

export default function use_user() {
    const {data, mutate, error} = useSWR("api_user", userInfo)
    const loading = !data && !error;
    const loggedOut = error && error.status === 403;
    return {loading, loggedOut, user: data, mutate}
}

const userInfo = async () => {
    let u = JSON.parse(localStorage.getItem("u"))
    let menu = JSON.parse(localStorage.getItem("menu"))
    if (!u || !menu) {
        const error = new Error("Not authorized!")
        error.status = 403
        throw error
    }
    return {u, menu}
}