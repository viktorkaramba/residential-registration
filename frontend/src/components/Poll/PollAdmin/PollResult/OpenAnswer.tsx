import React from "react";
import {format} from "date-fns";

const OpenAnswer = ({answer}:any) => {
    return(
        <div className={'answer flex flex-sb'}>
            <label>
                <span className={'m-5'}>{answer.content}</span>
            </label>
            <div className='announcements-item-info-item fw-6 fs-15 m-5'>
                <span>{format(answer.created_at, 'MMMM do yyyy, hh:mm:ss a')}</span>
            </div>
        </div>
    )
}

export default OpenAnswer