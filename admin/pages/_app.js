import '../styles/globals.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-toastify/dist/ReactToastify.css';
import {ToastContainer} from "react-toastify";
import {createTheme, ThemeProvider} from "@mui/material";
import {useEffect, useState} from "react";
import {Layout} from "../components/Layout";
import {useUser} from "../components/useUser";
import Router from "next/router";


function MyApp({Component, pageProps: {pageProps}, router}) {
    const {u, loading, loggedOut} = useUser()
    const [dark, setDark] = useState(true);
    const theme = createTheme({
        palette: {mode: dark ? 'dark' : 'light'}
    })
    const handleSetDark = () => {
        setDark(!dark)
        localStorage.setItem('dark', !dark)
    }
    useEffect(() => {
        setDark(localStorage.getItem('dark') === 'true')
        if (loggedOut) Router.replace('/login')
    }, [loggedOut, dark])
    if (loading) return <>loading...</>
    if (router.route === '/login') return <> <Component {...pageProps} /></>
    return <ThemeProvider theme={theme}>
        <Layout dark={dark} setDark={handleSetDark} theme={theme}> <Component {...pageProps} userInfo={u}/> </Layout>
        <ToastContainer position={'top-center'} theme={dark ? 'dark' : 'light'}/>
    </ThemeProvider>
}

export default MyApp
