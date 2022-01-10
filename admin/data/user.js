import axios from "axios";
import {LoginUrl, Server} from "../config";
import {token} from "../utils/util";
import Router from "next/router";
//user
export const userFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const addUser = (data) => axios.post(`${Server}/user/add`, data, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const delUser = (id) => axios.delete(`${Server}/user/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const updateUser = (data) => axios.put(`${Server}/user/update`, data, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const getUserById = (id) => axios.get(`${Server}/user/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
//loginLog
export const loginLogFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const addLoginLog = (data) => axios.post(Server + "/loginLog/add", data, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const delLoginLog = (id) => axios.delete(Server + "/loginLog/del?id=" + id, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const updateLoginLog = (data) => axios.put(Server + "/loginLog/update", data, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const getLoginLogById = (id) => axios.get(Server + "/loginLog/getById?id=" + id, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))

//userDetails
export const userDetailsFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const addUserDetails = (data) => axios.post(Server + "/userDetails/add", data, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const delUserDetails = (id) => axios.delete(Server + "/userDetails/del?id=" + id, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const updateUserDetails = (data) => axios.put(Server + "/userDetails/update", data, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
export const getUserDetailsById = (id) => axios.get(Server + "/userDetails/getById?id=" + id, {headers: {'token': token()}}).then(res => res.data).catch(() => Router.push(LoginUrl))
