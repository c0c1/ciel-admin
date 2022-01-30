import request from "../../utils/request";
import {useState} from "react";
import {defaultPageParam} from "../../data/const";
import useSWR from "swr";
import {toast} from "react-toastify";
import Meta from "../../components/meta";
import {Nav} from "../../components/nav";
import {SearchLi} from "../../components/searchLi";
import {TableTd} from "../../components/tableTd";
import {TableFooter} from "../../components/table";
import {DetailsLi2} from "../../components/detailsLi";
import RoleMenu from "./roleMenu";
import RoleApi from "./roleApi";

const fields = [
    {field: 'id', editHidden: 1},
    {field: 'name', search: 1},
    {field: 'status', type: 'select', items: [{label: 'yes', value: 1}, {label: 'no', value: 2}]},
    {field: 'created_at', editHidden: 1},
]
const name = 'Role'
const urlPrefix = "role"
const Role = () => {
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
        <header className={' hr'}>
            <h1>{name}</h1>
            <section>
                <ul>
                    {fields.filter(item => item.search).map(item =>
                        <SearchLi key={item.field} field={item.field} label={item.title} onChange={onChange} paramsTemp={paramsTemp} onKeyDown={onKeyDown}/>
                    )}
                </ul>
                <ul>
                    <input type="submit" value={'search'} onClick={() => onSearch()}/>
                    <input type="submit" value={'add'} onClick={() => setDetails({showType: "add"})}/>
                </ul>
            </section>
        </header>
        <main className={' hr'} hidden={!(details.showType === "main")}>
            {!list ? "" : (
                <table>
                    <thead>
                    <tr>
                        {[...fields.filter(item => !item.hidden), {field: "Operation"}].map(item =>
                            <th key={item.field}>{item.title ? item.title : item.field} </th>)}
                    </tr>
                    </thead>
                    <tbody>
                    {list.map(item => <tr key={item.id}>
                        {fields.filter(item => !item.hidden).map(i => (<TableTd key={i.field + 's'} value={item[i.field]} type={i.type} items={i.items}/>))}
                        <td>
                            <input type={'submit'} value={'Menu'} onClick={() => setDetails({...item, showType: "menu", rid: item.id})}/>
                            <input type={'submit'} value={'Api'} onClick={() => setDetails({...item, showType: "api", rid: item.id})}/>
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
                </ul>
            </section>
            <h2>Step2 Go!</h2>
            <p>Be patient!Operation may be some time...</p>
            <input type={'button'} value={'Go Back'} onClick={() => setDetails({showType: "main"})}/>
            <input type={'button'} value={'Submit'} onClick={onSubmit}/>
        </section>
        <section hidden={!(details.showType === 'menu')}>
            <RoleMenu query={{rid: details.rid}}/>
            <input type={'button'} onClick={() => setDetails({...details, showType: 'main'})} value={'back'}/>
        </section>
        <section hidden={!(details.showType === 'api')}>
            <RoleApi query={{rid: details.rid}}/>
            <input type={'button'} onClick={() => setDetails({...details, showType: 'main'})} value={'back'}/>
        </section>
    </>
}
export default Role