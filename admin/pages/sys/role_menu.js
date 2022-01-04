import useSWR from "swr";
import {useEffect, useState} from "react";
import {toast} from "react-toastify";
import {ButtonGroup, Card, CardContent, CardHeader, TableCell, TableRow, TextField} from "@mui/material";
import {PageCardActions, PageDialog, PageTable} from "../../components/page";
import {ConfirmMsg, DefaultPageParam, LoadingMsg} from "../../consts/const";
import {getQuery, resBaseCode, resCodeShowMsg} from "../../utils/util";
import {Server} from "../../config";
import {addRoleMenu, delRoleMenu, getRoleMenuById, noMenus, roleMenuFetcher, updateRoleMenu} from "../../data/sys";
import Button from "@mui/material/Button";

export default function RoleMenu() {
    const [params, setParams] = useState({...DefaultPageParam, rid: '', mid: ''})
    const [form, setForm] = useState({add: true})
    const [show, setShow] = useState(false);
    const [menuOptions, setMenuOptions] = useState([]);
    useEffect(() => {
        setParams({...params, rid: getQuery('rid')})
    }, [setParams])
    const {data, mutate} = useSWR({url: `${Server}/roleMenu/list`, params: params}, roleMenuFetcher)
    if (resBaseCode(data)) return <div/>
    const handleClose = () => setShow(false);
    const handleChange = (e, name) => setForm({
        ...form, [name]: Array.from(e.target.selectedOptions, option => {
            if (option.value !== "") {
                return option.value
            }
        })
    })
    const handleShow = async (add, id) => {
        let rid = getQuery('rid');
        const res = await noMenus(rid)
        const {normal_data} = res.data
        setMenuOptions(normal_data ? normal_data : [])
        setForm({add: add, rid: rid, mid: []})
        if (!add) {
            const {data} = await getRoleMenuById(id)
            if (resBaseCode(data)) return
            setForm({...data.normal_data})
        }
        setShow(true)
    }
    const handleSubmit = async () => {
        const tid = toast.loading(LoadingMsg);
        if (form.add) {
            const {code, message} = await addRoleMenu(form)
            if (resCodeShowMsg(code, message, tid)) return
        } else {
            const {code, message} = await updateRoleMenu(form)
            if (resCodeShowMsg(code, message, tid)) return
        }
        mutate(data)
        handleClose()
    }
    const handleDel = async (id) => {
        if (!confirm(ConfirmMsg)) return
        const tid = toast.loading(LoadingMsg);
        const {code, msg} = await delRoleMenu(id)
        if (resCodeShowMsg(code, msg, tid)) return
        mutate(data)
        handleClose()
    }
    return (<>
        <Card elevation={3}>
            <CardHeader title={`角色菜单`} action={<Search params={params} setParams={setParams} handleShow={handleShow} menuOptions={menuOptions}/>}/>
            <CardContent><PageTable cellNames={['id', '角色', '菜单', '操作']}>
                {data.data.list.map(item => (<TableRow key={item.id}>
                    <TableCell>{item.id}</TableCell>
                    <TableCell>{item.r_name}</TableCell>
                    <TableCell>{item.m_name}</TableCell>
                    <TableCell>
                        <Button color={'error'} onClick={() => handleDel(item.id)}>删除</Button>
                    </TableCell>
                </TableRow>))}
            </PageTable>
            </CardContent>
            <PageCardActions setParams={setParams} params={params} data={data}/>
        </Card>
        <PageDialog add={form.add} show={show} handleClose={handleClose} handleSubmit={handleSubmit}>
            <input type={'hidden'} value={form.id}/>
            <TextField label={'mid'} autoFocus margin={'dense'} size={'small'} value={form.mid || ''} onChange={e => handleChange(e, 'mid')} select SelectProps={{native: true, multiple: true}}>
                <option value={''}/>
                {menuOptions.map(item => <option key={item.id} value={item.id}>{item.name}</option>)}
            </TextField>
        </PageDialog>
    </>)
}

const Search = ({handleShow}) => {
    return (<ButtonGroup variant={'text'} className={'d-flex justify-content-between'}>
        <Button color={'success'} onClick={() => handleShow(true)} size={"small"}>添加</Button>
    </ButtonGroup>)
}