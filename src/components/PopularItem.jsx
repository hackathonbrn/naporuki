import React from 'react'

export default function PopularItem(props) {
    return (
        <div className="popular-item">
          <img
            className="popular-item__photo"
            src={props.photo}
            alt={`Фото ${props.name}`}
          />
          <div className="popular-item__info">
            <h3 className="popular-item__name">{props.name}</h3>
            <span className="rating">{`Рейтинг: ${props.rating}/5`}</span>
          </div>
        </div>
    )
}
