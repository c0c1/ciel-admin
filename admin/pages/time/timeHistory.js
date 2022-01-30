import request from "../../utils/request";
import {useState} from "react";
import {defaultPageParam} from "../../data/const";
import useSWR from "swr";
import {toast} from "react-toastify";
import Meta from "../../components/meta";
import {Nav} from "../../components/nav";
import {TableFooter} from "../../components/table";
import {SearchLi} from "../../components/searchLi";
import {DetailsLi} from "../../components/detailsLi";

const fields = [
    {field: 'id', editHidden: 1},
    {field: 'century', title: '世纪', search: 1},
    {field: 'year', title: '年份', search: 1},
    {field: 'month', title: '月份'},
    {field: 'day', hidden: 1},
    {field: 'type', title: '', search: 1,},
    {field: 'type_details', title: '', search: 1},
    {field: 'name', title: '事件名称', search: 1,},
    {field: 'summary', title: '描述'},
    {field: 'content', title: '内容'},
    {field: 'tag', title: '标签', search: 1},
    {field: 'sort', hidden: 1},
    {field: 'status', hidden: 1},
    {field: 'created_at', hidden: 1, editHidden: 1},
]
let urlPrefix = 'timeHistory'
let name = 'History'


const TimeHistory = () => {
    const fetcher = async ({url, params}) => request({url: url, method: 'get', params}),
        update = async (data) => request({url: `/${urlPrefix}/update`, method: 'put', data}),
        add = async (data) => request({url: `/${urlPrefix}/add`, method: 'post', data}),
        del = async (id) => request({url: `/${urlPrefix}/del?id=${id}`, method: 'delete'})
    const [params, setParams] = useState({...defaultPageParam}), [paramsTemp, setParamsTemp] = useState(params),
        [details, setDetails] = useState({showType: "main"}),
        {data, mutate} = useSWR({url: `/${urlPrefix}/list`, params: params}, fetcher),
        {currPage, list, totalPage, totalCount} = data ? data.data : ""
    const onSearch = () => setParams(paramsTemp),
        onChange = (e, name) => setParamsTemp({...paramsTemp, [name]: e.target.value}),
        onDetailsChange = (e, name) => setDetails({...details, [name]: e.target.value}),
        onKeyDown = (e) => (e.key === 'Enter') ? params.page = 1 && onSearch() : "",
        onDelete = async (id) => (confirm("Are you sure?")) ? await del(id) && toast('ok') && mutate() : "",
        onSubmit = async () => {
            let {code, msg} = details.showType == "update" ? await update(details) : await add(details)
            if (code == 0) {
                toast(msg)
                mutate()
                setDetails({showType: "main"})
            }
        }
    return <>
        <Meta/>
        <Nav/>
        <header className={'container hr'}>
            <h1>{name}</h1>
            <section>
                <ul>
                    <SearchLi field={"century"} paramsTemp={paramsTemp} onChange={onChange} onKeyDown={onKeyDown}/>
                    <SearchLi field={"year"} paramsTemp={paramsTemp} onChange={onChange} onKeyDown={onKeyDown}/>
                    <SearchLi field={"type"} paramsTemp={paramsTemp} onChange={onChange} onKeyDown={onKeyDown}/>
                </ul>
                <ul>
                    <SearchLi field={"name"} paramsTemp={paramsTemp} onChange={onChange} onKeyDown={onKeyDown}/>
                    <SearchLi field={"summary"} paramsTemp={paramsTemp} onChange={onChange} onKeyDown={onKeyDown}/>
                    <SearchLi field={"tag"} paramsTemp={paramsTemp} onChange={onChange} onKeyDown={onKeyDown}/>
                </ul>
                <ul>
                    <input type="submit" value={'search'} onClick={() => onSearch()}/>
                    <input type="submit" value={'add'} onClick={() => setDetails({showType: "add"})}/>
                </ul>
            </section>
        </header>
        <main className={'container hr'} hidden={!(details.showType === "main")}>
            {!list ? "" : (
                <table>
                    <thead>
                    <tr>
                        {[...fields.filter(item => !item.hidden), {field: "Operation"}].map(item => <th key={item.field}>{item.label ? item.label : item.field}</th>)}
                    </tr>
                    </thead>
                    <tbody>
                    {list.map(item => <tr key={item.id}>
                        {fields.filter(item => !item.hidden).map(i => <td key={i.field + 's'} title={i.desc ? i.desc : item[i.field]}><span>{item[i.field]}</span></td>)}
                        <td>
                            <input type={'button'} value={'Edit'} onClick={() => setDetails({...item, showType: "update"})}/>
                            <input type={'button'} value={'Del'} onClick={() => onDelete(item.id)}/>
                        </td>
                    </tr>)}
                    </tbody>
                    <TableFooter params={params} setParams={setParams} current={currPage} tCount={totalCount} tPage={totalPage}/>
                </table>
            )}
        </main>
        <section hidden={!(details.showType === "add" || details.showType === "update")}>
            <h2>Step1 Set info</h2>
            <section>
                <ul>
                    <DetailsLi field={'century'} details={details} onDetailsChange={onDetailsChange}/>
                    <DetailsLi field={'year'} details={details} onDetailsChange={onDetailsChange}/>
                    <DetailsLi field={'month'} details={details} onDetailsChange={onDetailsChange}/>
                    <DetailsLi field={'day'} details={details} onDetailsChange={onDetailsChange}/>
                    <DetailsLi field={'type'} details={details} onDetailsChange={onDetailsChange}/>
                    <DetailsLi field={'type_details'} details={details} onDetailsChange={onDetailsChange}/>
                </ul>
                <ul>
                    <DetailsLi field={'name'} details={details} onDetailsChange={onDetailsChange}/>
                    <DetailsLi field={'tag'} details={details} onDetailsChange={onDetailsChange}/>
                    <DetailsLi field={'sort'} details={details} onDetailsChange={onDetailsChange}/>
                    <DetailsLi field={'status'} details={details} onDetailsChange={onDetailsChange}/>
                </ul>
                <ul>
                    <DetailsLi field={'summary'} height={6} details={details} onDetailsChange={onDetailsChange}/>
                    <DetailsLi field={'content'} height={3} details={details} onDetailsChange={onDetailsChange}/>
                </ul>
            </section>
            <h2>Step2 Go!</h2>
            <p>Be patient!Operation may be some time...</p>
            <input type={'button'} value={'Go Back'} onClick={() => setDetails({showType: "main"})}/>
            <input type={'button'} value={'Submit'} onClick={onSubmit}/>
        </section>
    </>
}
export default TimeHistory