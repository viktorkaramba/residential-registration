import React from "react";
import "./Menu.css"

import Header from "../Header/Header";
import OSBBList from "../OSBB/OSBBList/OSBBList";
import OSBBForm from "../OSBB/OSBBForm/OSBBForm";
import InhabitantForm from "../Inhabitant/InhabitantForm/InhabitantForm";
import {useAppContext} from "../../AppContext";
import {Link} from "react-router-dom";
import { useNavigate } from "react-router-dom"
import AnnouncementList from "../Announcements/AnnouncementList/AnnouncementList";
import PollUserList from "../Poll/PollUser/PollUserList";
import AnnouncementForm from "../Announcements/AnnouncementForm/AnnouncementForm";
import PollForm from "../Poll/PollForm/PollForm";
import ProfileHeader from "../Header/ProfileHeader";
import PollMenu from "./PollMenu";

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
                    <div className='text-block' onClick={() => handleClick('AnnouncementForm')}>
                        Додати оголошення
                    </div>
                    <div className='text-block' onClick={() => handleClick('PollMenu')}>
                        Додати опитування
                    </div>
                    <div className='text-block' onClick={() => handleClick('')}>
                        Запити мешканців
                    </div>
                </section>
            }
            {activeOSBBElement === 'AnnouncementForm' && <AnnouncementForm/>}
            {activeOSBBElement === 'PollMenu' && <PollMenu/>}
        </div>
    )
}

export default ProfileMenu