import DefaultPage from "../../components/defaultPage";

const fields = [
    {field: 'id', search: 0, details: 0},
    {field: 'book_id'},
    {field: 'name', search: 1},
    {field: 'hidden_words', hidden: 1},
    {field: 'type'},
    {field: 'created_at'},
]
const BookChapter = () => <DefaultPage fields={fields} name={'章节'} urlPrefix={"bookChapter"}/>
export default BookChapter