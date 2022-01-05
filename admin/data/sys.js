import axios from "axios";
import {Server} from "../config";
import {token} from "../utils/util";


//menu
export const menuFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addMenu = (data) => axios.post(`${Server}/menu/add`, data, {headers: {'token': token()}}).then(res => res.data)
export const delMenu = (id) => axios.delete(`${Server}/menu/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateMenu = (data) => axios.put(`${Server}/menu/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getMenuById = (id) => axios.get(`${Server}/menu/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)

//api
export const apiFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addApi = (data) => axios.post(`${Server}/api/add`, data, {headers: {'token': token()}}).then(res => res.data)
export const delApi = (id) => axios.delete(`${Server}/api/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateApi = (data) => axios.put(`${Server}/api/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getApiById = (id) => axios.get(`${Server}/api/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)

//role
export const roleFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addRole = (data) => axios.post(`${Server}/role/add`, data, {headers: {'token': token()}}).then(res => res.data)
export const delRole = (id) => axios.delete(`${Server}/role/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateRole = (data) => axios.put(`${Server}/role/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getRoleById = (id) => axios.get(`${Server}/role/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)

//roleMenu
export const roleMenuFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addRoleMenu = (data) => axios.post(`${Server}/roleMenu/add`, data, {headers: {'token': token()}}).then(res => res.data)
export const delRoleMenu = (id) => axios.delete(`${Server}/roleMenu/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateRoleMenu = (data) => axios.put(`${Server}/roleMenu/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getRoleMenuById = (id) => axios.get(`${Server}/roleMenu/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const noMenus = (rid) => axios.get(`${Server}/roleMenu/noMenus?rid=${rid}`, {headers: {'token': token()}}).then(res => res.data)

//roleApi
export const roleApiFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addRoleApi = (data) => axios.post(`${Server}/roleApi/add`, data, {headers: {'token': token()}}).then(res => res.data)
export const delRoleApi = (id) => axios.delete(`${Server}/roleApi/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateRoleApi = (data) => axios.put(`${Server}/roleApi/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getRoleApiById = (id) => axios.get(`${Server}/roleApi/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const noApis = (rid) => axios.get(`${Server}/roleApi/noApis?rid=${rid}`, {headers: {'token': token()}}).then(res => res.data)

//admin
export const login = (data) => axios.post(`${Server}/admin/login`, data)
export const adminFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addAdmin = (data) => axios.post(`${Server}/admin/add`, data, {headers: {'token': token()}}).then(res => res.data)
export const delAdmin = (id) => axios.delete(`${Server}/admin/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateAdmin = (data) => axios.put(`${Server}/admin/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getAdminById = (id) => axios.get(`${Server}/admin/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updatePwd = (pwd, pwd2) => axios.put(`${Server}/admin/pwd`, {'pwd': pwd, 'pwd2': pwd2}, {headers: {'token': token()}}).then(res => res.data)

//dict
export const dictFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addDict = (data) => axios.post(`${Server}/dict/add`, data, {headers: {'token': token()}}).then(res => res.data)
export const delDict = (id) => axios.delete(`${Server}/dict/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateDict = (data) => axios.put(`${Server}/dict/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getDictById = (id) => axios.get(`${Server}/dict/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)

//file
export const fileFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const delFile = (id) => axios.delete(`${Server}/file/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateFile = (data) => axios.put(`${Server}/file/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getFileById = (id) => axios.get(`${Server}/file/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const uploadFile = (data) => axios.post(`${Server}/file/upload`, data, {headers: {'token': token()}}).then(res => res.data)

//icon
export const iconFetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const addIcon = (data) => axios.post(`${Server}/icon/add`, data, {headers: {'token': token()}}).then(res => res.data)
export const delIcon = (id) => axios.delete(`${Server}/icon/del?id=${id}`, {headers: {'token': token()}}).then(res => res.data)
export const updateIcon = (data) => axios.put(`${Server}/icon/update`, data, {headers: {'token': token()}}).then(res => res.data)
export const getIconById = (id) => axios.get(`${Server}/icon/getById?id=${id}`, {headers: {'token': token()}}).then(res => res.data)