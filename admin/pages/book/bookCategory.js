import DefaultPage from "../../components/defaultPage";

const fields = [
    {field: 'id', editHidden:1},
    {field: 'name', title: '', search: 0, desc: '',  detailsDesc: "", sys: 0},
    {field: 'icon', title: '', search: 1, desc: '',  detailsDesc: "", sys: 0},
    {field: 'color', title: '', hidden: 1, search: 0, desc: '',  detailsDesc: "", sys: 0},
    {field: 'hidden_words', hidden: 1, title: '', search: 0, desc: '',  detailsDesc: "", sys: 0},
    {field: 'type', title: '', search: 0, desc: '',  detailsDesc: "", sys: 0},
    {field: 'sort', title: '', search: 0, desc: '',  detailsDesc: "", sys: 0},
    {field: 'status', title: '', search: 0, desc: '',  detailsDesc: "", sys: 0},
    {field: 'created_at', editHidden: 1},
]
const BookCategory = () => <DefaultPage fields={fields} name={'BookCategory'} urlPrefix={"bookCategory"}/>
export default BookCategory