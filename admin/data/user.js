import axios from "axios";
import {Server} from "../config";
import {token} from "../utils/util";
//user
export const userFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addUser = (data) => axios.post(`${Server}/user/add`, data, {headers: {'token': token()}}).then(res => res.data)
export const delUser = (id) => axios.delete(`${Server}/user/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateUser = (data) => axios.put(`${Server}/user/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getUserById = (id) => axios.get(`${Server}/user/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)