import axios from "axios";
import {Server} from "../config";
import {token} from "../utils/util";
//user
export const userFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addUser = (data) => axios.post(`${Server}/user/add`, data, {headers: {'token': token()}}).then(res => res.data)
export const delUser = (id) => axios.delete(`${Server}/user/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateUser = (data) => axios.put(`${Server}/user/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getUserById = (id) => axios.get(`${Server}/user/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
//loginLog
export const loginLogFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addLoginLog = (data) => axios.post(Server+"/loginLog/add", data, {headers: {'token': token()}}).then(res => res.data)
export const delLoginLog = (id) => axios.delete(Server+"/loginLog/del?id="+id, {headers: {'token': token()}}).then(res => res.data)
export const updateLoginLog = (data) => axios.put(Server+"/loginLog/update", data, {headers: {'token': token()}}).then(res => res.data)
export const getLoginLogById = (id) => axios.get(Server+"/loginLog/getById?id="+id, {headers: {'token': token()}}).then(res => res.data)

//userDetails
export const userDetailsFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addUserDetails = (data) => axios.post(Server+"/userDetails/add", data, {headers: {'token': token()}}).then(res => res.data)
export const delUserDetails = (id) => axios.delete(Server+"/userDetails/del?id="+id, {headers: {'token': token()}}).then(res => res.data)
export const updateUserDetails = (data) => axios.put(Server+"/userDetails/update", data, {headers: {'token': token()}}).then(res => res.data)
export const getUserDetailsById = (id) => axios.get(Server+"/userDetails/getById?id="+id, {headers: {'token': token()}}).then(res => res.data)
