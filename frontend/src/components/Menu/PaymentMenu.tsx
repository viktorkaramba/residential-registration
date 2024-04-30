import React, {useEffect} from "react";
import "./Menu.css"

import {useAppContext} from "../../utils/AppContext";
import PollForm from "../Poll/PollForm/PollForm";
import PollTestForm from "../Poll/PollForm/PollTestForm";
import PollAdminList from "../Poll/PollAdmin/PollAdminList";
import PollResultItem from "../Poll/PollAdmin/PollResult/PollResultItem";
import PaymentForm from "../Payment/PaymentForm/PaymentForm";

const PaymentMenu = () => {

    // @ts-ignore
    const {activePollElement, payment, setActivePollElement} = useAppContext();

    const handleClick = (element: React.SetStateAction<string>) => {
        setActivePollElement(element);
    };

    useEffect(()=>{
        setActivePollElement('PaymentForm')
    },[])

    return(
        <div>
            <section className='menu flex flex-c flex-wrap'>
                <button className='menu-text m-5' onClick={() => handleClick('PaymentForm')}>
                    Додати платіж
                </button>
                <button className='menu-text m-5' onClick={() => handleClick('PaymentAdmin')}>
                    Звіти про платежі
                </button>
            </section>
            {activePollElement === 'PaymentForm' && <PaymentForm/>}
            {/*{activePollElement === 'PaymentAdmin' && <PaymentReport/>}*/}
        </div>
    )
}

export default PaymentMenu