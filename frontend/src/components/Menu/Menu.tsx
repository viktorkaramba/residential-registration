import React from "react";
import "./Menu.css"

import Header from "../Header/Header";
import OSBBList from "../OSBB/OSBBList/OSBBList";
import OSBBForm from "../OSBB/OSBBForm/OSBBForm";
import InhabitantForm from "../Inhabitant/InhabitantForm/InhabitantForm";
import {useOSBBContext} from "../OSBB/OSBBContext";

const Menu = () =>{

    // @ts-ignore
    const {activeOSBBElement, setActiveOSBBElement} = useOSBBContext();
    const handleClick = (element: React.SetStateAction<string>) => {
        setActiveOSBBElement(element);
    };

    return(
        <div>
            <Header/>
            <section className='menu'>
                <div className='text-block' onClick={() => handleClick('1')}>
                    Список ОСББ
                </div>
                <div className='text-block' onClick={() => handleClick('2')}>
                    Додати ОСББ
                </div>

            </section>
            {activeOSBBElement === '1' && <OSBBList/>}
            {activeOSBBElement === '2' && <OSBBForm/>}
            {activeOSBBElement === '3' && <InhabitantForm/>}
        </div>
    )
}

export default Menu