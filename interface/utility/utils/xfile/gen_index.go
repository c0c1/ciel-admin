package xfile

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

var IndexTemplate = `
import useSWR from "swr";
import {useState} from "react";
import {toast} from "react-toastify";
import {ButtonGroup, Card, CardContent, CardHeader, TableCell, TableRow, TextField} from "@mui/material";
import {PageCardActions, PageDataBtn, PageDataStatus, PageDialog, PageOptionStatus, PageSearchBtn, PageStatusOption, PageTable} from "../../components/page";
import {ConfirmMsg, DefaultPageParam, LoadingMsg} from "../../consts/const";
import {resBaseCode, resCodeShowMsg} from "../../utils/util";
import {Server} from "../../config";
import {add$Name$, del$Name$, get$Name$ById, $name$Fetcher, update$Name$} from "../../data/$group$";

export default function $Name$() {
    const [params, setParams] = useState({...DefaultPageParam, $p1$: '', $p2$: ''})
    const [form, setForm] = useState({add: true})
    const [show, setShow] = useState(false);
    const {data, mutate} = useSWR({url: Server+"/$name$/list", params: params}, $name$Fetcher)
    if (resBaseCode(data)) return <div/>
    const handleClose = () => setShow(false);
    const handleChange = (e, name) => setForm({...form, [name]: e.target.value})
    const handleShow = async (add, id) => {
        setForm({add: add})
        if (!add) {
            const {data} = await get$Name$ById(id)
            setForm({...data.normal_data})
        }
        setShow(true)
    }
    const handleSubmit = async () => {
        const tid = toast.loading(LoadingMsg);
        if (form.add) {
            const {code, message} = await add$Name$(form)
            if (resCodeShowMsg(code, message, tid)) return
        } else {
            const {code, message} = await update$Name$(form)
            if (resCodeShowMsg(code, message, tid)) return
        }
        mutate(data)
        handleClose()
    }
    const handleDel = async (id) => {
        if (!confirm(ConfirmMsg)) return
        const tid = toast.loading(LoadingMsg);
        const {code, msg} = await del$Name$(id)
        if (resCodeShowMsg(code, msg, tid)) return
        mutate(data)
        handleClose()
    }
    return (<>
        <Card elevation={3}>
            <CardHeader title={'$NameZh$'} action={<Search params={params} setParams={setParams} handleShow={handleShow}/>}/>
            <CardContent><PageTable cellNames={['id', '$p1name$', '$p2name$', '$p3name$','状态','创建时间','修改时间','操作']}>
                {data.data.list.map(item => (<TableRow key={item.id}>
                    <TableCell>{item.id}</TableCell>
                    <TableCell>{item.$p1$}</TableCell>
                    <TableCell>{item.$p2$}</TableCell>
                    <TableCell>{item.$p3$}</TableCell>
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
            <TextField label={'$p1name$'} autoFocus margin={'dense'} size={'small'} value={form.$p1$ || ''} onChange={e => handleChange(e, '$p1$')}/>
            <TextField label={'$p2name$'} autoFocus margin={'dense'} size={'small'} value={form.$p2$ || ''} onChange={e => handleChange(e, '$p2$')}/>
            <TextField label={'$p3name$'} autoFocus margin={'dense'} size={'small'} value={form.$p3$ || ''} onChange={e => handleChange(e, '$p3$')}/>
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
        <TextField label={'$p1name$'} size={"small"} className={'me-2'} value={form.$p1$} onKeyUp={handleKeyUp} onChange={e => handleChange(e, '$p1$')}/>
        <TextField label="$p2name$" size={"small"} className={'me-2'} value={form.$p2$} onKeyUp={handleKeyUp} onChange={e => handleChange(e, '$p2$')}/>
        <TextField label={'状态'} size={"small"} className={'me-2'} value={form.status} onChange={e => handleChange(e, 'status')} select SelectProps={{native: true}}> <PageOptionStatus/></TextField>
        <PageSearchBtn handleSearch={handleSearch} handleShow={handleShow}/>
    </ButtonGroup>)
}
`

func genIndex(c TemplateConfig) {
	file, err := gfile.Create(c.RootPath + c.PathIndex + "/" + gstr.CaseCamelLower(c.EntityName) + ".js")
	template := IndexTemplate
	template = strings.ReplaceAll(template, "$Name$", c.EntityName)
	template = strings.ReplaceAll(template, "$name$", gstr.CaseCamelLower(c.EntityName))
	template = strings.ReplaceAll(template, "$group$", c.group)
	template = strings.ReplaceAll(template, "$p1$", c.P1)
	template = strings.ReplaceAll(template, "$p2$", c.P2)
	template = strings.ReplaceAll(template, "$p3$", c.P3)
	template = strings.ReplaceAll(template, "$p1name$", c.P1Name)
	template = strings.ReplaceAll(template, "$p2name$", c.P2Name)
	template = strings.ReplaceAll(template, "$p3name$", c.P3Name)
	template = strings.ReplaceAll(template, "$NameZh$", c.NameZh)
	if _, err = file.WriteString(template); err != nil {
		panic(err)
	}
	if err = file.Close(); err != nil {
		panic(err)
	}
}
