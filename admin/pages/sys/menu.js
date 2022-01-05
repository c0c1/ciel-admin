import useSWR from "swr";
import {useState} from "react";
import {toast} from "react-toastify";
import {ButtonGroup, Card, CardContent, CardHeader, TableCell, TableRow, TextField} from "@mui/material";
import {PageCardActions, PageDataBtn, PageDataStatus, PageDialog, PageOptionStatus, PageSearchBtn, PageStatusOption, PageTable} from "../../components/page";
import {ConfirmMsg, DefaultPageParam, LoadingMsg} from "../../consts/const";
import {resBaseCode, resCodeShowMsg} from "../../utils/util";
import {Server} from "../../config";
import {addMenu, delMenu, getMenuById, menuFetcher, updateMenu} from "../../data/sys";

export default function Menu() {
    const [params, setParams] = useState({...DefaultPageParam, pid: '', name: ''})
    const [form, setForm] = useState({add: true})
    const [show, setShow] = useState(false);
    const {data, mutate} = useSWR({url: `${Server}/menu/list`, params: params}, menuFetcher)
    if (resBaseCode(data)) return <div/>
    const handleClose = () => setShow(false);
    const handleChange = (e, name) => setForm({...form, [name]: e.target.value})
    const handleShow = async (add, id) => {
        setForm({add: add})
        if (!add) {
            const {data} = await getMenuById(id)
            setForm({...data.normal_data})
        }
        setShow(true)
    }
    const handleSubmit = async () => {
        const tid = toast.loading(LoadingMsg);
        if (form.add) {
            const {code, message} = await addMenu(form)
            if (resCodeShowMsg(code, message, tid)) return
        } else {
            const {code, message} = await updateMenu(form)
            if (resCodeShowMsg(code, message, tid)) return
        }
        mutate(data)
        handleClose()
    }
    const handleDel = async (id) => {
        if (!confirm(ConfirmMsg)) return
        const tid = toast.loading(LoadingMsg);
        const {code, msg} = await delMenu(id)
        if (resCodeShowMsg(code, msg, tid)) return
        mutate(data)
        handleClose()
    }
    return (<>
        <Card elevation={3}>
            <CardHeader title={'菜单'} action={<Search params={params} setParams={setParams} handleShow={handleShow}/>}/>
            <CardContent><PageTable cellNames={['id', '父ID', '名称', '排序', '路径', '类型', '状态', '操作']}>
                {data.data.list.map(item => (<TableRow key={item.id}>
                    <TableCell>{item.id}</TableCell>
                    <TableCell>{item.pid}</TableCell>
                    <TableCell>{item.name}</TableCell>
                    <TableCell>{item.sort}</TableCell>
                    <TableCell>{item.path}</TableCell>
                    <TableCell>{item.type === 0 ? '菜单' : item.type === 1 ? '分组' : "分隔线"}</TableCell>
                    <TableCell><PageDataStatus status={item.status}/></TableCell>
                    <TableCell><ButtonGroup variant={'text'}><PageDataBtn handleShow={() => handleShow(false, item.id)} handleDel={() => handleDel(item.id)}/></ButtonGroup></TableCell>
                </TableRow>))}
            </PageTable>
            </CardContent>
            <PageCardActions setParams={setParams} params={params} data={data}/>
        </Card>
        <PageDialog add={form.add} show={show} handleClose={handleClose} handleSubmit={handleSubmit}>
            <input type={'hidden'} value={form.id}/>
            <TextField label={'父ID'} autoFocus margin={'dense'} size={'small'} value={form.pid || ''} onChange={e => handleChange(e, 'pid')}/>
            <TextField label={'名称'} autoFocus margin={'dense'} size={'small'} value={form.name || ''} onChange={e => handleChange(e, 'name')}/>
            <TextField label={'路径'} autoFocus margin={'dense'} size={'small'} value={form.path || ''} onChange={e => handleChange(e, 'path')}/>
            <TextField label={'排序'} autoFocus margin={'dense'} size={'small'} value={form.sort || ''} onChange={e => handleChange(e, 'sort')}/>
            <TextField label={'图标'} autoFocus margin={'dense'} size={'small'} value={form.icon || ''} onChange={e => handleChange(e, 'icon')}/>
            <TextField label={'类型'} autoFocus margin={'dense'} size={'small'} value={form.type} onChange={e => handleChange(e, 'type')} select SelectProps={{native: true}}>
                <option value={''}/>
                <option value={0}>菜单</option>
                <option value={1}>分组</option>
                <option value={2}>分隔线</option>
            </TextField>
            <TextField label={'status'} autoFocus margin={'dense'} size={'small'} value={form.status} onChange={e => handleChange(e, 'status')} select SelectProps={{native: true}}><PageStatusOption/></TextField>
        </PageDialog>
    </>)
}

const Search = ({params, setParams, handleShow}) => {
    const [form, setForm] = useState(params);
    const handleChange = (e, name) => setForm({...form, [name]: e.target.value})
    const handleSubmit = () => setParams(form)
    const handleSearch = () => {
        form.page = 1
        setParams(form)
    }
    const handleKeyUp = (e) => {
        form.page = 1
        if (e.keyCode === 13) handleSubmit()
    }
    return (<ButtonGroup variant={'text'} className={'d-flex justify-content-between'}>
        <TextField label={'父ID'} size={"small"} className={'me-2'} value={form.pid} onKeyUp={handleKeyUp} onChange={e => handleChange(e, 'pid')}/>
        <TextField label="名称" size={"small"} className={'me-2'} value={form.name} onKeyUp={handleKeyUp} onChange={e => handleChange(e, 'name')}/>
        <TextField label={'状态'} size={"small"} className={'me-2'} value={form.status} onChange={e => handleChange(e, 'status')} select SelectProps={{native: true}}> <PageOptionStatus/></TextField>
        <PageSearchBtn handleSearch={handleSearch} handleShow={handleShow}/>
    </ButtonGroup>)
}