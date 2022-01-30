export function Table({fields, params, setParams, currPage, totalCount, totalPage, children}) {
    return <table>
        <thead>
        <tr>{fields.map(item => <th key={item}>{item}</th>)}</tr>
        </thead>
        <tbody>
        {children}
        </tbody>
        <TableFooter params={params} setParams={setParams} current={currPage} tCount={totalCount} tPage={totalPage}/>
    </table>;
}

export const TableFooter = ({setParams, params, current, tCount, tPage}) => {
    current = +current || 1
    let before = params.page - 1 <= 0 ? 1 : params.page - 1
    let next = params.page + 1 >= tPage ? tPage : params.page + 1
    const toPage = ({p, size}) => setParams({...params, page: p, size: size})
    return <tfoot>
    <tr>
        <td colSpan={100}>
            <ul>
                page
                <select onChange={e => toPage({p: params.page, size: e.target.value})}>
                    <option value={10}>10</option>
                    <option value={5}>5</option>
                    <option value={50}>50</option>
                    <option value={100}>100</option>
                    <option value={200}>200</option>
                </select>
                <li><a href="#" onClick={() => setParams({...params, page: 1})}>1</a></li>
                <span hidden={tPage < 3}>...</span>
                <li hidden={tPage < 2}><a href="#" onClick={() => setParams({...params, page: tPage})}>{tPage}</a></li>
                <li><a href="#" onClick={() => setParams({...params, page: before})}>Back</a></li>
                <li><a href="#" onClick={() => setParams({...params, page: next})}>Next</a></li>
                <span>{current}</span>
            </ul>
        </td>
    </tr>
    </tfoot>;
}