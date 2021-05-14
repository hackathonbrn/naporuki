import React from 'react';

export default function Category(props) {
    return (
        <li onClick={props.onClick} className={props.className}>{props.categoryName}</li>
    )
}
