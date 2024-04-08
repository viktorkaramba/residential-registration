import React, {useCallback, useEffect, useState} from "react";
import "./Menu.css"

import {useAppContext} from "../../utils/AppContext";
import ProfileHeader from "../Header/ProfileHeader";
import PollMenu from "./PollMenu";
import AnnouncementAdminList from "../Announcements/AnnouncementAdmin/AnnouncementAdminList";
import InhabitantRequestList from "../Inhabitant/InhabitantRequest/InhabitantRequestList";
import ProfileForm from "../Profile/ProfileForm/ProfileForm";
import {useNavigate} from "react-router-dom";
import config from "../../utils/config";
import err from "../../utils/err";

const ProfileMenu = () => {

    // @ts-ignore
    const {activeOSBBElement, setActiveOSBBElement, setOsbbID, user} = useAppContext();
    const [is, setIS] = useState<any>(false);
    
    useEffect(()=>{
        setActiveOSBBElement('ProfileForm')
    },[])

    useEffect(()=>{
        if(user!=null){
            setIS(true)
            setOsbbID(user.osbbid)
            setActiveOSBBElement('ProfileForm')
        }
    },[user])

    const handleClick = (element: React.SetStateAction<string>) => {
        setActiveOSBBElement(element);
    };

    return(
        <div>
            <ProfileHeader/>
            {user?.role === 'osbb_head'   &&
                <section className='menu flex flex-c flex-wrap'>
                    <button className='menu-text m-5' onClick={() => handleClick('ProfileForm')}>
                        Профіль
                    </button>
                    <button className='menu-text m-5' onClick={() => handleClick('AnnouncementAdminList')}>
                        Оголошення
                    </button>
                    <button className='menu-text m-5' onClick={() => handleClick('PollMenu')}>
                        Опитування
                    </button>
                    <button className='menu-text m-5' onClick={() => handleClick('InhabitantRequestList')}>
                        Запити мешканців
                    </button>
                </section>
            }
            {activeOSBBElement === 'ProfileForm' && is && <ProfileForm profile_user={user}/>}
            {activeOSBBElement === 'AnnouncementAdminList' && <AnnouncementAdminList/>}
            {activeOSBBElement === 'PollMenu' && <PollMenu/>}
            {activeOSBBElement === 'InhabitantRequestList' && <InhabitantRequestList/>}
        </div>
    )
}

export default ProfileMenu