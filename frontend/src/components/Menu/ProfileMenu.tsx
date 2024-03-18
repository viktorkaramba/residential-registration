import React from "react";
import "./Menu.css"

import {useAppContext} from "../../AppContext";
import ProfileHeader from "../Header/ProfileHeader";
import PollMenu from "./PollMenu";
import AnnouncementAdminList from "../Announcements/AnnouncementAdmin/AnnouncementAdminList";
import InhabitantRequestList from "../Inhabitant/InhabitantRequest/InhabitantRequestList";

const ProfileMenu = ({role}:any) => {

    // @ts-ignore
    const {activeOSBBElement, setActiveOSBBElement} = useAppContext();

    const handleClick = (element: React.SetStateAction<string>) => {
        setActiveOSBBElement(element);
    };

    return(
        <div>
            <ProfileHeader/>
            {role === 'osbb_head' &&
                <section className='menu'>
                    <div className='text-block' onClick={() => handleClick('AnnouncementAdminItem')}>
                        Оголошення
                    </div>
                    <div className='text-block' onClick={() => handleClick('PollMenu')}>
                        Опитування
                    </div>
                    <div className='text-block' onClick={() => handleClick('InhabitantRequestList')}>
                        Запити мешканців
                    </div>
                </section>
            }
            {activeOSBBElement === 'AnnouncementAdminItem' && <AnnouncementAdminList/>}
            {activeOSBBElement === 'PollMenu' && <PollMenu/>}
            {activeOSBBElement === 'InhabitantRequestList' && <InhabitantRequestList/>}

        </div>
    )
}

export default ProfileMenu