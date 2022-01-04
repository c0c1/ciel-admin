import React, {useState} from "react";
import {toast, ToastContainer} from "react-toastify";
import {Grid, TextField, Typography} from "@mui/material";
import Paper from '@mui/material/Paper';
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import {login} from "../data/sys";
import {KeyUserInfo, KeyUserToken, successTime} from "../consts/const";
import Router from "next/router";

const Login = () => {
    const [uname, setUname] = useState('');
    const [pwd, setPwd] = useState('');
    const submitLogin = async (e) => {
        e.preventDefault()
        const res = await login({uname: uname, pwd: pwd})
        const {code, message, data} = await res.data
        if (code !== 0) {
            toast.error(message)
        } else {
            const {menus, u, token} = data.normal_data
            let user = {menus: menus, info: u}
            localStorage.setItem(KeyUserInfo, JSON.stringify(user))
            localStorage.setItem(KeyUserToken, token)
            toast.success('success', {
                autoClose: successTime, onClose: () => Router.push('/')
            })
        }
    }

    return (<Grid container component={'main'} sx={{height: '100vh'}}>
        <Grid item xs={false} sm={4} md={9} sx={{backgroundImage: 'url(https://source.unsplash.com/random)', backgroundRepeat: 'no-repeat', backgroundSize: 'cover', backgroundPosition: 'center'}}/>
        <Grid item xs={12} sm={8} md={3} component={Paper} elevation={6} square>
            <Box sx={{my: 8, mx: 4, display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
                <Typography component="h1" variant="h5"> 管理员登录 </Typography>
                <Box component="form" noValidate onSubmit={submitLogin} sx={{mt: 1}}>
                    <TextField margin="normal" fullWidth label="用户名" value={uname} onChange={e => setUname(e.target.value)}/>
                    <TextField margin="normal" fullWidth label="密码" value={pwd} onChange={e => setPwd(e.target.value)} type={'password'}/>
                    <Button type="submit" fullWidth variant="contained" sx={{mt: 3, mb: 2}}> 登录 </Button>
                </Box>
            </Box>
        </Grid>
        <ToastContainer position={'top-center'} theme={'dark'}/>
    </Grid>)
}
export default Login