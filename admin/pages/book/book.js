import DefaultPage from "../../components/defaultPage";

const fields = [
    {field: 'id', search: 0, editHidden: 1},
    {field: 'tid', hidden: 1},
    {field: 'category_id', search: 1},
    {field: 'writer_id', hidden: 1},
    {field: 'name'},
    {field: 'icon'},
    {field: 'summary', detailsHeight: 3},
    {field: 'first_chapter_id', hidden: 1},
    {field: 'type_details', hidden: 1},
    {field: 'hidden_words', hidden: 1},
    {field: 'publish_time'},
    {field: 'sort'},
    {field: 'theme', hidden: 1},
    {field: 'status', hidden: 1},
    {field: 'scope', hidden: 1},
    {field: 'click_num', hidden: 1},
    {field: 'created_at', editHidden: 1},
]
const Book = () => <DefaultPage fields={fields} name={'Book'} urlPrefix={"book"}/>
export default Book