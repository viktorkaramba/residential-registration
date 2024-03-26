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
    const {activeOSBBElement, setActiveOSBBElement} = useAppContext();
    // @ts-ignore
    const {token} = useAppContext();
    const navigate = useNavigate();
    // @ts-ignore
    const {osbbID, setOsbbID} = useAppContext()
    const [currentUser, setCurrentUser] = useState<any>(null);
    const [is, setIS] = useState<any>(false);
    const fetchUserProfile = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers:{ 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            };
            fetch(config.apiUrl+'osbb/' + osbbID + '/inhabitants/profile', requestOptions)
                .then(response => response.json())
                .then(data =>{
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchUserProfile, navigate:navigate});
                    }else {
                        if(data){
                            const {osbbid, apartment, full_name, phone_number, role}:any = data;
                            const userInfo = {
                                osbbid: osbbid,
                                apartment: apartment,
                                full_name: full_name,
                                phone_number: phone_number,
                                role : role,
                            };
                            setCurrentUser(userInfo);

                        }else {
                            setCurrentUser(null);
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, [token]);

    useEffect(()=>{
        setActiveOSBBElement('ProfileForm')
        fetchUserProfile();
    },[])

    useEffect(()=>{
        if(currentUser!=null){
            setIS(true)
            setOsbbID(currentUser.osbbid)
            setActiveOSBBElement('ProfileForm')
        }
    },[currentUser])

    const handleClick = (element: React.SetStateAction<string>) => {
        setActiveOSBBElement(element);
    };



    return(
        <div>
            <ProfileHeader/>
            {currentUser?.role === 'osbb_head'   &&
                <section className='menu flex flex-c'>
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
            {activeOSBBElement === 'ProfileForm' && is && <ProfileForm profile_user={currentUser}/>}
            {activeOSBBElement === 'AnnouncementAdminList' && <AnnouncementAdminList/>}
            {activeOSBBElement === 'PollMenu' && <PollMenu/>}
            {activeOSBBElement === 'InhabitantRequestList' && <InhabitantRequestList/>}
        </div>
    )
}

export default ProfileMenu