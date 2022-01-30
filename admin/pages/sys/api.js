import DefaultPage from "../../components/defaultPage";


const fields = [
    {field: 'id', editHidden: 1,},
    {field: 'url', search: 1},
    {field: 'method', search: 1},
    {field: 'group'},
    {field: 'desc'},
    {field: 'status', type: 'select', items: [{label: 'yes', value: 1}, {label: 'no', value: 2}]},
    {field: 'created_at', editHidden: 1},
]
const Api = () => <DefaultPage fields={fields} name={'Api'} urlPrefix={'api'}/>
export default Api