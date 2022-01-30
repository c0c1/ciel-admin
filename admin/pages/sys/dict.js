import DefaultPage from "../../components/defaultPage";

const fields = [
    {field: 'id', editHidden: 1},
    {field: 'k', search: 1},
    {field: 'v', search: 1},
    {field: 'desc'},
    {
        field: 'group', search: 1, type: 'select', items: [
            {label: 'system', value: '1'},
            {label: 'website', value: '2'},
            {label: 'book', value: '3'},
        ]
    },
    {
        field: 'type', search: 1, type: 'select', items: [
            {label: 'img', value: 1},
            {label: 'icon', value: 2},
            {label: 'file', value: 3},
            {label: 'text', value: 4},
            {label: 'html', value: 5},
            {label: 'other', value: 6},
        ]
    },
    {field: 'status', type: 'select', items: [{label: 'yes', value: 1}, {label: 'no', value: 2}]},
    {field: 'created_at', editHidden: 1},
]
const Dict = () => <DefaultPage name={'Dict'} urlPrefix={'dict'} fields={fields}/>
export default Dict