
import useSWR from "swr";
import {useState} from "react";
import {toast} from "react-toastify";
import {ButtonGroup, Card, CardContent, CardHeader, TableCell, TableRow, TextField} from "@mui/material";
import {PageCardActions, PageDataBtn, PageDataStatus, PageDialog, PageOptionStatus, PageSearchBtn, PageStatusOption, PageTable} from "../../components/page";
import {ConfirmMsg, DefaultPageParam, LoadingMsg} from "../../consts/const";
import {resBaseCode, resCodeShowMsg} from "../../utils/util";
import {Server} from "../../config";
import {addUserDetails, delUserDetails, getUserDetailsById, userDetailsFetcher, updateUserDetails} from "../../data/user";

export default function UserDetails() {
    const [params, setParams] = useState({...DefaultPageParam, uid: '', real_name: ''})
    const [form, setForm] = useState({add: true})
    const [show, setShow] = useState(false);
    const {data, mutate} = useSWR({url: Server+"/userDetails/list", params: params}, userDetailsFetcher)
    if (resBaseCode(data)) return <div/>
    const handleClose = () => setShow(false);
    const handleChange = (e, name) => setForm({...form, [name]: e.target.value})
    const handleShow = async (add, id) => {
        setForm({add: add})
        if (!add) {
            const {data} = await getUserDetailsById(id)

            setForm({...data.normal_data})
        }
        setShow(true)
    }
    const handleSubmit = async () => {
        const tid = toast.loading(LoadingMsg);
        if (form.add) {
            const {code, message} = await addUserDetails(form)
            if (resCodeShowMsg(code, message, tid)) return
        } else {
            const {code, message} = await updateUserDetails(form)
            if (resCodeShowMsg(code, message, tid)) return
        }
        mutate(data)
        handleClose()
    }
    const handleDel = async (id) => {
        if (!confirm(ConfirmMsg)) return
        const tid = toast.loading(LoadingMsg);
        const {code, msg} = await delUserDetails(id)
        if (resCodeShowMsg(code, msg, tid)) return
        mutate(data)
        handleClose()
    }
    return (<>
        <Card elevation={3}>
            <CardHeader title={'用户详情'} action={<Search params={params} setParams={setParams} handleShow={handleShow}/>}/>
            <CardContent><PageTable cellNames={['id', '用户ID', '姓名', '备注','状态','创建时间','修改时间','操作']}>
                {data.data.list.map(item => (<TableRow key={item.id}>
                    <TableCell>{item.id}</TableCell>
                    <TableCell>{item.uid}</TableCell>
                    <TableCell>{item.real_name}</TableCell>
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
            <TextField label={'姓名'} autoFocus margin={'dense'} size={'small'} value={form.real_name || ''} onChange={e => handleChange(e, 'real_name')}/>
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
        <TextField label="姓名" size={"small"} className={'me-2'} value={form.real_name} onKeyUp={handleKeyUp} onChange={e => handleChange(e, 'real_name')}/>
        <TextField label={'状态'} size={"small"} className={'me-2'} value={form.status} onChange={e => handleChange(e, 'status')} select SelectProps={{native: true}}> <PageOptionStatus/></TextField>
        <PageSearchBtn handleSearch={handleSearch} handleShow={handleShow}/>
    </ButtonGroup>)
}
