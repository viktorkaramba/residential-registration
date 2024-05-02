import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import err from "../../../utils/err";
import {useAppContext} from "../../../utils/AppContext";
import {format} from "date-fns";

const PaymentUserItem = ({payment, updatePurchase}:any) => {
    // @ts-ignore
    const {osbbID, user} = useAppContext()

    function updatePurchaseRequest(purchaseID:any, payment_status:any){

        const requestOptions = {
            method: 'PUT',
            headers:config.headers,
            body: JSON.stringify({ payment_status: payment_status})
        }
        let isOsbbHead="";
        if(user.role === 'osbb_head'){
            isOsbbHead="-osbb-head";
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/payments/'+payment.payment_id+'/purchase'  , requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data)
                if(data){
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:updatePurchase});
                    }else {
                        payment.payment_status = payment_status
                        updatePurchase(payment)
                    }
                }
            });
    }

    return(
        <div className='poll-item'>
            <div className="inner-wrap">
                <div className='flex flex-sb flex-wrap'>
                    <div className='announcements-item-info-item fw-7 fs-26'>
                        <span> {payment.amount}UAH</span>
                    </div>
                    <div className='announcements-item-info-item fw-6 fs-15'>
                        <span> {format(payment.created_at, 'MMMM do yyyy, hh:mm:ss a')}</span>
                    </div>
                </div>
                <div className={'announcements-item-info'}>
                    <div className='announcements-item-info-item fw-6 fs-18'>
                        <span>{payment.appointment}</span>
                    </div>
                </div>
                </div>
            {payment.payment_status==="not_paid" && user?.id === payment.user_id && <div>
                <button className='button announcement_button' style={{marginRight:'5px'}} type="submit" onClick={()=>updatePurchaseRequest(payment.purchase_id, "paid")} name="to_result">
                    <span className="button_content poll_button_content_form">Оплатити</span>
                </button>
            </div>}
            {payment.payment_status==="not_paid" && user?.id !== payment.user_id && <div>
                <div className='announcements-item-info-item fw-7 fs-26'>
                    <span>Не сплачено</span>
                </div>
            </div>}
            {payment.payment_status==="paid" && <div>
                <button className='button announcement_button' style={{marginRight:'5px'}} type="submit" onClick={()=>updatePurchaseRequest(payment.purchase_id, "not_paid")}name="delete_purchase">
                    <span className="button_content poll_button_content_form">Видалити</span>
                </button>
            </div>}
        </div>
    )
}

export default PaymentUserItem