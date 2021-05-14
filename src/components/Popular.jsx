import React from 'react'
import PopularItem from './PopularItem'

export default function Popular(props) {
    return (
      <div className="popular__wrapper">
        {
            props.items.map((item, key) => {
                return <PopularItem key={`${item}_${key}`} {...props.items[key]} />
            })
        }
      </div>
    )
}
