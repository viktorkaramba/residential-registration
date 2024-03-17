import React from "react";
import "./Menu.css"

import Header from "../Header/Header";
import OSBBList from "../OSBB/OSBBList/OSBBList";
import OSBBForm from "../OSBB/OSBBForm/OSBBForm";
import InhabitantForm from "../Inhabitant/InhabitantForm/InhabitantForm";
import {useAppContext} from "../../AppContext";
import {Link} from "react-router-dom";
import { useNavigate } from "react-router-dom"
import AnnouncementUserList from "../Announcements/AnnouncementListUser/AnnouncementUserList";
import PollUserList from "../Poll/PollUser/PollUserList";
import AnnouncementForm from "../Announcements/AnnouncementForm/AnnouncementForm";
import PollForm from "../Poll/PollForm/PollForm";
import ProfileHeader from "../Header/ProfileHeader";
import PollMenu from "./PollMenu";
import AnnouncementAdminList from "../Announcements/AnnouncementAdmin/AnnouncementAdminList";
import PollAdminList from "../Poll/PollAdmin/PollAdminList";

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
                    <div className='text-block' onClick={() => handleClick('')}>
                        Запити мешканців
                    </div>
                </section>
            }
            {activeOSBBElement === 'AnnouncementAdminItem' && <AnnouncementAdminList/>}
            {activeOSBBElement === 'PollMenu' && <PollMenu/>}

        </div>
    )
}

export default ProfileMenu