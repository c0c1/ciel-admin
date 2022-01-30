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
    {field: 'm_name',},
    {field: 'r_name',}
]
const name = 'RoleMenus',
    urlPrefix = "roleMenu",
    noEdit = 1
const RoleMenu = ({query}) => {
    const fetcher = async ({url, params}) => request({url: url, method: 'get', params}),
        add = async (data) => request({url: `/${urlPrefix}/add`, method: 'post', data}),
        del = async (id) => request({url: `/${urlPrefix}/del?id=${id}`, method: 'delete'}),
        noMenus = async (rid) => request({url: `${urlPrefix}/noMenus?rid=${rid}`, method: 'get'})
    const [params, setParams] = useState({...defaultPageParam, ...query}), [paramsTemp, setParamsTemp] = useState(params),
        [details, setDetails] = useState({showType: "main"}),
        {data, mutate} = useSWR({url: `/${urlPrefix}/list`, params: params}, fetcher),
        {currPage, list, totalPage, totalCount} = data ? data.data : "",
        [noMenusOptions, setNoMenusOptions] = useState([]),
        [rid, setRid] = useState(0)
    useEffect(async () => {
        setParams({...params, ...query})
        setDetails({showType: 'main'})
        setRid(query.rid)
        setNoMenusOptions([])
    }, [query])
    const
        onDetailsChange = (e, name) => {
            if (name == 'mid') return setDetails({
                ...details, [name]: Array.from(e.target.selectedOptions, option => {
                    if (option.value) return option.value
                })
            })
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
            let {data} = await noMenus(query.rid)
            if (data) {
                setNoMenusOptions(data.map(item => {
                    return {
                        label: item.name,
                        value: item.id
                    }
                }))
            }
            setDetails({...details, showType: 'add'})
        }
    return <>
        <Meta/>
        <header className={'hr'}>
            <h2>{name}-{rid}</h2>
            <section>
                <ul><input type="submit" value={'add'} onClick={() => onShowNoMenus()}/></ul>
            </section>
        </header>
        <main hidden={!(details.showType === "main")}>
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
                        label={'菜单选择'}
                        field={'mid'}
                        details={details}
                        items={noMenusOptions}
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
export default RoleMenu