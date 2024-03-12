import React from "react";
import "./Menu.css"

import Header from "../Header/Header";
import OSBBList from "../OSBB/OSBBList/OSBBList";
import OSBBForm from "../OSBB/OSBBForm/OSBBForm";
import InhabitantForm from "../Inhabitant/InhabitantForm/InhabitantForm";
import {useOSBBContext} from "../OSBB/OSBBContext";
import {Link} from "react-router-dom";
import { useNavigate } from "react-router-dom"

const Menu = () =>{

    // @ts-ignore
    const {activeOSBBElement, setActiveOSBBElement} = useOSBBContext();
    // @ts-ignore
    const {isLogin} = useOSBBContext();
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
                <div className='text-block' onClick={() => handleClick('1')}>
                    Список ОСББ
                </div>
                {!isLogin &&
                    <div className='text-block' onClick={() => handleClick('2')}>
                        Додати ОСББ
                    </div>}
                {isLogin &&
                    <div className='text-block' onClick={() => handleClick('4')}>
                        Профіль ОСББ
                    </div>}
            </section>
            {activeOSBBElement === '1' && <OSBBList/>}
            {activeOSBBElement === '2' && <OSBBForm/>}
            {activeOSBBElement === '3' && <InhabitantForm/>}
        </div>
    )
}

export default Menu