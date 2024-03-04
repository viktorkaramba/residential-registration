import React, {useState} from "react";
import "./Menu.css"
import OSBBList from "../OSBB/OSBBList";
import OSBBForm from "../OSBB/OSBBForm";
import InhabitantForm from "../OSBB/Inhabitant/InhabitantForm";

const Menu = () =>{
    const [activeElement, setActiveElement] = useState('1');
    const handleClick = (element: React.SetStateAction<string>) => {
        setActiveElement(element);
    };

    return(
        <div>
            <section className='menu'>
                <div className='text-block' onClick={() => handleClick('1')}>
                    Список ОСББ
                </div>
                <div className='text-block' onClick={() => handleClick('2')}>
                    Додати ОСББ
                </div>
                <div className='text-block' onClick={() => handleClick('3')}>
                    Приєднатися до ОСББ
                </div>
            </section>
            {activeElement === '1' && <OSBBList/>}
            {activeElement === '2' && <OSBBForm/>}
            {activeElement === '3' && <InhabitantForm/>}
        </div>
    )
}

export default Menu