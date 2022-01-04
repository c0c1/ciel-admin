import useSWR from "swr";
import {useEffect, useState} from "react";
import {toast} from "react-toastify";
import {ButtonGroup, Card, CardContent, CardHeader, TableCell, TableRow, TextField} from "@mui/material";
import {PageCardActions, PageDialog, PageTable} from "../../components/page";
import {ConfirmMsg, DefaultPageParam, LoadingMsg} from "../../consts/const";
import {getQuery, resBaseCode, resCodeShowMsg} from "../../utils/util";
import {Server} from "../../config";
import {addRoleApi, delRoleApi, getRoleApiById, noApis, roleApiFetcher, updateRoleApi} from "../../data/sys";
import Button from "@mui/material/Button";

export default function RoleApi() {
    const [params, setParams] = useState({...DefaultPageParam, rid: '', aid: ''})
    const [form, setForm] = useState({add: true})
    const [show, setShow] = useState(false);
    const [apiOptions, setApiOptions] = useState([]);
    useEffect(() => {
        setParams({...params, rid: getQuery('rid')})
    }, [setParams])
    const {data, mutate} = useSWR({url: `${Server}/roleApi/list`, params: params}, roleApiFetcher)
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
        const res = await noApis(rid)
        const {normal_data} = res.data
        setApiOptions(normal_data ? normal_data : [])
        setForm({add: add, rid: rid, aid: []})
        if (!add) {
            const {data} = await getRoleApiById(id)
            if (resBaseCode(data)) return
            setForm({...data.normal_data})
        }
        setShow(true)
    }
    const handleSubmit = async () => {
        const tid = toast.loading(LoadingMsg);
        if (form.add) {
            const {code, message} = await addRoleApi(form)
            if (resCodeShowMsg(code, message, tid)) return
        } else {
            const {code, message} = await updateRoleApi(form)
            if (resCodeShowMsg(code, message, tid)) return
        }
        mutate(data)
        handleClose()
    }
    const handleDel = async (id) => {
        if (!confirm(ConfirmMsg)) return
        const tid = toast.loading(LoadingMsg);
        const {code, msg} = await delRoleApi(id)
        if (resCodeShowMsg(code, msg, tid)) return
        mutate(data)
        handleClose()
    }
    return (<>
        <Card elevation={3}>
            <CardHeader title={'角色禁用API'} action={<Search params={params} setParams={setParams} handleShow={handleShow}/>}/>
            <CardContent><PageTable cellNames={['id', '角色名称', '分组', '方法', '路径', '描述', '操作']}>
                {data.data.list.map(item => (<TableRow key={item.id}>
                    <TableCell>{item.id}</TableCell>
                    <TableCell>{item.role_name}</TableCell>
                    <TableCell>{item.group}</TableCell>
                    <TableCell>{item.method}</TableCell>
                    <TableCell>{item.url}</TableCell>
                    <TableCell>{item.desc}</TableCell>
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
            <TextField label={'aid'} autoFocus margin={'dense'} size={'small'} value={form.aid || ''} onChange={e => handleChange(e, 'aid')} select SelectProps={{native: true, multiple: true}}>
                <option value={''}/>
                {apiOptions.map((item) => <option value={item.id} key={item.id + item.url + item.method}>{item.group + ' ' + item.method + ' ' + item.url + ' ' + item.desc}</option>)}
            </TextField>
        </PageDialog>
    </>)
}

const Search = ({params, setParams, handleShow}) => {
    return (<ButtonGroup variant={'text'} className={'d-flex justify-content-between'}>
        <Button color={'success'} onClick={() => handleShow(true)} size={"small"}>添加</Button>
    </ButtonGroup>)
}