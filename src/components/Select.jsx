import React, { useState } from "react";

export default function Select(props) {
    const [active, setActive] = useState(false);
    return <option className="teacher-form__option category" value={props.value} onClick={() => {setActive(!active)}} >{props.value}</option>

}