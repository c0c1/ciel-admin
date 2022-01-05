import useSWR from "swr";
import {useState} from "react";
import {toast} from "react-toastify";
import {ButtonGroup, Card, CardContent, CardHeader, TableCell, TableRow, TextField} from "@mui/material";
import {PageCardActions, PageDataBtn, PageDataStatus, PageDialog, PageOptionStatus, PageSearchBtn, PageStatusOption, PageTable} from "../../components/page";
import {ConfirmMsg, DefaultPageParam, LoadingMsg} from "../../consts/const";
import {isEmpty, resBaseCode, resCodeShowMsg} from "../../utils/util";
import {ImgPrefix, Server} from "../../config";
import {delFile, fileFetcher, getFileById, updateFile, uploadFile} from "../../data/sys";

const groups = [{label: '文件', value: 'file'}, {label: '头像', value: 'icon'}, {label: '视频', value: 'video'}, {label: '音频', value: 'audio'}, {label: '其他', value: 'other'},]
export default function File() {
    const [params, setParams] = useState({...DefaultPageParam, name: '', group: ''})
    const [form, setForm] = useState({add: true})
    const [show, setShow] = useState(false);
    const {data, mutate} = useSWR({url: `${Server}/file/list`, params: params}, fileFetcher)
    if (resBaseCode(data)) return <div/>
    const handleClose = () => setShow(false);
    const handleChange = (e, name) => {
        if (name === 'files') {
            setForm({...form, [name]: e.target.files})
            return
        }
        setForm({...form, [name]: e.target.value})
    }
    const handleShow = async (add, id) => {
        setForm({add: add})
        if (!add) {
            const {data} = await getFileById(id)

            setForm({...data.normal_data})
        }
        setShow(true)
    }
    const handleSubmit = async () => {
        if (isEmpty(form.group)) return toast.warning('分组不能为空')
        if (isEmpty(form.files)) return toast.warning('文件不能为空')
        const tid = toast.loading(LoadingMsg);
        if (form.add) {
            let formData = new FormData();
            for (let i = 0; i < form.files.length; i++) {
                formData.append("files", form.files[i])
            }
            formData.append('group', form.group)
            const {code, message} = await uploadFile(formData)
            if (resCodeShowMsg(code, message, tid)) return
            mutate(data)
        } else {
            const {code, message} = await updateFile(form)
            if (resCodeShowMsg(code, message, tid)) return
        }
        mutate(data)
        handleClose()
    }
    const handleDel = async (id) => {
        if (!confirm(ConfirmMsg)) return
        const tid = toast.loading(LoadingMsg);
        const {code, msg} = await delFile(id)
        if (resCodeShowMsg(code, msg, tid)) return
        mutate(data)
        handleClose()
    }
    const getGroup = (group) => {
        switch (group) {
            case "file":
                return "文件"
            case "icon":
                return "头像"
            case "video":
                return '视频'
            case "audio":
                return "音频"
            default:
                return "其他"
        }
    }
    return (<>
        <Card elevation={3}>
            <CardHeader title={'文件'} action={<Search params={params} setParams={setParams} handleShow={handleShow}/>}/>
            <CardContent><PageTable cellNames={['id', '图片', '文件名', '分组', '状态', '创建时间', '修改时间', '操作']}>
                {data.data.list.map(item => (<TableRow key={item.id}>
                    <TableCell>{item.id}</TableCell>
                    <TableCell><a href={ImgPrefix + '/' + item.name} target="_blank" rel="noreferrer noopener"><img className={'img-fluid'} src={ImgPrefix + '/' + item.name} alt="can't find" style={{width: '25px'}}/></a></TableCell>
                    <TableCell>{item.name}</TableCell>
                    <TableCell>{getGroup(item.group)}</TableCell>
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
            <TextField disabled label={'文件名'} autoFocus margin={'dense'} size={'small'} value={form.name || ''} onChange={e => handleChange(e, 'name')} fullWidth/>
            <TextField label={'分组'} autoFocus margin={'dense'} size={'small'} value={form.group || ''} onChange={e => handleChange(e, 'group')}
                       select SelectProps={{native: true}}>
                <option value={''}/>
                {groups.map(item => <option value={item.value} key={item.value}>{item.label}</option>)}
            </TextField>
            <TextField label={'状态'} autoFocus margin={'dense'} size={'small'} value={form.status} onChange={e => handleChange(e, 'status')} select SelectProps={{native: true}}><PageStatusOption/></TextField> <br/>
            <input hidden={!form.add} type="file" className={'mt-2'} multiple onChange={e => handleChange(e, 'files')}/>
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
        <TextField label={'文件名'} size={"small"} className={'me-2'} value={form.name} onKeyUp={handleKeyUp} onChange={e => handleChange(e, 'name')}/>
        <TextField label={'分组'} size={"small"} className={'me-2'} value={form.group || ''} onChange={e => handleChange(e, 'group')}
                   select SelectProps={{native: true}}>
            <option value={''}/>
            {groups.map(item => <option value={item.value} key={item.value}>{item.label}</option>)}
        </TextField>
        <TextField label={'状态'} size={"small"} className={'me-2'} value={form.status} onChange={e => handleChange(e, 'status')} select SelectProps={{native: true}}> <PageOptionStatus/></TextField>
        <PageSearchBtn handleSearch={handleSearch} handleShow={handleShow}/>
    </ButtonGroup>)
}