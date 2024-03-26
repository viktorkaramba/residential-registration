import React, {useEffect} from "react";
import "./Menu.css"

import {useAppContext} from "../../utils/AppContext";
import PollForm from "../Poll/PollForm/PollForm";
import PollTestForm from "../Poll/PollForm/PollTestForm";
import PollAdminList from "../Poll/PollAdmin/PollAdminList";

const PollMenu = () => {

    // @ts-ignore
    const {activePollElement, setActivePollElement} = useAppContext();

    const handleClick = (element: React.SetStateAction<string>) => {
        setActivePollElement(element);
    };

    useEffect(()=>{
        setActivePollElement('PollAdminList')
    },[])

    return(
        <div>
            <section className='menu flex flex-c'>
                <button className='menu-text m-5' onClick={() => handleClick('PollAdminList')}>
                    Опитування
                </button>
                <button className='menu-text m-5' onClick={() => handleClick('PollForm')}>
                    Додати опитування з відкритою відповідю
                </button>
                <button className='menu-text m-5' onClick={() => handleClick('PollTestForm')}>
                    Додати опитування у вигляді тесту
                </button>
            </section>
            {activePollElement === 'PollAdminList' && <PollAdminList/>}
            {activePollElement === 'PollForm' && <PollForm/>}
            {activePollElement === 'PollTestForm' && <PollTestForm/>}
        </div>
    )
}

export default PollMenu