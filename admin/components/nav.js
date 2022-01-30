import Router from "next/router";
import use_user from "../data/use_user";
import request from "../utils/request";
import {toast} from "react-toastify";

export function Nav() {
    const {user, loading, loggedOut, mutate} = use_user()
    let profile = "loading..."
    let menu = "loading..."
    const updatePwd = (pwd) => request({url: "/admin/pwd", method: 'put', data: {pwd: pwd}})
    const handleLogout = () => {
        localStorage.removeItem("u")
        localStorage.removeItem("token")
        localStorage.removeItem("menu")
        Router.push("/login")
        mutate()
    }
    const initData = () => {
        profile = (<li><a href={"#"} onClick={() => Router.replace('/login')}>登录</a></li>)
        menu = ""
    }
    const handleUpdatePwd = async () => {
        let pwd = prompt('请输入新的密码');
        if (pwd) {
            let {code, msg} = await updatePwd(pwd);
            toast.success(msg)
        }
    }
    if (loading) {
        profile = "loading..."
    }
    if (user) {
        profile = (<li className={'profile'}><a href={"#"}>{user.u.uname}</a>
            <ul>
                <li><a href="#" onClick={() => handleUpdatePwd()}>修改密码</a></li>
                <li><a href="#" onClick={() => handleLogout()}>退出</a></li>
            </ul>
        </li>)
        menu = (user.menu.map(item => {
            return (
                <li key={item.name}>
                    <a href={'#'}>{item.name}</a>
                    <ul>{(item.Children || []).map(item => <li key={item.name}><a href="#" onClick={() => Router.push(item.path)}>{item.name}</a></li>)} </ul>
                </li>
            )
        }))
    }
    if (loggedOut) initData()
    return <nav>
        <ul>
            <li><a href="#" onClick={() => Router.push('/')}>Home</a></li>
            {menu}
            {profile}
        </ul>
    </nav>;
}