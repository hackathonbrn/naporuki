import React, { useState } from 'react'
import Category from './Category'

export default function Categories(props) {
    React.useState(null);
    const [active, setActive] = useState(null);
    return (
        <ul className="categories__wrapper">
            <Category categoryName="Все" onClick={() => {setActive(null)}} className={`category ${active === null ? 'category--active' : '' }`}/>
            {
                props.items.map((category, key) => {
                    return <Category onClick={() => {setActive(key)}} className={`category ${active === key ? 'category--active' : '' }`} key={`${category}_${key}`} categoryName={category} />
                })
            }
        </ul>
    )
}
