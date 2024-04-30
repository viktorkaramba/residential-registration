import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import PaymentUserItem from "./PaymentUserItem";
import Select from 'react-select'


const PaymentUserList = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const options = [
        { value: 'all', label: 'Усі' },
        { value: 'not_paid', label: 'Не оплачені' },
        { value: 'paid', label: 'Оплачені' }
    ]
    const [userChoice, setUserChoice] = useState('all')
    const [payments, setPayments] = useState([]);
    const navigate = useNavigate();
    const fetchPayments = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: { 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            }
            fetch(config.apiUrl+'osbb/'+ osbbID+ '/payments', requestOptions)
                .then(response => response.json())
                .then(data => {
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchPayments, navigate:navigate});
                    }else {
                        if(data){
                            const payments = data.map(
                                (paymentSingle: { id:any, amount: any; appointment: any; createdAt: any; updatedAt: any }) => {
                                    const {id, amount, appointment, createdAt, updatedAt} = paymentSingle;
                                    return {
                                        id: id,
                                        amount: amount,
                                        appointment: appointment,
                                        createdAt: createdAt,
                                        updatedAt: updatedAt
                                    }
                                });
                            setPayments(payments)
                        }else {
                            setPayments([])
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, []);

    function handleFilter(choice:any){

        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json',
                'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            body: JSON.stringify({ payment_status: choice})
        }
        fetch(config.apiUrl+'osbb/'+ osbbID+ '/purchase-user', requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:fetchPayments, navigate:navigate});
                }else {
                    if(data){
                        const payments = data.map(
                            (paymentSingle: { id:any, amount: any; appointment: any; createdAt: any; updatedAt: any }) => {
                                const {id, amount, appointment, createdAt, updatedAt} = paymentSingle;
                                return {
                                    id: id,
                                    amount: amount,
                                    appointment: appointment,
                                    createdAt: createdAt,
                                    updatedAt: updatedAt
                                }
                            });
                        setPayments(payments)
                    }else {
                        setPayments([])
                    }
                }
            });
    }
    useEffect(() => {
        fetchPayments();
    }, [userChoice]);

    // @ts-ignore
    // @ts-ignore
    return(
        <section className='poll_list'>
            <div className='container'>
                <div className='poll_content grid'>
                    {payments.length === 0 && <h1 style={{color:"white"}}>Немає платежів</h1>}
                    {payments.length === 0 &&
                        <div className={'flex'} >
                            <div style={{flexGrow:2}}>

                            </div>
                            <div style={{flexGrow:8}}>

                            </div>
                            <div style={{flexGrow:2, width: '100px'}} >
                                <Select options={options} defaultValue={options[0]}
                                        onChange={(choice:any) => handleFilter(choice.value)}/>
                                <p>{userChoice}</p>
                            </div>

                        </div>

                    }
                    {
                        payments.map((payment:{id:any, amount:any, appointment:any}) => {
                            return (
                                <PaymentUserItem
                                    payment={payment}
                                    key={payment.id}
                                />
                            )
                        })
                    }
                </div>
            </div>
        </section>
    )
}

export default PaymentUserList