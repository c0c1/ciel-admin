import useSWR from "swr";
import {useState} from "react";
import {toast} from "react-toastify";
import {ButtonGroup, Card, CardContent, CardHeader, Checkbox, FormControlLabel, TableCell, TableRow, TextField} from "@mui/material";
import {PageCardActions, PageDataBtn, PageDataStatus, PageDialog, PageOptionStatus, PageSearchBtn, PageStatusOption, PageTable} from "../../components/page";
import {ConfirmMsg, DefaultPageParam, LoadingMsg} from "../../consts/const";
import {isEmpty, resBaseCode, resCodeShowMsg} from "../../utils/util";
import {Server} from "../../config";
import {addAdmin, adminFetcher, delAdmin, getAdminById, updateAdmin} from "../../data/sys";

export default function Admin() {
    const [params, setParams] = useState({...DefaultPageParam, rid: '', uname: ''}), [form, setForm] = useState({add: true}), [show, setShow] = useState(false), {data, mutate} = useSWR({url: `${Server}/admin/list`, params: params}, adminFetcher)
    if (resBaseCode(data)) return <div/>
    const {list, other} = data.data, handleClose = () => setShow(false), handleChange = (e, name) => {
        if (name === 'ex1') {
            setForm({...form, [name]: !form.ex1})
        } else {
            setForm({...form, [name]: e.target.value})
        }
    }, handleShow = async (add, id) => {
        setForm({add: add})
        if (!add) {
            const {data} = await getAdminById(id)
            if (resBaseCode(data)) return
            setForm({...data.normal_data, ex1: false})
        }
        setShow(true)
    }, handleSubmit = async () => {
        if (isEmpty(form.rid)) {
            toast.warning(`角色不能为空`)
            return
        }
        const tid = toast.loading(LoadingMsg);
        if (form.add) {
            const {code, message} = await addAdmin(form)
            if (resCodeShowMsg(code, message, tid)) return
        } else {
            const {code, message} = await updateAdmin(form)
            if (resCodeShowMsg(code, message, tid)) return
        }
        mutate(data)
        handleClose()
    }, handleDel = async (id) => {
        if (!confirm(ConfirmMsg)) return
        const tid = toast.loading(LoadingMsg);
        const {code, msg} = await delAdmin(id)
        if (resCodeShowMsg(code, msg, tid)) return
        mutate(data)
        handleClose()
    }
    return (<>
        <Card elevation={3}>
            <CardHeader title={'管理员'} action={<Search params={params} setParams={setParams} handleShow={handleShow} other={other}/>}/>
            <CardContent><PageTable cellNames={['id', '角色', '用户名', '备注', '状态', '操作']}>
                {list.map(item => (<TableRow key={item.id}>
                    <TableCell>{item.id}</TableCell>
                    <TableCell>{item.role_name}</TableCell>
                    <TableCell>{item.uname}</TableCell>
                    <TableCell>{item.desc}</TableCell>
                    <TableCell><PageDataStatus status={item.status}/></TableCell>
                    <TableCell>
                        <ButtonGroup variant={'text'}>
                            <PageDataBtn handleShow={() => handleShow(false, item.id)} handleDel={() => handleDel(item.id)}/>
                        </ButtonGroup>
                    </TableCell>
                </TableRow>))}
            </PageTable>
            </CardContent>
            <PageCardActions setParams={setParams} params={params} data={data}/>
        </Card>
        <PageDialog add={form.add} show={show} handleClose={handleClose} handleSubmit={handleSubmit}>
            <input type={'hidden'} value={form.id}/>
            <TextField label={'用户名'} autoFocus margin={'dense'} size={'small'} value={form.uname || ''} onChange={e => handleChange(e, 'uname')} fullWidth/>
            <FormControlLabel hidden={form.add} checked={form.ex1} control={<Checkbox onChange={e => handleChange(e, 'ex1')}/>} label="修改用户名"/>
            <TextField label={'密码'} autoFocus margin={'dense'} size={'small'} value={form.pwd || ''} onChange={e => handleChange(e, 'pwd')} fullWidth/>
            <TextField label={'状态'} autoFocus margin={'dense'} size={'small'} value={form.status} onChange={e => handleChange(e, 'status')} select SelectProps={{native: true}}><PageStatusOption/></TextField>
            <TextField label={'角色'} autoFocus margin={'dense'} size={'small'} value={form.rid || ''} onChange={e => handleChange(e, 'rid')} select SelectProps={{native: true}}>
                <option/>
                {other.map(item => <option key={item.id} value={item.id}>{item.name}</option>)}
            </TextField>
            <TextField label={'备注'} autoFocus margin={'dense'} size={'small'} value={form.desc || ''} onChange={e => handleChange(e, 'desc')}/>
        </PageDialog>
    </>)
}

const Search = ({params, setParams, handleShow, other}) => {
    const [form, setForm] = useState(params), handleChange = (e, name) => setForm({...form, [name]: e.target.value}), handleSubmit = () => setParams(form), handleSearch = () => {
        form.page = 1
        setParams(form)
    }, handleKeyUp = (e) => {
        form.page = 1
        if (e.keyCode === 13) handleSubmit()
    };
    return (<ButtonGroup variant={'text'} className={'d-flex justify-content-between'}>
        <TextField label={'角色'} size={"small"} className={'me-2'} value={form.rid} onKeyUp={handleKeyUp} onChange={e => handleChange(e, 'rid')}
                   select SelectProps={{native: true}}
        >
            <option value={''}/>
            {other.map(item => <option key={item.id} value={item.id}>{item.name}</option>)}
        </TextField>
        <TextField label="用户名" size={"small"} className={'me-2'} value={form.uname} onKeyUp={handleKeyUp} onChange={e => handleChange(e, 'uname')}/>
        <TextField label={'状态'} size={"small"} className={'me-2'} value={form.status} onChange={e => handleChange(e, 'status')} select SelectProps={{native: true}}> <PageOptionStatus/></TextField>
        <PageSearchBtn handleSearch={handleSearch} handleShow={handleShow}/>
    </ButtonGroup>)
}