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

const OSBBProfileMenu = () => {

    // @ts-ignore
    const {activeOSBBElement, setActiveOSBBElement} = useAppContext();

    const handleClick = (element: React.SetStateAction<string>) => {
        setActiveOSBBElement(element);
    };

    return(
        <div>
            <Header/>
            <section className='menu'>
                <div className='text-block' onClick={() => handleClick('AnnouncementList')}>
                    Оголошення
                </div>
                <div className='text-block' onClick={() => handleClick('PollUserList')}>
                    Опитування
                </div>
            </section>
            {activeOSBBElement === 'AnnouncementList' && <AnnouncementList/>}
            {activeOSBBElement === 'PollUserList' && <PollUserList/>}
        </div>
    )
}

export default OSBBProfileMenu