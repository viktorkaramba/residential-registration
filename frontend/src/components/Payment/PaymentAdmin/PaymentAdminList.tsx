import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import {MdAddCircle} from "react-icons/md";
import {IoListCircleSharp} from "react-icons/io5";
import PaymentForm from "../PaymentForm/PaymentForm";
import PaymentAdminItem from "./PaymentAdminItem";

const PaymentAdminList = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()

    const [payments, setPayments] = useState([]);
    const navigate = useNavigate();
    const [isAddedChecked, setIsAddedChecked] = useState(false);

    const fetchPayments = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: { 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            }
            fetch(config.apiUrl+'osbb/'+ osbbID+ '/payments', requestOptions)
                .then(response => response.json())
                .then(data=>{
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchPayments,
                            navigate:navigate});
                    }else {
                        if(data){
                            const payments = data.map(
                                (paymentSingle: { ID: any; OSBBID: any; Amount: any; Appointment : any; CreatedAt: any; updatedAt: any}) => {
                                    const {ID, OSBBID, Amount, Appointment, CreatedAt, updatedAt} = paymentSingle;
                                    return {
                                        ID: ID,
                                        OSBBID: OSBBID,
                                        Amount: Amount,
                                        Appointment: Appointment,
                                        CreatedAt: CreatedAt,
                                        updatedAt: updatedAt
                                    }
                                });
                            setPayments(payments);
                        }else {
                            setPayments([]);
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, []);

    function updatePayment(paymentID:any, amount:any, appointment:any, setIsAppointmentChecked:any){
        let amountJSON = null
        let appointmentJSON = null
        if(amount !== "" ){
            amountJSON = {amount: amount};
        }
        if(appointment != null){
            appointmentJSON={appointment:appointment};
        }
        let body = {...amountJSON, ...appointmentJSON};
        const requestOptions = {
            method: 'PUT',
            headers:config.headers,
            body: JSON.stringify(body)
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/payments/'+paymentID, requestOptions)
            .then(response => response.json())
            .then(data => {
                if(data){
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:updatePayment, navigate:navigate});
                    }else {
                        setPayments((currentPayment:any) => {
                            return currentPayment.map((payment:any)=>{
                                if(payment.ID === paymentID){
                                    return {...payment, amount, appointment}
                                }
                                return payment
                            })
                        })
                        setIsAppointmentChecked(false);
                    }
                }
            });
    }

    function deletePayment(paymentID:any){
        const requestOptions = {
            method: 'DELETE',
            headers:config.headers,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/payments/'+paymentID, requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:deletePayment, navigate:navigate});
                }else {
                    setPayments((currentPayments: any) => {
                        return currentPayments.filter((payment:any) => payment.ID !== paymentID)
                    })
                }
            });
    }

    function addPayment (payment: any){
        // @ts-ignore
        setPayments(currentPayment => {
            return [
                ...currentPayment,
                payment,
            ]
        })
    }

    useEffect(() => {
        fetchPayments();
    }, [fetchPayments]);

    return(
        <section className='announcements-list'>
            <div className='container'>
                <div className={'flex flex-end m-5'}>
                    {!isAddedChecked && <MdAddCircle fontSize={'40px'} style={{color:'var(--blue-color)'}} onClick={()=>setIsAddedChecked(!isAddedChecked)}/>}
                    {isAddedChecked &&
                        <IoListCircleSharp fontSize={'40px'} style={{color:'var(--blue-color)'}} onClick={()=>setIsAddedChecked(!isAddedChecked)}/>}

                </div>
                {isAddedChecked &&  <PaymentForm addPayment={addPayment} setIsAddedChecked={setIsAddedChecked}/>}
                {!isAddedChecked &&  <div className='announcements-content grid'>
                    {payments.length === 0 && <h1 style={{color:"white"}}>Немає платежів</h1>}
                    {
                        payments.map((payment:{ID:any, Amount:any, Appointment:any})=> {
                            return (
                                <PaymentAdminItem key={payment.ID}
                                                       payment={payment}
                                                  deletePayment={deletePayment}
                                                  updatePayment={updatePayment}/>
                            )
                        })
                    }
                </div>}
            </div>
        </section>
    )
}

export default PaymentAdminList