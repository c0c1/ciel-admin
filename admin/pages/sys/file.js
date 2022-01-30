import request from "../../utils/request";
import {defaultPageParam, imgPrefix} from "../../data/const";
import {useState} from "react";
import useSWR from "swr";
import {toast} from "react-toastify";
import Meta from "../../components/meta";
import {Nav} from "../../components/nav";
import {SearchLi} from "../../components/searchLi";
import {TableTd} from "../../components/tableTd";
import {TableFooter} from "../../components/table";
import {DetailsLi2} from "../../components/detailsLi";

const fields = [
    {field: 'id', editHidden: 1},
    {field: 'name'},
    {
        field: 'group', search: 1, type: 'select', items: [
            {label: 'img', value: 1},
            {label: 'icon', value: 2},
            {label: 'file', value: 3},
            {label: 'other', value: 4},
        ]
    },
    {
        field: 'status', type: 'select', items: [
            {label: 'yes', value: 1}, {label: 'no', value: 2}
        ]
    },
    {field: 'created_at', editHidden: 1},
]
const urlPrefix = "file"
const name = 'File'
const File = () => {
    const fetcher = async ({url, params}) => request({url: url, method: 'get', params}),
        update = async (data) => request({url: `/${urlPrefix}/update`, method: 'put', data}),
        del = async (id) => request({url: `/${urlPrefix}/del?id=${id}`, method: 'delete'}),
        uploadFile = async (data) => request({url: `/file/upload`, method: 'post', data})
    const [params, setParams] = useState({...defaultPageParam}), [paramsTemp, setParamsTemp] = useState(params),
        [details, setDetails] = useState({showType: "main"}),
        {data, mutate} = useSWR({url: `/${urlPrefix}/list`, params: params}, fetcher),
        {currPage, list, totalPage, totalCount} = data ? data.data : ""
    const onSearch = () => setParams(paramsTemp),
        onChange = (e, name) => setParamsTemp({...paramsTemp, [name]: e.target.value}),
        onDetailsChange = (e, name) => {
            if (name === 'files') {
                setDetails({...details, [name]: e.target.files})
                return
            }
            setDetails({...details, [name]: e.target.value})
        },
        onKeyDown = (e) => (e.key === 'Enter') ? params.page = 1 && onSearch() : "",
        onDelete = async (id) => (confirm("Are you sure?")) ? await del(id) && toast('ok') && mutate() : "",
        onSubmit = async () => {
            if (!details.group) {
                toast.error('Group must be chose!')
                return
            }

            if (details.showType == "update") {
                await update(details)
            } else {
                if (!details.files) {
                    toast.error('Files must be chose!')
                    return
                }
                let forData = new FormData()
                for (let i = 0; i < details.files.length; i++) {
                    forData.append('files', details.files[i])
                }
                forData.append('group', details.group)
                await uploadFile(forData)
            }
            mutate()
            toast('ok')
            setDetails({showType: "main"})
        }
    return <>
        <Meta/>
        <Nav/>
        <header className={'hr'}>
            <h1>{name}</h1>
            <section>
                <ul>
                    {fields.filter(item => item.search).map(item =>
                        <SearchLi key={item.field}
                                  field={item.field} type={item.type} items={item.items}
                                  onChange={onChange} paramsTemp={paramsTemp} onKeyDown={onKeyDown}
                        />
                    )}
                </ul>
                <ul>
                    <input type="submit" value={'search'} onClick={() => onSearch()}/>
                    <input type="submit" value={'add'} onClick={() => setDetails({showType: "add"})}/>
                </ul>
            </section>
        </header>
        <main className={'hr'} hidden={!(details.showType === "main")}>
            {!list ? "" : (
                <table>
                    <thead>
                    <tr>
                        {['id', 'name', 'img', 'group', 'status', 'created_at', 'Operation'].map(item =>
                            <th key={item}>{item}</th>)}
                    </tr>
                    </thead>
                    <tbody>
                    {list.map(item => <tr key={item.id}>
                        <TableTd value={item.id}/>
                        <TableTd value={item.name}/>
                        <TableTd value={`${imgPrefix}/${item.name}`} type={'img'}/>
                        <TableTd value={item.group} type={'select'} items={[
                            {label: 'img', value: 1},
                            {label: 'icon', value: 2},
                            {label: 'file', value: 3},
                            {label: 'other', value: 4},
                        ]}/>
                        <TableTd value={item.status} type={'select'} items={[{label: 'yse', value: 1}, {label: 'no', value: 2}]}/>
                        <TableTd value={item.created_at}/>
                        <td>
                            <input type={'submit'} value={'Edit'} onClick={() => setDetails({...item, showType: "update"})}/>
                            <input type={'submit'} value={'Del'} onClick={() => onDelete(item.id)}/>
                        </td>
                    </tr>)}
                    </tbody>
                    <TableFooter params={params} setParams={setParams} current={currPage} tCount={totalCount} tPage={totalPage}/>
                </table>
            )}
        </main>
        <section className={'details'} hidden={!(details.showType === "add" || details.showType === "update")}>
            <h2>Step1 Set info</h2>
            <section>
                <ul>
                    {fields.filter(item => !item.editHidden).map(item =>
                        <DetailsLi2 items={item.items} type={item.type} key={item.field} field={item.field} label={item.title} height={item.detailsHeight} placeholder={item.detailsDesc} details={details} onDetailsChange={onDetailsChange}/>
                    )}
                    <li hidden={details.showType === 'update'}>文件<input type={"file"} multiple={true} onChange={e => onDetailsChange(e, 'files')}/></li>
                </ul>
            </section>
            <h2>Step2 Go!</h2>
            <p>Be patient! It may take a little while to finish your request...</p>
            <input type={'button'} value={'Go Back'} onClick={() => setDetails({...details, showType: "main"})}/>
            <input type={'button'} value={'Submit'} onClick={onSubmit}/>
        </section>
    </>
}
export default File