import React from "react";
import { format } from 'date-fns';

const AnnouncementUserItem = ((announcement:any) => {

    return(
        <div className='announcements-item'>
            <div className='flex flex-sb flex-wrap'>
                <div className='announcements-item-info-item fw-7 fs-26'>
                    <span> {announcement.Title}</span>
                </div>
                <div className='announcements-item-info-item fw-6 fs-15'>
                    <span> {format(announcement.CreatedAt, 'MMMM do yyyy, hh:mm:ss a')}</span>
                </div>
            </div>
            <div className={'announcements-item-info'}>
                <div className='announcements-item-info-item fw-6 fs-18'>
                    <span>{announcement.Content}</span>
                </div>
            </div>
        </div>
    )
});

export default AnnouncementUserItem