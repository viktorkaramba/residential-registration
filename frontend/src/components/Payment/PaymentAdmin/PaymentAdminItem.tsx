import React, {useEffect, useState} from "react";
import {format} from "date-fns";
import Checkbox from '@mui/material/Checkbox';
import PaymentForm from "../PaymentForm/PaymentForm";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import {Button} from "@mui/material";

const PaymentAdminItem = ({payment, deletePayment, updatePayment, addPayment}:any) => {
    const [newAmount, setNewAmount] = useState(payment.Amount);
    const [newAppointment, setNewAppointment] = useState(payment.Appointment);
    const [isPaymentChecked, setIsPaymentChecked] = useState(false);
    const [isAddedChecked, setIsAddedChecked] = useState(false);
    const [show, setShow] = useState(false)

    function handleClickDelete(){
        setShow(true)
    }
    const handleClose = ()=>{
        setShow(false)
    }

    function handleDelete(){
        deletePayment(payment.ID);
    }

    function handleUpdate(){
        updatePayment(payment.ID, newAmount, newAppointment, setIsPaymentChecked)
    }

    return(
        <div>
            <Dialog
                open={show}
                onClose={handleClose}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                <DialogTitle id="alert-dialog-title">
                    <div className={'text-black fw-7 fs-24 dialog-style'}>{"Попередження"}</div>
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        <div className={'text-black fw-4 fs-20 dialog-style'}>Ви впевнені, що хочете видалити послугу?</div>
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}><span className={'fw-4 fs-16 dialog-style'}>Відмінити</span></Button>
                    <Button onClick={handleDelete} autoFocus>
                        <span className={'fw-4 fs-16 dialog-style'}>Видалити</span>
                    </Button>
                </DialogActions>
            </Dialog>
            <div className={'flex flex-end m-5'}>
                {!isAddedChecked &&
                    <>
                        <Checkbox
                            name="payment_check_box"
                            id='payment_check_box'
                            checked={isPaymentChecked}
                            size="large"
                            style={{color:'var(--blue-color)'}}
                            onChange={()=>{setIsPaymentChecked(!isPaymentChecked)}}
                        />
                    </>}
            </div>
            {isAddedChecked &&  <PaymentForm addPayment={addPayment} setIsAddedChecked={setIsAddedChecked}/>}
            {!isAddedChecked &&
                <div className='announcements-item'>
                    {!isPaymentChecked &&
                        <div className="inner-wrap">
                            <div className='flex flex-sb flex-wrap'>
                                <div className='announcements-item-info-item fw-7 fs-26'>
                                    <span> {newAmount} UAH</span>
                                </div>
                                <div className='announcements-item-info-item fw-6 fs-15'>
                                    <span> {format(payment.CreatedAt, 'MMMM do yyyy, hh:mm:ss a')}</span>
                                </div>
                            </div>
                            <div className={'announcements-item-info'}>
                                <div className='announcements-item-info-item fw-6 fs-18'>
                                    <span>{newAppointment}</span>
                                </div>
                            </div>
                        </div>}
                    {isPaymentChecked &&
                        <div className="inner-wrap">
                            <div className='flex flex-sb flex-wrap'>
                                <label className='announcements-item-info-item fw-7 fs-26' form={'amount'}>Нова ціна сплати
                                    <input name="payment_update_amount"
                                           placeholder=""
                                           type='number'
                                           onChange={e=>setNewAmount(e.target.value)}
                                           value={newAmount}
                                           id='payment_update_amount'/>
                                </label>
                            </div>
                            <div className={'announcements-item-info'}>
                                <label form={'appointment'}>Нове признаення платежу
                                    <textarea minLength={2} onChange={e=>{setNewAppointment(e.target.value)}} required={true} disabled={false} name="appoinment" placeholder="" id='appoinment' value={newAppointment}/>
                                </label>
                            </div>
                        </div>
                    }
                    {isPaymentChecked &&
                        <button className='button announcement_button' type="submit" onClick={()=>handleUpdate()} name="update_payment">
                            <span className="button_content announcement_button_content">Оновити</span>
                        </button>
                    }
                    <button className='button announcement_button' type="submit" onClick={()=>handleClickDelete()} name="delete_payment">
                        <span className="button_content announcement_button_content">Видалити</span>
                    </button>
                </div>
            }
        </div>
    )
}

export default PaymentAdminItem