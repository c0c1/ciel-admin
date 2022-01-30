import 'bootstrap/dist/css/bootstrap.min.css'
import 'react-toastify/dist/ReactToastify.css';
import '../styles/globals.scss'
import {ToastContainer} from "react-toastify";
import Footer from "../components/footer";

function MyApp({Component, pageProps, router}) {
    let pathname = router.pathname;
    if (pathname === '/login' || pathname === '/') {
        return <div className={'w'}>
            <Component {...pageProps} />
            <ToastContainer position={"bottom-left"} autoClose={2000}/>
            {/*https://fkhadra.github.io/react-toastify/introduction/*/}
        </div>
    }
    return <div className={'w'}>
        <Component {...pageProps} />
        <ToastContainer position={"bottom-left"} autoClose={2000}/>
        <Footer/>
    </div>
}

export default MyApp
