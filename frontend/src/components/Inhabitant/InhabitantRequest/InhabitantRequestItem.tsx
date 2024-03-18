import React from "react";

const InhabitantRequestItem = ({inhabitantWaitApprove, updateStatus}:any) => {
    return(
        <li>
            <div>
                {inhabitantWaitApprove.full_name.first_name}
                <br/>
                {inhabitantWaitApprove.full_name.surname}
                <br/>
                {inhabitantWaitApprove.full_name.patronymic}
                <br/>
                {inhabitantWaitApprove.phone_number}
            </div>
            <div>
                <button onClick={()=>updateStatus(inhabitantWaitApprove.id, true)}>Підтвердити</button>
                <button onClick={()=>updateStatus(inhabitantWaitApprove.id, false)}>Відхилити</button>
            </div>
        </li>
    )
}

export default InhabitantRequestItem