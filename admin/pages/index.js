import Meta from "../components/meta";
import {Nav} from "../components/nav";

export default function Home() {
    return (
        <>
            <Meta/>
            <Nav/>
            <main>
                <h1 className={'hr'}>Admin</h1>
                <h2>hello</h2>
                <p>欢迎来到我の后台系统</p>
            </main>
        </>
    )
}