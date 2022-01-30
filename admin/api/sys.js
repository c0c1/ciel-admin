import request from "../utils/request";

export const login = (data) => request({url: "/admin/login", method: 'post', data})
export const userInfo = async () => {
    let u = JSON.parse(localStorage.getItem("u"))
    let menu = JSON.parse(localStorage.getItem("menu"))
    if (!u || !menu) {
        const error = new Error("Not authorized!")
        error.status = 403
        throw error
    }
    return {u, menu}
}