import DefaultPage from "../../components/defaultPage";

const fields = [
    {field: 'id', editHidden: 1},
    {
        field: 'rid', search: 1, type: 'select', items: [
            {label: '管理员', value: 9},
            {label: '超级管理员', value: 8},
        ]
    },
    {field: 'uname', search: 1},
    {field: 'ex1', title: '修改用户名', type: 'select', items: [{label: 'yes', value: 1}, {label: 'no', value: 0}], hidden: 1},
    {field: 'pwd', hidden: 1},
    {field: 'desc'},
    {field: 'status', type: 'select', items: [{label: 'yes', value: 1}, {label: 'no', value: 2}]},
    {field: 'created_at', editHidden: 1},
]
const name = 'Admin'
const urlPrefix = "admin"
const Admin = () => <DefaultPage name={name} fields={fields} urlPrefix={urlPrefix}/>
export default Admin