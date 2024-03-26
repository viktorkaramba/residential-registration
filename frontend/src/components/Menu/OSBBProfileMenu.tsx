import React, {useEffect} from "react";
import "./Menu.css"

import Header from "../Header/Header";
import {useAppContext} from "../../utils/AppContext";
import AnnouncementUserList from "../Announcements/AnnouncementListUser/AnnouncementUserList";
import PollUserList from "../Poll/PollUser/PollUserList";
import OSBBDescription from "../OSBB/OSBBDescription";

const OSBBProfileMenu = () => {

    // @ts-ignore
    const {activeOSBBElement, setActiveOSBBElement} = useAppContext();

    const handleClick = (element: React.SetStateAction<string>) => {
        setActiveOSBBElement(element);
    };

    useEffect(()=>{
        setActiveOSBBElement('OSBBDescription')
    },[])
    return(
        <div>
            <Header  withWelcomeBlock={false}/>
            <section className='menu flex flex-c flex-wrap'>
                <button className='menu-text m-5' onClick={() => handleClick('OSBBDescription')}>
                    Профіль
                </button>
                <button className='menu-text m-5' onClick={() => handleClick('AnnouncementUserList')}>
                    Оголошення
                </button>
                <button className='menu-text m-5' onClick={() => handleClick('PollUserList')}>
                    Опитування
                </button>
            </section>
            {activeOSBBElement === 'OSBBDescription' && <OSBBDescription/>}
            {activeOSBBElement === 'AnnouncementUserList' && <AnnouncementUserList/>}
            {activeOSBBElement === 'PollUserList' && <PollUserList/>}
        </div>
    )
}

export default OSBBProfileMenu