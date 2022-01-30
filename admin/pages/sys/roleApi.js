import React, {useEffect, useState} from "react";
import request from "../../utils/request";
import {defaultPageParam} from "../../data/const";
import useSWR from "swr";
import {toast} from "react-toastify";
import Meta from "../../components/meta";
import {TableTd} from "../../components/tableTd";
import {TableFooter} from "../../components/table";
import {DetailsLi2} from "../../components/detailsLi";

const fields = [
    {field: 'id', editHidden: 1},
    {field: 'role_name'},
    {field: 'group'},
    {field: 'method'},
    {field: 'url'},
]
const name = 'RoleApi',
    urlPrefix = "roleApi",
    noEdit = 1
const RoleApi = ({query}) => {
    const
        fetcher = async ({url, params}) => request({url: url, method: 'get', params}),
        add = async (data) => request({url: `/${urlPrefix}/add`, method: 'post', data}),
        del = async (id) => request({url: `/${urlPrefix}/del?id=${id}`, method: 'delete'}),
        noApis = async (rid) => request({url: `${urlPrefix}/noApis?rid=${rid}`, method: 'get'})
    const
        [params, setParams] = useState({...defaultPageParam, ...query}),
        [details, setDetails] = useState({showType: "main"}),
        {data, mutate} = useSWR({url: `/${urlPrefix}/list`, params: params}, fetcher),
        {currPage, list, totalPage, totalCount} = data ? data.data : "",
        [noApisOptions, setNoApisOptions] = useState([]),
        [rid, setRid] = useState(0)
    const
        onDetailsChange = (e, name) => {
            if (name == 'aid') {
                return setDetails({
                    ...details, [name]: Array.from(e.target.selectedOptions, option => {
                        if (option.value) return option.value
                    })
                })
            }
            setDetails({...details, [name]: e.target.value})
        },
        onDelete = async (id) => (confirm("Are you sure?")) ? await del(id) && toast('ok') && mutate() : "",
        onSubmit = async () => {
            await add({...details, 'rid': rid})
            mutate()
            toast('ok')
            setDetails({showType: "main"})
        },
        onShowNoMenus = async () => {
            let {data} = await noApis(query.rid)
            if (data) {
                setNoApisOptions(data.map(item => {
                    const {desc, group, method, url} = item
                    return {label: `${group} ${url} ${method} ${desc}`, value: item.id}
                }))
            }
            setDetails({...details, showType: 'add'})
        }
    useEffect(async () => {
        setParams({...params, ...query})
        setRid(query.rid)
        setDetails({showType: 'main'})
    }, [query])
    return <>
        <Meta/>
        <header className={'hr'}>
            <h2>{name}-{rid}</h2>
            <section>
                <ul><input type="submit" value={'add'} onClick={() => onShowNoMenus()}/></ul>
            </section>
        </header>
        <main className={'hr'} hidden={!(details.showType === "main")}>
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
                            <input hidden={noEdit} type={'submit'} value={'Edit'} onClick={() => setDetails({...item, showType: "update"})}/>
                            <input type={'submit'} value={'Del'} onClick={() => onDelete(item.id)}/>
                        </td>
                    </tr>)}
                    </tbody>
                    <TableFooter params={params} setParams={setParams} current={currPage} tCount={totalCount} tPage={totalPage}/>
                </table>
            )}
        </main>
        <section className={'details'} hidden={!(details.showType === "add" || details.showType === "update")}>
            <section>
                <ul>
                    <DetailsLi2
                        label={'API选择'}
                        field={'aid'}
                        details={details}
                        items={noApisOptions}
                        type={'select_multiple'}
                        onDetailsChange={onDetailsChange}
                    />
                </ul>
            </section>
            <input type={'button'} value={'Go Back'} onClick={() => setDetails({showType: "main"})}/>
            <input type={'button'} value={'Submit'} onClick={onSubmit}/>
        </section>
    </>
}
export default RoleApi