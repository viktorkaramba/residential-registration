import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import PaymentUserItem from "./PaymentUserItem";
import Select from 'react-select'


const PaymentUserList = () =>{
    // @ts-ignore
    const {osbbID, user} = useAppContext()


    const options = [
        { value: 'all', label: 'Усі' },
        { value: 'not_paid', label: 'Не оплачені' },
        { value: 'paid', label: 'Оплачені' }
    ]
    const [userChoice, setUserChoice] = useState<any>(0)
    const [purchaseChoice, setPurchaseChoice] = useState('all')
    const [paymentChoice, setPaymentChoice] = useState(0)
    const [payments, setPayments] = useState([]);
    const [allUsers, setAllUsers] = useState([]);
    const optionsUser = allUsers.map((user:any) => ({
        label: user.full_name.first_name + " " + user.full_name.surname + " " + user.full_name.patronymic,
        value: user.id
    }));
    const navigate = useNavigate();

    const fetchAllUsers = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: { 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            }
            fetch(config.apiUrl+'osbb/'+ osbbID+ '/inhabitants', requestOptions)
                .then(response => response.json())
                .then(data=>{
                    console.log(data)
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchAllUsers,
                            navigate:navigate});
                    }else {
                        if(data){
                            const inhabitants = data.map(
                                (inhabitantSingle: { id: any; osbbid: any; apartment: any; full_name: any; photo: any; phone_number: any; role: any; is_approved:any }) => {
                                    const {id, osbbid, apartment, full_name, photo, phone_number, role, is_approved} = inhabitantSingle;
                                    return {
                                        id: id,
                                        osbbid: osbbid,
                                        apartment: apartment,
                                        full_name: full_name,
                                        photo: photo,
                                        phone_number: phone_number,
                                        role: role,
                                        is_approved:is_approved
                                    }
                                });
                            setAllUsers(inhabitants);
                        }else {
                            setAllUsers([]);
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, []);

    function handleFilter(){
        let userJSON = null
        let purchaseJSON = null
        let paymentJSON = null
        if(purchaseChoice !== "all" ){
            purchaseJSON={payment_status:purchaseChoice};
        }
        if(userChoice !== 0 ){
            userJSON = {user_id: userChoice};
        }
        if(paymentChoice !== 0) {
            paymentJSON = {payment_id:paymentChoice};
        }
        let body = {...userJSON, ...paymentJSON, ...purchaseJSON};
        console.log(body);
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json',
                'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            body: JSON.stringify(body)
        }
        let isOsbbHead="user";
        if(user.role === 'osbb_head'){
            isOsbbHead="osbb-head";
        }
        fetch(config.apiUrl+'osbb/'+ osbbID+ '/purchase-'+isOsbbHead, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if(data == null){
                    setPayments([])
                    return
                }
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:handleFilter, navigate:navigate});
                }else {
                    if(data){
                        const payments = data.map(
                            (paymentSingle: { payment_id:any, purchase_id: any; user_id: any;amount: any; appointment: any; payment_status: any; created_at: any; updated_at:any }) => {
                                const {payment_id, purchase_id, user_id, amount, appointment, payment_status, created_at,updated_at } = paymentSingle;
                                return {
                                    payment_id: payment_id,
                                    purchase_id: purchase_id,
                                    user_id: user_id,
                                    amount: amount,
                                    appointment: appointment,
                                    payment_status: payment_status,
                                    created_at: created_at,
                                    updated_at: updated_at
                                }
                            });
                        setPayments(payments)
                    }else {
                        setPayments([])
                    }
                }
            });
    }

    function updatePurchase (newPurchase: any){
        if(userChoice === "all"){
            setPayments((currentPurchase:any) => {
                return currentPurchase.map((purchase:any)=>{
                    if(purchase.purchase_id === newPurchase.purchase_id){
                        return {...purchase}
                    }
                    return purchase
                })
            })
        }else {
            // @ts-ignore
            setPayments((currentPurchase: any) => {
                return currentPurchase.filter((purchase:any) => purchase.purchase_id !== newPurchase.purchase_id)
            })
        }
    }

    useEffect(() => {
        handleFilter();
    }, [userChoice, purchaseChoice]);

    useEffect(() => {
        fetchAllUsers();
    }, [fetchAllUsers]);
    return(
        <section className='poll_list'>
            <div className='container'>
                <div className='poll_content grid'>
                    <div className={'flex'} >
                        <div className={'m-5'} style={{flexGrow:2}}>

                        </div>
                        {user.role === "osbb_head" && <>
                            <div className={'m-5'} style={{flexGrow:6}}>

                            </div>
                            <div className={'m-5'} style={{flexGrow:2, width: '100px'}} >
                                <Select options={optionsUser} defaultValue={optionsUser[0]}
                                        onChange={(choice:any) => setUserChoice(choice.value)}/>
                            </div>
                            <div className={'m-5'} style={{flexGrow:2, width: '100px'}} >
                                <Select options={options} defaultValue={options[0]}
                                        onChange={(choice:any) => setPurchaseChoice(choice)}/>
                            </div>
                        </>}
                        {user.role !== "osbb_head" && <>
                            <div className={'m-5'} style={{flexGrow:8}}>

                            </div>
                            <div className={'m-5'} style={{flexGrow:2, width: '100px'}} >
                                <Select options={options} defaultValue={options[0]}
                                        onChange={(choice:any) => setPurchaseChoice(choice)}/>
                            </div>
                        </>}
                    </div>
                    {payments.length === 0 && <h1 style={{color:"white"}}>Немає платежів</h1>}
                    {
                        payments.map((payment:{payment_id:any, purchase_id: any; amount: any; appointment: any; payment_status: any; created_at: any; updated_at:any }) => {
                            return (
                                <PaymentUserItem
                                    payment={payment}
                                    key={payment.purchase_id}
                                    updatePurchase={updatePurchase}
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