import {bgGround} from "../utils/utils";

export const SearchLi = ({type, items, field, label, paramsTemp, onChange, onKeyDown}) => {
    let title = label ? label : field
    switch (type) {
        case 'select':
            return (
                <li>
                    {title}
                    <select value={paramsTemp[field]} onChange={e => onChange(e, field)}>
                        <option value={''}>Chose</option>
                        {items.map(item =>
                            <option key={item.value}
                                    style={{background: bgGround(item.value, item.bgGround)}}
                                    value={item.value}>
                                {item.label}
                            </option>)}
                    </select>
                </li>
            )
        default:
            return (
                <li>
                    {title}
                    <input
                        type="text"
                        value={paramsTemp[field] || ''}
                        onChange={e => onChange(e, field)}
                        onKeyDown={onKeyDown}/>
                </li>
            );
    }
}

