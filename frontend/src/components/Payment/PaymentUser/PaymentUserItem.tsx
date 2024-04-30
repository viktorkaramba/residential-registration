import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import err from "../../../utils/err";
import {useAppContext} from "../../../utils/AppContext";
import {useNavigate} from "react-router-dom";
import {format} from "date-fns";

const PaymentUserItem = ({payment}:any) => {
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [userPurchase, setUserPurchase] = useState<any>(null);
    const navigate = useNavigate();
    const [isExist, setIsExist] = useState(false);

    const fetchUserPurchase = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: { 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            }
            fetch(config.apiUrl+'osbb/'+ osbbID+ '/purchases-user', requestOptions)
                .then(response => response.json())
                .then(data => {
                    if(data){
                        const {error}:any = data;
                        if(error){
                            err.HandleError({errorMsg:error, func:fetchUserPurchase, navigate:navigate});
                        }else {
                            const { id, osbb_id, payment_id, user_id, payment_status, createdAt}:any =data
                            const newPurchase ={
                                id: id,
                                osbb_id: osbb_id,
                                payment_id: payment_id,
                                user_id: user_id,
                                payment_status: payment_status,
                                createdAt: createdAt,
                            }
                            console.log(newPurchase)
                            setUserPurchase(newPurchase)
                        }
                    }else {
                        setUserPurchase(null)
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, []);

    function updatePurchase(purchaseID:any, payment_status:any, setIsChecked:any){
        const requestOptions = {
            method: 'PUT',
            headers:config.headers,
            body: JSON.stringify({ payment_status: payment_status})
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/purchase-user', requestOptions)
            .then(response => response.json())
            .then(data => {
                if(data){
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:userPurchase});
                    }else {
                        setIsChecked(false);
                    }
                }

            });
    }

    function deletePurchase(purchaseID:any){
        const requestOptions = {
            method: 'DELETE',
            headers:config.headers,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/purchase-user/'+purchaseID, requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:deletePurchase});
                }
            });
    }

    useEffect(() => {
        fetchUserPurchase();
    }, []);

    useEffect(() => {
        if(userPurchase != null){
            setIsExist(userPurchase.content.length!==0)
        }
    }, [userPurchase]);

    return(
        <div className='poll-item'>
            <div className="inner-wrap">
                <div className='flex flex-sb flex-wrap'>
                    <div className='polls-item-info-item fw-7 fs-26'>
                        <span>{payment.amount}</span>
                    </div>
                    <div className='polls-item-info-item fw-6 fs-15'>
                        <span> {format(payment.created_at, 'MMMM do yyyy, hh:mm:ss a')}</span>
                    </div>
                </div>
                {/*<div className='flex flex-sb flex-wrap'>*/}
                {/*    <div className='polls-item-info-item fw-7 fs-15'>*/}
                {/*        <span>Завершується - {format(poll.finished_at, 'MMMM do yyyy, hh:mm:ss a')}</span>*/}
                {/*    </div>*/}
                {/*</div>*/}
            </div>
        </div>
    )
}

export default PaymentUserItem