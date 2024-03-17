import React from "react";
import "./Menu.css"

import {useAppContext} from "../../AppContext";
import PollForm from "../Poll/PollForm/PollForm";
import PollTestForm from "../Poll/PollForm/PollTestForm";
import PollAdminList from "../Poll/PollAdmin/PollAdminList";

const PollMenu = () => {

    // @ts-ignore
    const {activePollElement, setActivePollElement} = useAppContext();

    const handleClick = (element: React.SetStateAction<string>) => {
        setActivePollElement(element);
    };

    return(
        <div>
            <section className='menu'>
                <div className='text-block' onClick={() => handleClick('PollForm')}>
                    Додати опитування з відкритою відповідю
                </div>
                <div className='text-block' onClick={() => handleClick('PollTestForm')}>
                    Додати опитування у вигляді тесту
                </div>
            </section>
            {activePollElement === 'PollForm' && <PollForm/>}
            {activePollElement === 'PollTestForm' && <PollTestForm/>}
            <PollAdminList/>
        </div>
    )
}

export default PollMenu