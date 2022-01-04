import Button from "@mui/material/Button";
import {CardActions, Chip, Dialog, DialogActions, DialogContent, DialogTitle, FormControl, InputLabel, MenuItem, Pagination, Select, Table, TableBody, TableCell, TableHead, TableRow} from "@mui/material";

export const PageSearchBtn = ({handleSearch, handleShow}) => {
    return (<>
        <Button color={'primary'} onClick={() => handleSearch()} size={"small"}>搜索</Button>
        <Button color={'success'} onClick={() => handleShow(true)} size={"small"}>添加</Button>
    </>)
}
export const PageOptionStatus = () => {
    return (<>
        <option value={''}/>
        <option value={1}>&nbsp; &nbsp; 是</option>
        <option value={2}>&nbsp; &nbsp; 否</option>
    </>)
}

export const PageTable = ({children, cellNames}) => {
    return (<Table sx={{minWidth: 650}} size={"small"}>
        <TableHead><TableRow>{cellNames.map(item => <TableCell key={item}>{item}</TableCell>)}</TableRow></TableHead>
        <TableBody>{children}</TableBody>
    </Table>)
}
export const PageDataStatus = ({status}) => status === 1 ? <Chip label={'是'} color={'success'} size={"small"}/> : <Chip label={'否'} color={'error'} size={"small"}/>

export const PageDataBtn = ({handleShow, handleDel}) => (<>
    <Button color={'warning'} onClick={handleShow}>编辑</Button>
    <Button color={'error'} onClick={handleDel}>删除</Button>
</>)

export const PageCardActions = ({params, setParams, data}) => {
    const {totalCount, totalPage} = data.data
    return (<CardActions className={'d-flex justify-content-end'}>
        <FormControl size={"small"}>
            <InputLabel id="demo-simple-select-label">分页大小</InputLabel>
            <Select
                label="分页大小" labelId="demo-simple-select-label" id="demo-simple-select" sx={{width: '120px'}}
                value={params.pageSize || 10}
                onChange={(e) => setParams({...params, page: 1, pageSize: e.target.value})}>
                <MenuItem value={2}>2</MenuItem>
                <MenuItem value={5}>5</MenuItem>
                <MenuItem value={10}>10</MenuItem>
                <MenuItem value={20}>20</MenuItem>
                <MenuItem value={30}>30</MenuItem>
            </Select>
        </FormControl>
        <Pagination
            count={data ? totalPage : 0} color={'primary'} variant={'outlined'}
            page={params.page}
            onChange={(e, page) => setParams({...params, page: page})}/>
        <Chip size='small' label={` 总条数:${data ? totalCount : 0}`} variant="outlined"/>
    </CardActions>)
}

export const PageDialog = ({add, children, show, handleClose, handleSubmit, name}) => {
    return (<Dialog open={show} onClose={handleClose}>
        <DialogTitle>{name ? name : add ? '添加' : '修改'}</DialogTitle>
        <DialogContent>
            {children}
        </DialogContent>
        <DialogActions>
            <Button onClick={handleClose} color={'secondary'} size={'small'}>取消</Button>
            <Button onClick={handleSubmit} color={'primary'} size={'small'}>提交</Button>
        </DialogActions>
    </Dialog>)
}

export const PageStatusOption = () => {
    return <>
        <option value={''}/>
        <option value={1}>是</option>
        <option value={2}>否</option>
    </>
}
