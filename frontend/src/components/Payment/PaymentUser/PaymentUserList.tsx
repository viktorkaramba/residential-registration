import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import PaymentUserItem from "./PaymentUserItem";
import Select from 'react-select'
import "../Payment.css"
import { RiSortAsc, RiSortDesc } from "react-icons/ri";

const PaymentUserList = () =>{
    // @ts-ignore
    const {osbbID, user} = useAppContext();
    const options = [
        { value: 'all', label: 'Усі' },
        { value: 'not_paid', label: 'Не оплачені' },
        { value: 'paid', label: 'Оплачені' }
    ];
    const [userChoice, setUserChoice] = useState<any>(0);
    const [purchaseChoice, setPurchaseChoice] = useState('all');
    const [paymentChoice, setPaymentChoice] = useState(0);
    const [dateFromChoice, setDataFromChoice] = useState<any>(null);
    const [dateToChoice, setDateToChoice] = useState<any>(null);
    const [filterPayments, setFilterPayments] = useState<any>([]);
    const [isAsc, setIsAsc] = useState<any>(true);
    const [payments, setPayments] = useState<any>([]);
    const [allUsers, setAllUsers] = useState([]);
    const optionsUser = allUsers.map((user:any) => ({
        label: user.full_name.first_name + " " + user.full_name.surname + " " + user.full_name.patronymic,
        value: user.id
    }));
    const indexToMove = optionsUser.findIndex(currentUser => {
        return currentUser.value === user.id; // Припустимо, ваша умова - це співпадіння value з якимось значенням
    });
    if (indexToMove !== -1) {
        const elementToMove = optionsUser.splice(indexToMove, 1)[0]; // Видаляємо елемент із масиву
        optionsUser.unshift(elementToMove);
    }
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
            userJSON = {user_id: parseInt(userChoice)};
        }
        if(paymentChoice !== 0) {
            paymentJSON = {payment_id:paymentChoice};
        }
        let body = {...userJSON, ...paymentJSON, ...purchaseJSON};
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
                if(data == null){
                    setPayments([])
                    setFilterPayments([])
                    return
                }
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:handleFilter, navigate:navigate});
                }else {
                    if(data){
                        const newPayments = data.map(
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

                        console.log(isAsc);
                        if(isAsc){
                            setPayments(newPayments.sort(compareCreatedAtAsc));
                            setFilterPayments(payments);
                        }else {
                            setPayments(newPayments.sort(compareCreatedAtDesc));
                            setFilterPayments(payments);
                        }
                    }else {
                        setPayments([]);
                        setFilterPayments([]);
                    }
                }
            });
    }

    function handleDateFilter(){
        if((dateFromChoice == null || dateFromChoice==='') && dateToChoice!=null && dateToChoice!==''){
            const toDateObj = new Date(dateToChoice);
            return payments.filter((payment: any) => {
                const updatedAtDate = new Date(payment.created_at);
                return updatedAtDate <= toDateObj;
            });
        }
        if((dateFromChoice!=null && dateFromChoice!=='') && (dateToChoice==null || dateToChoice==='')){
            const fromDateObj = new Date(dateFromChoice);
            return payments.filter((payment: any) => {
                const updatedAtDate = new Date(payment.created_at);
                return updatedAtDate >= fromDateObj;
            });
        }
        if(dateFromChoice!=null && dateFromChoice!=='' && dateToChoice!=null && dateToChoice!==''){
            const fromDateObj = new Date(dateFromChoice);
            const toDateObj = new Date(dateToChoice);
            return payments.filter((payment: any) => {
                const updatedAtDate = new Date(payment.created_at);
                return updatedAtDate >= fromDateObj && updatedAtDate <= toDateObj;
            });
        }
        if((dateFromChoice==null && dateToChoice==null) || (dateFromChoice==='' && (dateToChoice==='' || dateToChoice==null))
        || ((dateFromChoice===''|| dateFromChoice==null) && dateToChoice===null)||
            (dateFromChoice==null && (dateToChoice==='' || dateToChoice==null))){
            return payments;
        }
    }

    function updatePurchase (newPurchase: any){
        if(purchaseChoice === "all"){
            setPayments((currentPurchase:any) => {
                return currentPurchase.map((purchase:any)=>{
                    if(purchase.purchase_id === newPurchase.purchase_id){
                        return {...purchase}
                    }
                    return purchase
                })
            })
            setFilterPayments((currentPurchase:any) => {
                return currentPurchase.map((purchase:any)=>{
                    if(purchase.purchase_id === newPurchase.purchase_id){
                        return {...purchase}
                    }
                    return purchase
                })
            })
        }else {
            setPayments((currentPurchase: any) => {
                return currentPurchase.filter((purchase:any) => purchase.purchase_id !== newPurchase.purchase_id)
            })
            setFilterPayments((currentPurchase: any) => {
                return currentPurchase.filter((purchase:any) => purchase.purchase_id !== newPurchase.purchase_id)
            })
        }
    }

    function compareCreatedAtDesc(a:any, b:any) {
        const dateA = new Date(a.created_at);
        const dateB = new Date(b.created_at);

        if (dateA > dateB) {
            return -1;
        }
        if (dateA < dateB) {
            return 1;
        }
        return 0;
    }

    function compareCreatedAtAsc(a:any, b:any) {
        const dateA = new Date(a.created_at);
        const dateB = new Date(b.created_at);

        if (dateA < dateB) {
            return -1;
        }
        if (dateA > dateB) {
            return 1;
        }
        return 0;
    }

    useEffect(() => {
        handleFilter();
    }, [userChoice, purchaseChoice]);

    useEffect(() => {
        if(isAsc){
            setFilterPayments(handleDateFilter().sort(compareCreatedAtAsc));
        }else {
            setFilterPayments(handleDateFilter().sort(compareCreatedAtDesc));
        }
    }, [dateFromChoice, dateToChoice, handleFilter]);

    useEffect(() => {
        fetchAllUsers();
    }, [fetchAllUsers]);

    useEffect(() => {
        if(!isAsc){
            setFilterPayments(filterPayments.sort(compareCreatedAtAsc));
        }else {
            setFilterPayments(filterPayments.sort(compareCreatedAtDesc));
        }

    }, [filterPayments, isAsc]);

    return(
        <section className='poll_list'>
            <div className='container'>
                <div className='poll_content grid'>
                    <div className={'flex'} >
                        <div style={{flexGrow:1}}>
                            {isAsc && <RiSortAsc size={"30px"} style={{color:"var(--blue-color)"}} onClick={()=>setIsAsc(!isAsc)}/>}
                            {!isAsc && <RiSortDesc size={"30px"} style={{color:"var(--blue-color)"}} onClick={()=>setIsAsc(!isAsc)}/>}
                        </div>
                        {user.role === "osbb_head" && <>
                            <div className={'payment-date-item m-5'} style={{flexGrow:1}}>
                                <label form={'date_from'}>Від
                                    <input className={'m-5'} required={true} onChange={(event:any)=>setDataFromChoice(event.target.value)} name="date_from" placeholder="" type='datetime-local' step="1" id='date_from'/>
                                </label>
                            </div>
                            <div className={'payment-date-item m-5'} style={{flexGrow:1}}>
                                <label form={'finished_at'}>До
                                    <input className={'m-5'} required={true} onChange={(event:any)=>setDateToChoice(event.target.value)} name="finished_at" placeholder="" type='datetime-local' step="1" id='finished_at'/>
                                </label>
                            </div>
                            <div className={'m-5'} style={{flexGrow:4, width: '250px'}} >
                                <Select options={optionsUser}
                                        defaultValue={{
                                            label: user.full_name.first_name + " " + user.full_name.surname + " " + user.full_name.patronymic,
                                            value: user.id }}
                                        onChange={(choice:any) => setUserChoice(choice.value)}/>
                            </div>
                            <div className={'m-5'} style={{flexGrow:3, width: '150px'}} >
                                <Select options={options} defaultValue={options[0]}
                                        onChange={(choice:any) => setPurchaseChoice(choice.value)}/>
                            </div>
                        </>}
                        {user.role !== "osbb_head" && <>
                            <div className={'m-5'} style={{flexGrow:5}}>
                                <label form={'payment-date-item finished_at'}>Дата завершення
                                    <input required={true} name="finished_at" onChange={(event:any)=>setDataFromChoice(event.target.value)} placeholder="" type='datetime-local' step="1" id='finished_at'/>
                                </label>
                            </div>
                            <div className={'payment-date-item m-5'} style={{flexGrow:5}}>
                                <label form={'finished_at'}>Дата завершення
                                    <input required={true} name="finished_at" onChange={(event:any)=>setDateToChoice(event.target.value)} placeholder="" type='datetime-local' step="1" id='finished_at'/>
                                </label>
                            </div>
                            <div className={'m-5'} style={{flexGrow:2, width: '100px'}} >
                                <Select options={options} defaultValue={options[0]}
                                        onChange={(choice:any) => setPurchaseChoice(choice.value)}/>
                            </div>
                        </>}
                    </div>
                    {filterPayments?.length === 0 && <h1 style={{color:"white"}}>Немає платежів</h1>}
                    {
                        filterPayments?.map((payment:{payment_id:any, purchase_id: any; amount: any; appointment: any; payment_status: any; created_at: any; updated_at:any }) => {
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