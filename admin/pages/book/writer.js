import DefaultPage from "../../components/defaultPage";

const fields = [
    {field: 'id', search: 0, details: 0},
    {field: 'tid', hidden: 1},
    {field: 'name', search: 1},
    {field: 'icon',hidden: 1},
    {field: 'summary'},
    {field: 'ex1', hidden: 1},
    {field: 'words', hidden: 1},
    {field: 'sort',hidden:1},
    {field: 'status',hidden: 1},
    {field: 'created_at'},
]
const Writer = () => <DefaultPage fields={fields} name={'Writer'} urlPrefix={"writer"}/>
export default Writer