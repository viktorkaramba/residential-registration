import React, {useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import Checkbox from "@mui/material/Checkbox";
import './Answer.css'
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import {Button} from "@mui/material";

const AnswerForm = ({pollID, userAnswer, isExist, setIsExist, updateAnswer, deleteAnswer}:any) =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [isChecked, setIsChecked] = useState(false);
    const [newAnswer, setNewAnswer] = useState(userAnswer?.content);
    const [show, setShow] = useState(false)

    function handleClickDelete(){
        setShow(true)
    }
    const handleClose = ()=>{
        setShow(false)
    }

   const addOpenAnswer = (event: any) => {
        event.preventDefault();
        // 👇️ access input values using name prop
        const content = event.target.answer?.value;
        const requestOptions = {
            method: 'POST',
            headers:{ 'Content-Type': 'application/json',
                'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            body: JSON.stringify({ content: content})
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/polls/'+ pollID + '/answers', requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:addOpenAnswer});
                }else {
                    if(data){
                        setIsExist(true);
                    }
                }
            });
        // 👇️ clear all input values in the form
        // event.target.reset();
    };
    useEffect(() => {
        if(userAnswer!=null){
            setNewAnswer(userAnswer?.content)
            setIsExist(true);
        }
    }, [userAnswer]);

    function handleDelete(){
        deleteAnswer(pollID);
        setNewAnswer('')
        setIsExist(false);
        setShow(false)
    }

    function handleUpdate(){
        updateAnswer(pollID, newAnswer, setIsChecked);
        setIsChecked(false);
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
                        <div className={'text-black fw-4 fs-20 dialog-style'}>Ви впевнені, що хочете видалити свою відповідь?</div>
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}><span className={'fw-4 fs-16 dialog-style'}>Відмінити</span></Button>
                    <Button onClick={handleDelete} autoFocus>
                        <span className={'fw-4 fs-16 dialog-style'}>Видалити</span>
                    </Button>
                </DialogActions>
            </Dialog>
        <form method='post'  onSubmit={addOpenAnswer}>
            {!isExist && <div className="inner-wrap">
                <label form={'answer'}>Відповідь
                    <input
                        maxLength={256}
                        minLength={2}
                        required={true}
                        name="answer"
                        placeholder=""
                        value={newAnswer}
                        type='text'
                        id='answer'
                        onChange={(event)=>{
                            setNewAnswer(event.target.value)
                        }}
                    />
                    <button className='button' style={{marginTop:'10px'}}>
                        <span className="button_content" >Відповісти</span>
                    </button>
                </label>
            </div>}
                    {isExist && !isChecked &&
                        <div className={'answer flex flex-sb'}>
                            <label>
                                <Checkbox
                                    name="announcement_check_box"
                                    id='announcement_check_box'
                                    checked={isChecked}
                                    size="medium"
                                    style={{color:'var(--blue-color)'}}
                                    onChange={e=>setIsChecked(!isChecked)}
                                />
                                {!isChecked && <span className={'m-5'}>{newAnswer}</span>}
                            </label>
                            <div className={'flex'}>
                                <button className='button'>
                                    <span className="button_content" onClick={()=>handleClickDelete()}>Видалити</span>
                                </button>
                            </div>
                        </div>
                      }
            {isExist && isChecked &&
                <div className={'answer flex'}>
                    <Checkbox
                        name="announcement_check_box"
                        id='announcement_check_box'
                        checked={isChecked}
                        size="medium"
                        style={{color:'var(--blue-color)'}}
                        onChange={e=>setIsChecked(!isChecked)}
                    />
                    <div style={{width:'100%'}}>
                        <input maxLength={256}
                               minLength={2}
                               required={true}
                               name="poll_answer_update_content"
                               placeholder=""
                               type='text'
                               onChange={e=>setNewAnswer(e.target.value)}
                               value={newAnswer}
                               className='m-5'
                               id='poll__answer_update_content'/>
                        <button className='button m-5' onClick={()=>handleUpdate()}>
                            <span className="button_content">Оновити</span>
                        </button>
                        <button className='button'>
                            <span className="button_content" onClick={()=>handleClickDelete()}>Видалити</span>
                        </button>
                    </div>
                </div>
            }
        </form>
        </div>
    )
}

export default AnswerForm