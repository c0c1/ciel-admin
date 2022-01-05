import * as React from 'react';
import {useEffect, useState} from 'react';
import {styled} from '@mui/material/styles';
import Box from '@mui/material/Box';
import MuiDrawer from '@mui/material/Drawer';
import MuiAppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import CssBaseline from '@mui/material/CssBaseline';
import Divider from '@mui/material/Divider';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import {useUser} from "./useUser";
import {ExpandLess, ExpandMore} from "@mui/icons-material";
import {Collapse, Dialog, DialogActions, DialogContent, DialogTitle, Menu, MenuItem, Switch, TextField, Typography} from "@mui/material";
import {getIcon} from "../utils/util";
import Router from "next/router";
import Button from "@mui/material/Button";
import {toast} from "react-toastify";
import {errorTime, KeyUserInfo, KeyUserToken, SubmitBtnMsg, successTime, warningTime} from "../consts/const";
import {updatePwd} from "../data/sys";

const drawerWidth = 240;
export const Layout = ({children, dark, setDark, theme}) => {
    const {u, loading, loggedOut} = useUser()
    const [drawerOpen, setDrawerDrawerOpen] = React.useState(true);
    const [pwdShow, setPwdShow] = React.useState(false);
    const [pwdForm, setPwdForm] = useState({pwd: "", pwd2: ""});
    const [menuOpen, setMenuOpen] = useState({})
    const [userOption, setUserOption] = useState(null);
    useEffect(() => {
        let menuOpenOption = {}
        if (!loading && u) {
            for (let i = 0; i < u.menus.length; i++) {
                menuOpenOption['c' + i] = false
            }
            setMenuOpen(menuOpenOption)
        }
    }, [loading, u])
    if (loading || loggedOut) {
        return <div>loading...</div>
    }
    const assistant = new PageAssistant(setPwdShow, setUserOption, setDrawerDrawerOpen, pwdForm, setPwdForm)
    return (<Box sx={{display: 'flex'}}>
        <CssBaseline/>
        <AppBar position="fixed" open={drawerOpen}>
            <Toolbar className={'d-flex'}>
                <IconButton color="inherit" onClick={assistant.showDrawer} edge="start" sx={{marginRight: '36px', ...(drawerOpen && {display: 'none'}),}}> <MenuIcon/> </IconButton>
                <Button color={'inherit'} onClick={() => Router.push("/")}>首页</Button>
                <div className={'me-auto'}/>
                <Button
                    onClick={assistant.showUserOptions}
                    color={'inherit'}
                >
                    {u.info.uname}
                </Button>
                <Menu
                    keepMounted
                    anchorEl={userOption}
                    open={Boolean(userOption)}
                    onClose={assistant.closeUserOptions}
                >
                    <MenuItem> <Typography textAlign={'center'} onClick={() => assistant.showPwd()}>修改密码</Typography> </MenuItem>
                    <MenuItem> <Typography textAlign={'center'} onClick={() => assistant.handleLogout()}>退出</Typography> </MenuItem>
                </Menu>
                <Switch size={'small'} checked={dark} onChange={() => setDark(!dark)}/>
            </Toolbar>
        </AppBar>
        <Drawer variant="permanent" open={drawerOpen}>
            <DrawerHeader>
                <Button fullWidth variant={'text'} size={"large"} color={'inherit'} style={{fontSize: '16px', fontWeight: 'bold'}} onClick={() => Router.push("/")}>SCiel</Button>
                <IconButton onClick={assistant.closeDrawer}>
                    {theme.direction === 'rtl' ? <ChevronRightIcon/> : <ChevronLeftIcon/>}
                </IconButton>
            </DrawerHeader>
            <Divider/>
            <List>
                {u.menus.map((item, index) => {
                    let filed = 'c' + index
                    return (<div key={index}>
                        <ListItem button key={item.id} onClick={() => {
                            setMenuOpen({...menuOpen, [filed]: !menuOpen[filed]})
                        }}>
                            <ListItemIcon> {getIcon(item.icon)}</ListItemIcon>
                            <ListItemText primary={item.name}/>{menuOpen[filed] ? <ExpandLess/> : <ExpandMore/>}
                        </ListItem>
                        <Collapse component={'li'} in={menuOpen[filed]} unmountOnExit timeout={'auto'}>
                            <List disablePadding className={'ms-2'}>
                                {item.Children.map((item) => item.type === 2 ? <Divider key={item.id}/> : <ListItem button key={item.id} onClick={() => Router.push(item.path)}>
                                    <ListItemIcon>{getIcon(item.icon)}</ListItemIcon>
                                    <ListItemText primary={item.name}/>
                                </ListItem>)}
                            </List>
                        </Collapse>
                        <Divider/>
                    </div>)
                })}
            </List>
        </Drawer>
        <Box component="main" sx={{flexGrow: 1, p: 3}}> <DrawerHeader/>
            {children}
        </Box>
        <Dialog open={pwdShow} onClose={assistant.closePwd}>
            <DialogTitle>修改密码</DialogTitle>
            <DialogContent>
                <TextField
                    margin={'dense'}
                    label={'密码'}
                    size={'small'}
                    fullWidth
                    value={pwdForm.pwd}
                    onChange={e => setPwdForm({...pwdForm, 'pwd': e.target.value})}
                />
                <TextField
                    label={'确认密码'} margin={'dense'} size={'small'} fullWidth
                    value={pwdForm.pwd2}
                    onChange={e => setPwdForm({...pwdForm, 'pwd2': e.target.value})}
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={() => assistant.submitPwd()}>{SubmitBtnMsg}</Button>
            </DialogActions>
        </Dialog>
    </Box>);
}

class PageAssistant {
    constructor(setPwdShow, setUserOption, setDrawerDrawerOpen, pwdForm, setPwdForm) {
        this.setPwdShow = setPwdShow
        this.setUserOption = setUserOption
        this.setDrawerDrawerOpen = setDrawerDrawerOpen
        this.pwdForm = pwdForm
        this.setPwdForm = setPwdForm
    }

    showDrawer = () => this.setDrawerDrawerOpen(true)
    closeDrawer = () => this.setDrawerDrawerOpen(false)
    showUserOptions = (event) => this.setUserOption(event.currentTarget)
    closeUserOptions = () => this.setUserOption(null)
    closePwd = () => this.setPwdShow(false)
    showPwd = () => {
        this.setPwdShow(true)
        this.setPwdForm({pwd: '', pwd2: ''})
        this.closeUserOptions()
    }
    handleLogout = () => {
        localStorage.removeItem(KeyUserToken)
        localStorage.removeItem(KeyUserInfo)
        Router.push('/login')
    }

    async submitPwd() {
        const {pwd, pwd2} = this.pwdForm
        if (pwd !== pwd2) {
            toast.warning('两次密码不相同', {autoClose: warningTime})
            return
        }
        const {code, message} = await updatePwd(pwd, pwd2)
        if (code !== 0) {
            toast.error(message, {autoClose: errorTime})
            return
        }
        toast.success("ok", {autoClose: successTime})
        this.setPwdShow(false)
    }
}


const openedMixin = (theme) => ({
    width: drawerWidth, transition: theme.transitions.create('width', {
        easing: theme.transitions.easing.sharp, duration: theme.transitions.duration.enteringScreen,
    }), overflowX: 'hidden',
});

const closedMixin = (theme) => ({
    transition: theme.transitions.create('width', {
        easing: theme.transitions.easing.sharp, duration: theme.transitions.duration.leavingScreen,
    }), overflowX: 'hidden', width: `calc(${theme.spacing(7)} + 1px)`, [theme.breakpoints.up('sm')]: {
        width: `calc(${theme.spacing(9)} + 1px)`,
    },
});

const DrawerHeader = styled('div')(({theme}) => ({
    display: 'flex', alignItems: 'center', justifyContent: 'flex-end', padding: theme.spacing(0, 1), // necessary for content to be below app bar
    ...theme.mixins.toolbar,
}));

const AppBar = styled(MuiAppBar, {
    shouldForwardProp: (prop) => prop !== 'open',
})(({theme, open}) => ({
    zIndex: theme.zIndex.drawer + 1, transition: theme.transitions.create(['width', 'margin'], {
        easing: theme.transitions.easing.sharp, duration: theme.transitions.duration.leavingScreen,
    }), ...(open && {
        marginLeft: drawerWidth, width: `calc(100% - ${drawerWidth}px)`, transition: theme.transitions.create(['width', 'margin'], {
            easing: theme.transitions.easing.sharp, duration: theme.transitions.duration.enteringScreen,
        }),
    }),
}));

const Drawer = styled(MuiDrawer, {shouldForwardProp: (prop) => prop !== 'open'})(({theme, open}) => ({
    width: drawerWidth, flexShrink: 0, whiteSpace: 'nowrap', boxSizing: 'border-box', ...(open && {
        ...openedMixin(theme), '& .MuiDrawer-paper': openedMixin(theme),
    }), ...(!open && {
        ...closedMixin(theme), '& .MuiDrawer-paper': closedMixin(theme),
    }),
}),);