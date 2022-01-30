import DefaultPage from "../../components/defaultPage";

const Menu = () => <DefaultPage
    fields={[
        {field: 'id', editHidden: 1},
        {field: 'pid', search: 1, type: 'number'},
        {field: 'name', search: 1},
        {field: 'type', title: '类型', search: 1, type: "select", items: [{label: 'group', value: 1, bgGround: "#80deea"}, {label: 'point', value: 2, bgGround: "#e0e0e0"}]},
        {field: 'path'},
        {field: 'status', type: 'select', items: [{label: 'yes', value: 1}, {label: 'no', value: 2}]},
        {field: 'sort', type: 'number', step: '0.1', detailsDesc: "请给一个排序"},
        {field: 'created_at', editHidden: 1},
    ]} name={'Menu'} urlPrefix={'menu'}/>
export default Menu