import Meta from "../components/meta";
import {Nav} from "../components/nav";
import {useState} from "react";
import {login} from "../api/sys";
import Router from "next/router";

export default function Login() {
    const [msg, setMsg] = useState("");
    const [uname, setUname] = useState("");
    const [pwd, setPwd] = useState("");
    const handleLogin = async (e) => {
        e.preventDefault()
        const {code, data, msg} = await login({uname, pwd})
        if (code !== 0) {
            setMsg(msg)
            return
        }
        setMsg("")
        const {menus, token, u} = data
        localStorage.setItem("token", token)
        localStorage.setItem("menu", JSON.stringify(menus))
        localStorage.setItem("u", JSON.stringify(u))
        Router.push("/")
    }

    return <>
        <Meta/>
        <Nav/>
        <main>
            <h1 className={'hr'}>Admin</h1>
            <form className={'data-details'}>
                <h2>Login</h2>
                <p>username <input value={uname} onChange={e => setUname(e.target.value)}/>
                    <span className={'text-danger'} hidden={msg === ""}>{msg}</span>
                </p>
                <p>password <input value={pwd} type={'password'} onChange={e => setPwd(e.target.value)}/></p>
                <input type={'submit'} value={'submit'} onClick={e => handleLogin(e)}/>
            </form>
            <p className={'hr'}/>
        </main>
    </>
}