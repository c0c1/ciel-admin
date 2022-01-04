import {toast} from "react-toastify";
import Router from "next/router";
import {errorTime, infoTime, KeyUserToken, SuccessMsg, successTime, warningTime} from "../consts/const";
import MenuIcon from '@mui/icons-material/Menu';
import ListIcon from '@mui/icons-material/List';
import ContactPageIcon from '@mui/icons-material/ContactPage';
import SupervisorAccountIcon from '@mui/icons-material/SupervisorAccount';
import MenuBookIcon from '@mui/icons-material/MenuBook';
import AttachFileIcon from '@mui/icons-material/AttachFile';
import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet';
import LocalAtmIcon from '@mui/icons-material/LocalAtm';
import QuestionAnswerIcon from '@mui/icons-material/QuestionAnswer';
import PeopleIcon from '@mui/icons-material/People';
import FormatListNumberedRtlIcon from '@mui/icons-material/FormatListNumberedRtl';
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import CreditCardIcon from '@mui/icons-material/CreditCard';
import PaidIcon from '@mui/icons-material/Paid';
import AttachMoneyIcon from '@mui/icons-material/AttachMoney';
import SavingsIcon from '@mui/icons-material/Savings';
import AlignHorizontalLeftIcon from '@mui/icons-material/AlignHorizontalLeft';
import FormatListNumberedIcon from '@mui/icons-material/FormatListNumbered';
import AddRoadIcon from '@mui/icons-material/AddRoad';
import LeaderboardIcon from '@mui/icons-material/Leaderboard';
import CameraFrontIcon from '@mui/icons-material/CameraFront';
import SettingsIcon from "@mui/icons-material/Settings";
import AutoFixHighIcon from '@mui/icons-material/AutoFixHigh';
import BadgeIcon from '@mui/icons-material/Badge';
import BrandingWatermarkIcon from '@mui/icons-material/BrandingWatermark';
import AddShoppingCartIcon from '@mui/icons-material/AddShoppingCart';
import LiquorIcon from '@mui/icons-material/Liquor';
import ListAltIcon from '@mui/icons-material/ListAlt';
import FaceIcon from '@mui/icons-material/Face';
import FaceRetouchingNaturalIcon from '@mui/icons-material/FaceRetouchingNatural';

export const isEmpty = (value) => {
    return value === null || value === undefined || value === '' || value.length === 0
}
export const resBaseCode = (data, tid) => {
    if (!data) return true
    const {code, msg} = data
    if (code === 0) return false
    return resCodeShowMsg(code, msg, tid)
}
export const resCodeShowMsg = (code, msg, tid) => {
    if (isEmpty(tid)) {
        switch (code) {
            case 3: // error
                toast.error(msg, {autoClose: errorTime})
                return true
            case 65: //  data not found
            case 51: // Data validation failed
            case 2: // warning
                toast.warning(msg, {autoClose: warningTime})
                return true
            case 1: // info
                toast.info(msg, {autoClose: infoTime})
                return false
            case 0:
                toast.success(SuccessMsg, {autoClose: successTime})
                return false
            case -1:
                toast.error(msg, {autoClose: errorTime})
                return true
            case -2:
                toast.warning(msg, {autoClose: warningTime})
                Router.push('/login')
                return true
        }
    } else {
        switch (code) {
            case 3: // error
                toast.update(tid, {autoClose: errorTime, render: msg, type: "error", isLoading: false})
                return true
            case 65: //  data not found
            case 51: // Data validation failed
            case 2: // warning
                toast.update(tid, {autoClose: warningTime, render: msg, type: 'warning', isLoading: false})
                return true
            case 1: // info
                toast.update(tid, {autoClose: infoTime, render: msg, type: 'info', isLoading: false})
                return false
            case 0:
                toast.update(tid, {autoClose: successTime, render: SuccessMsg, type: 'success', isLoading: false})
                return false
            case -1:
                toast.update(tid, {autoClose: errorTime, render: msg, type: 'error', isLoading: false})
                return true
            case -2:
                toast.update(tid, {autoClose: warningTime, render: msg, type: 'warning', isLoading: false})
                Router.push('/login')
                return true
        }
    }
}
export const getIcon = (icon) => {
    switch (icon) {
        case "FaceRetouchingNaturalIcon":
            return <FaceRetouchingNaturalIcon/>
        case "FaceIcon":
            return <FaceIcon/>
        case "ListAltIcon":
            return <ListAltIcon/>
        case "LiquorIcon":
            return <LiquorIcon/>
        case "AddShoppingCartIcon":
            return <AddShoppingCartIcon/>
        case "BrandingWatermarkIcon":
            return <BrandingWatermarkIcon/>
        case "BadgeIcon":
            return <BadgeIcon/>
        case "AutoFixHighIcon":
            return <AutoFixHighIcon/>
        case "CameraFrontIcon":
            return <CameraFrontIcon/>
        case "LeaderboardIcon":
            return <LeaderboardIcon/>
        case "AddRoadIcon":
            return <AddRoadIcon/>
        case "FormatListNumberedIcon":
            return <FormatListNumberedIcon/>
        case "AlignHorizontalLeftIcon":
            return <AlignHorizontalLeftIcon/>
        case "SavingsIcon":
            return <SavingsIcon/>
        case "AttachMoneyIcon":
            return <AttachMoneyIcon/>
        case "PaidIcon":
            return <PaidIcon/>
        case "CreditCardIcon":
            return <CreditCardIcon/>
        case "AccountBalanceIcon":
            return <AccountBalanceIcon/>
        case "FormatListNumberedRtlIcon":
            return <FormatListNumberedRtlIcon/>
        case "PeopleIcon":
            return <PeopleIcon/>
        case "QuestionAnswerIcon":
            return <QuestionAnswerIcon/>
        case "LocalAtmIcon":
            return <LocalAtmIcon/>
        case "AccountBalanceWalletIcon":
            return <AccountBalanceWalletIcon/>
        case "AttachFileIcon":
            return <AttachFileIcon/>
        case "MenuBookIcon":
            return <MenuBookIcon/>
        case "SupervisorAccountIcon":
            return <SupervisorAccountIcon/>
        case "ContactPageIcon":
            return <ContactPageIcon/>
        case "ListIcon":
            return <ListIcon/>
        case "MenuIcon":
            return <MenuIcon/>
        case 'SettingsIcon':
            return <SettingsIcon/>
        default:
            return <SettingsIcon/>
    }
}
export const token = () => localStorage.getItem(KeyUserToken)
export const getQuery = (name) => {
    let reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    let r = window.location.search.substr(1).match(reg);
    if (r != null) return decodeURI(r[2]);
    return null;
}