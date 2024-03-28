import React, {useEffect} from "react";
import "./Menu.css"

import Header from "../Header/Header";
import OSBBList from "../OSBB/OSBBList/OSBBList";
import OSBBForm from "../OSBB/OSBBForm/OSBBForm";
import InhabitantForm from "../Inhabitant/InhabitantForm/InhabitantForm";
import {useAppContext} from "../../utils/AppContext";
import { useNavigate } from "react-router-dom"

const HomeMenu = () =>{

    // @ts-ignore
    const {activeOSBBElement, setActiveOSBBElement} = useAppContext();
    // @ts-ignore
    const {isLogin} = useAppContext();
    const navigate = useNavigate()
    const handleClick = (element: React.SetStateAction<string>) => {
        if(element === 'OSBBProfile'){
            navigate('/osbbs/profile')
        }else {
            setActiveOSBBElement(element);
        }
    };

    useEffect(()=>{
        setActiveOSBBElement('OSBBList');
    }, [])
    return(
        <div>
            <Header  withWelcomeBlock={true}/>
            <section className='menu flex flex-c'>
                <button className='menu-text m-5' onClick={() => handleClick('OSBBList')}>
                    Список ОСББ
                </button>
                {!isLogin &&
                    <button className='menu-text m-5' onClick={() => handleClick('OSBBForm')}>
                        Додати ОСББ
                    </button>}
                {isLogin &&
                    <button className='menu-text m-5' onClick={() => handleClick('OSBBProfile')}>
                        Профіль ОСББ
                    </button>}
            </section>
            {activeOSBBElement === 'OSBBList' && <OSBBList/>}
            {activeOSBBElement === 'OSBBForm' && <OSBBForm/>}
            {activeOSBBElement === 'InhabitantForm' && <InhabitantForm/>}
        </div>
    )
}

export default HomeMenu