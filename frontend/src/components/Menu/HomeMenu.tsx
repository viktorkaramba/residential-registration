import React from "react";
import "./Menu.css"

import Header from "../Header/Header";
import OSBBList from "../OSBB/OSBBList/OSBBList";
import OSBBForm from "../OSBB/OSBBForm/OSBBForm";
import InhabitantForm from "../Inhabitant/InhabitantForm/InhabitantForm";
import {useAppContext} from "../../AppContext";
import {Link} from "react-router-dom";
import { useNavigate } from "react-router-dom"

const HomeMenu = () =>{

    // @ts-ignore
    const {activeOSBBElement, setActiveOSBBElement} = useAppContext();
    // @ts-ignore
    const {isLogin} = useAppContext();
    const navigate = useNavigate()
    const handleClick = (element: React.SetStateAction<string>) => {
        if(element === '4'){
            navigate('/osbbs/profile')
        }else {
            setActiveOSBBElement(element);
        }
    };

    return(
        <div>
            <Header/>
            <section className='menu'>
                <div className='text-block' onClick={() => handleClick('OSBBList')}>
                    Список ОСББ
                </div>
                {!isLogin &&
                    <div className='text-block' onClick={() => handleClick('OSBBForm')}>
                        Додати ОСББ
                    </div>}
                {isLogin &&
                    <div className='text-block' onClick={() => handleClick('InhabitantForm')}>
                        Профіль ОСББ
                    </div>}
            </section>
            {activeOSBBElement === 'OSBBList' && <OSBBList/>}
            {activeOSBBElement === 'OSBBForm' && <OSBBForm/>}
            {activeOSBBElement === 'InhabitantForm' && <InhabitantForm/>}
        </div>
    )
}

export default HomeMenu