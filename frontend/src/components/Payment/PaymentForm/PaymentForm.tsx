import React, {useEffect, useRef, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import Alert from "@mui/material/Alert";
import {Stack} from "@mui/material";

const PaymentForm = ({addPayment}:any) =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [isSuccess, setIsSuccess]= useState(false);
    const navigate = useNavigate();
    const focusRef = useRef(null);
    useEffect(() => {
        // @ts-ignore
        focusRef.current.scrollIntoView({behavior: 'smooth'});
    }, []);

    const handleAddPayment = (event: any) => {
        event.preventDefault();

        // 👇️ access input values using name prop
        const amount = event.target.amount.value;
        const appointment = event.target.appointment.value;
        const requestOptions = {
            method: 'POST',
            headers: config.headers,
            body: JSON.stringify({ amount: amount, appointment: appointment})
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/payments', requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:handleAddPayment, navigate:navigate});
                }else {
                    const {id}:any = data
                    const newPayment = {
                        ID: id,
                        OSBBID: osbbID,
                        Amount: amount,
                        Appointment: appointment,
                        CreatedAt:new Date()
                    };
                    addPayment(newPayment)
                    setIsSuccess(true);
                }
            });
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };

    return(
        <form method='post'  onSubmit={handleAddPayment} ref={focusRef} style={{margin:"35px"}}>
            <div className={'flex flex-wrap align-items-start'}>
                <div className="form announcement_form">
                    <h1>Ціна сплати та призначення платежу</h1>
                    <div className="inner-wrap">
                        <label form={'amount'}>Вартість
                            <input required={true} name="amount" placeholder="" type='number' id='amount'/>
                        </label>
                        <label form={'appointment'}>Призначення
                            <textarea minLength={2} required={true} name="appointment" placeholder="" id='appointment'/>
                        </label>
                    </div>
                    <div className={'flex flex-c'}>
                        <button className='button add_osbb' type="submit" name="submit_osbb">
                            <span className="button_content add_osbb_content">Додати Платіж</span>
                        </button>
                    </div>
                    {isSuccess &&
                        <Stack sx={{margin: '10px'}} spacing={2}>
                            <Alert variant={'filled'} severity="success" style={{fontSize:'15px'}}>Платіж успішно доданий!</Alert>
                        </Stack>}
                </div>
            </div>
        </form>
    )
}

export default PaymentForm