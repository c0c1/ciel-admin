import axios from "axios";
import Router from 'next/router'
import {toast} from "react-toastify";

const service = axios.create({
    baseURL: "http://localhost:2022",
    timeout: 500
})
service.interceptors.request.use(
    config => {
        config.headers["token"] = localStorage.getItem("token")
        return config
    },
    error => {
        console.log(error)
        return Promise.reject(error)
    }
)

service.interceptors.response.use(
    response => {
        const res = response.data
        switch (res.code) {
            case 61:
                Router.push('/login')
        }
        if (res.code !== 0) {
            console.log('err:', res.msg)
            toast.error(res.msg)
            return new Error(res.msg)
        }
        return res
    },
    error => {
        return Promise.reject(error)
    }
)
export default service