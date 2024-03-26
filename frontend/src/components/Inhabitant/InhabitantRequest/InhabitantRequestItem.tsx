import React from "react";
import {format} from "date-fns";

const InhabitantRequestItem = ({inhabitantWaitApprove, updateStatus}:any) => {
    return(
        <div className='announcements-item'>
            <div className='flex flex-sb flex-wrap'>
                <div className='announcements-item-info-item fw-7 fs-26'>
                    ПІБ: <span>{inhabitantWaitApprove.full_name.first_name}  {inhabitantWaitApprove.full_name.surname}  {inhabitantWaitApprove.full_name.patronymic}</span>
                </div>
                <div className='announcements-item-info-item fw-6 fs-15'>
                    <span>{format(inhabitantWaitApprove.createdAt, 'MMMM do yyyy, hh:mm:ss a')}</span>
                </div>
            </div>
            <div className={'announcements-item-info'}>
                <div className='announcements-item-info-item fw-6 fs-15'>
                    Номер телефону: <span>{inhabitantWaitApprove.phone_number}</span>
                </div>
            </div>
            <div className={'announcements-item-info'}>
                <div className='announcements-item-info-item fw-6 fs-15'>
                    Номер квартири: <span>{inhabitantWaitApprove.apartment.number}</span>
                </div>
            </div>
            <div className={'announcements-item-info'}>
                <div className='announcements-item-info-item fw-6 fs-15'>
                    Площа квартири: <span>{inhabitantWaitApprove.apartment.area}</span>
                </div>
            </div>
            <div className={'flex flex-c'}>
                <button className='button announcement_button' style={{marginTop:'10px', marginRight:'5px'}} onClick={()=>updateStatus(inhabitantWaitApprove.id, true)}>
                    <span className="button_content announcement_button_content" >Підтвердити</span>
                </button>
                <button className='button announcement_button' style={{marginTop:'10px'}} onClick={()=>updateStatus(inhabitantWaitApprove.id, false)}>
                    <span className="button_content announcement_button_content" >Відхилити</span>
                </button>
            </div>
        </div>
    )
}

export default InhabitantRequestItem