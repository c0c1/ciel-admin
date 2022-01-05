
import useSWR from "swr";
import {useState} from "react";
import {toast} from "react-toastify";
import {ButtonGroup, Card, CardContent, CardHeader, TableCell, TableRow, TextField} from "@mui/material";
import {PageCardActions, PageDataBtn, PageDataStatus, PageDialog, PageOptionStatus, PageSearchBtn, PageStatusOption, PageTable} from "../../components/page";
import {ConfirmMsg, DefaultPageParam, LoadingMsg} from "../../consts/const";
import {resBaseCode, resCodeShowMsg} from "../../utils/util";
import {Server} from "../../config";
import {addLoginLog, delLoginLog, getLoginLogById, loginLogFetcher, updateLoginLog} from "../../data/user";

export default function LoginLog() {
    const [params, setParams] = useState({...DefaultPageParam, uid: '', ip: ''})
    const [form, setForm] = useState({add: true})
    const [show, setShow] = useState(false);
    const {data, mutate} = useSWR({url: Server+"/loginLog/list", params: params}, loginLogFetcher)
    if (resBaseCode(data)) return <div/>
    const handleClose = () => setShow(false);
    const handleChange = (e, name) => setForm({...form, [name]: e.target.value})
    const handleShow = async (add, id) => {
        setForm({add: add})
        if (!add) {
            const {data} = await getLoginLogById(id)

            setForm({...data.normal_data})
        }
        setShow(true)
    }
    const handleSubmit = async () => {
        const tid = toast.loading(LoadingMsg);
        if (form.add) {
            const {code, message} = await addLoginLog(form)
            if (resCodeShowMsg(code, message, tid)) return
        } else {
            const {code, message} = await updateLoginLog(form)
            if (resCodeShowMsg(code, message, tid)) return
        }
        mutate(data)
        handleClose()
    }
    const handleDel = async (id) => {
        if (!confirm(ConfirmMsg)) return
        const tid = toast.loading(LoadingMsg);
        const {code, msg} = await delLoginLog(id)
        if (resCodeShowMsg(code, msg, tid)) return
        mutate(data)
        handleClose()
    }
    return (<>
        <Card elevation={3}>
            <CardHeader title={'登录日志'} action={<Search params={params} setParams={setParams} handleShow={handleShow}/>}/>
            <CardContent><PageTable cellNames={['id', '用户ID', 'IP', '备注','状态','创建时间','修改时间','操作']}>
                {data.data.list.map(item => (<TableRow key={item.id}>
                    <TableCell>{item.id}</TableCell>
                    <TableCell>{item.uid}</TableCell>
                    <TableCell>{item.ip}</TableCell>
                    <TableCell>{item.desc}</TableCell>
                    <TableCell><PageDataStatus status={item.status}/></TableCell>
                    <TableCell>{item.created_at}</TableCell>
                    <TableCell>{item.updated_at}</TableCell>
                    <TableCell><ButtonGroup variant={'text'}><PageDataBtn handleShow={() => handleShow(false, item.id)} handleDel={() => handleDel(item.id)}/></ButtonGroup></TableCell>
                </TableRow>))}
            </PageTable>
            </CardContent>
            <PageCardActions setParams={setParams} params={params} data={data}/>
        </Card>
        <PageDialog add={form.add} show={show} handleClose={handleClose} handleSubmit={handleSubmit}>
            <input type={'hidden'} value={form.id}/>
            <TextField label={'用户ID'} autoFocus margin={'dense'} size={'small'} value={form.uid || ''} onChange={e => handleChange(e, 'uid')}/>
            <TextField label={'IP'} autoFocus margin={'dense'} size={'small'} value={form.ip || ''} onChange={e => handleChange(e, 'ip')}/>
            <TextField label={'备注'} autoFocus margin={'dense'} size={'small'} value={form.desc || ''} onChange={e => handleChange(e, 'desc')}/>
            <TextField label={'状态'} autoFocus margin={'dense'} size={'small'} value={form.status} onChange={e => handleChange(e, 'status')} select SelectProps={{native: true}}><PageStatusOption/></TextField>
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
    return (<ButtonGroup className={'d-flex justify-content-between'} variant={'text'}>
        <TextField label={'用户ID'} size={"small"} className={'me-2'} value={form.uid} onKeyUp={handleKeyUp} onChange={e => handleChange(e, 'uid')}/>
        <TextField label="IP" size={"small"} className={'me-2'} value={form.ip} onKeyUp={handleKeyUp} onChange={e => handleChange(e, 'ip')}/>
        <TextField label={'状态'} size={"small"} className={'me-2'} value={form.status} onChange={e => handleChange(e, 'status')} select SelectProps={{native: true}}> <PageOptionStatus/></TextField>
        <PageSearchBtn handleSearch={handleSearch} handleShow={handleShow}/>
    </ButtonGroup>)
}
