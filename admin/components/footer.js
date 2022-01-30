import use_user from "../data/use_user";
import Router from "next/router";

export default function Footer() {
    const {user, loading, loggedOut} = use_user()
    let menu = "O(∩_∩)O"
    if (loading || loggedOut || !user) {
        menu = "O(∩_∩)O"
    }
    if (loggedOut) {
        menu = "O(∩_∩)O"
    }
    if (user) {
        menu = (user.menu.map(item => {
            return (
                <li key={item.name}>
                    <mark>{item.name}</mark>
                    <ul>{(item.Children || []).map(item => <li key={item.name}><a href="#" onClick={() => Router.push(item.path)}>{item.name}</a></li>)} </ul>
                </li>
            )
        }))
    }
    return <footer>
        <ul>{menu} </ul>
    </footer>;
}