import {bgGround} from "../utils/utils";

export const DetailsLi = ({label, field, height, details, onDetailsChange, placeholder}) => {
    return height ? <li>{label ? label : field}<textarea rows={height} value={details[field] || ''} onChange={e => onDetailsChange(e, field)} placeholder={placeholder}/></li>
        : <li>{label ? label : field}<input value={details[field] || ''} onChange={e => onDetailsChange(e, field)} placeholder={placeholder}/></li>
}

export const DetailsLi2 = ({type, step, label, field, height, items, details, onDetailsChange, placeholder}) => {
    let title = label ? label : field
    switch (type) {
        case 'select_multiple':
        case 'select':
            return <li>
                {title}
                <select multiple={type == 'select_multiple'} value={details[field]} onChange={e => onDetailsChange(e, field)}>
                    <option value={''}>Chose</option>
                    {items.map(item =>
                        <option key={item.value}
                                style={{background: bgGround(item.value, item.bgGround)}}
                                value={item.value}>
                            {item.label}
                        </option>)}
                </select>
            </li>
        case "number":
            return <li>{title}<input type={"number"} step={step} value={details[field] || ''} onChange={e => onDetailsChange(e, field)} placeholder={placeholder}/></li>
        case 'textarea':
            return <li>{title}<textarea rows={height} value={details[field] || ''} onChange={e => onDetailsChange(e, field)} placeholder={placeholder}/></li>
        default:
            return <li>{title}<input value={details[field] || ''} onChange={e => onDetailsChange(e, field)} placeholder={placeholder}/></li>
    }
}
