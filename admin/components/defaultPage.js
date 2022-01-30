import useSWR from "swr";
import {useEffect, useState} from "react";
import {toast} from "react-toastify";
import Meta from "./meta";
import {Nav} from "./nav";
import {TableFooter} from "./table";
import {defaultPageParam} from "../data/const";
import request from "../utils/request";
import {DetailsLi2} from "./detailsLi";
import {SearchLi} from "./searchLi";
import {TableTd} from "./tableTd";

// defaultPage
const DefaultPage = ({
                         name, fields, urlPrefix,
                         noEdit, noNav, query, noSearch// options
                     }) => {
    const fetcher = async ({url, params}) => request({url: url, method: 'get', params}),
        update = async (data) => request({url: `/${urlPrefix}/update`, method: 'put', data}),
        add = async (data) => request({url: `/${urlPrefix}/add`, method: 'post', data}),
        del = async (id) => request({url: `/${urlPrefix}/del?id=${id}`, method: 'delete'})
    const [params, setParams] = useState({...defaultPageParam, ...query}), [paramsTemp, setParamsTemp] = useState(params),
        [details, setDetails] = useState({showType: "main"}),
        {data, mutate} = useSWR({url: `/${urlPrefix}/list`, params: params}, fetcher),
        {currPage, list, totalPage, totalCount} = data ? data.data : ""
    useEffect(() => setParams({...params, ...query}), [query])
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
    let tbody = ''
    let search = <ul>
        {fields.filter(item => item.search).map(item =>
            <SearchLi key={item.field}
                      field={item.field} label={item.title}
                      type={item.type} items={item.items}
                      onChange={onChange} paramsTemp={paramsTemp} onKeyDown={onKeyDown}
            />
        )}
    </ul>
    if (list) {
        tbody = list.map(item => <tr key={item.id}>
            {fields.filter(item => !item.hidden).map(i => (<TableTd key={i.field + 's'} value={item[i.field]} type={i.type} items={i.items}/>))}
            <td>
                <input hidden={noEdit} type={'submit'} value={'Edit'} onClick={() => setDetails({...item, showType: "update"})}/>
                <input type={'submit'} value={'Del'} onClick={() => onDelete(item.id)}/>
            </td>
        </tr>)
    }
    return <>
        <Meta/>
        {noNav ? '' : <Nav/>}
        <header className={' hr'}>
            <h1>{name}</h1>
            <section>
                {search}
                <ul>
                    <input hidden={noSearch} type="submit" value={'Search'} onClick={() => onSearch()}/>
                    <input type="submit" value={`Add ${name}`} onClick={() => setDetails({showType: "add", rid: ''})}/>
                </ul>
            </section>
        </header>
        <main className={' hr'} hidden={!(details.showType === "main")}>
            <table>
                <thead>
                <tr>
                    {[...fields.filter(item => !item.hidden), {field: "Operation"}].map(item =>
                        <th key={item.field}>{item.title ? item.title : item.field} </th>)}
                </tr>
                </thead>
                <tbody>{tbody}</tbody>
                <TableFooter params={params} setParams={setParams} current={currPage} tCount={totalCount} tPage={totalPage}/>
            </table>
        </main>
        <section className={'details'} hidden={!(details.showType === "add" || details.showType === "update")}>
            <h2>Step1 Set info</h2>
            <section>
                <ul> {fields.filter(item => !item.editHidden).map(item =>
                    <DetailsLi2
                        key={item.field} field={item.field} type={item.type} details={details} onDetailsChange={onDetailsChange}
                        label={item.title} placeholder={item.detailsDesc}
                        items={item.items} step={item.step} height={item.detailsHeight}
                    />)}
                </ul>
            </section>
            <h2>Step2 Go!</h2>
            <p>Be patient!Operation may be some time...</p>
            <input type={'button'} value={'Go Back'} onClick={() => setDetails({showType: "main"})}/>
            <input type={'button'} value={'Submit'} onClick={onSubmit}/>
        </section>
    </>
}
export default DefaultPage