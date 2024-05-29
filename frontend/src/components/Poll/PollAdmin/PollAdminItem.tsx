import React, {useState} from "react";
import TestAnswerAdminList from "../TestAnswer/TestAnswerAdmin/TestAnswerAdminList";
import Checkbox from "@mui/material/Checkbox";
import {format} from "date-fns";
import '../Poll.css';
import {useAppContext} from "../../../utils/AppContext";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import {Button} from "@mui/material";


const PollAdminItem = ({poll, updatePoll, deletePoll}:any) => {
    const [isPollChecked, setIsPollChecked] = useState(false);
    const [newQuestion, setNewQuestion] = useState(poll.question);
    const [newFinished, setNewFinished] = useState(poll.finished_at);
    const [newIsClosed, setNewIsClosed] = useState(poll.is_closed);
    const [errorDate, setErrorDate]= useState(false);
    // @ts-ignore
    const {setPoll, setActivePollElement} = useAppContext();
    const [show, setShow] = useState(false)

    function handleClickDelete(){
        setShow(true)
    }
    const handleClose = ()=>{
        setShow(false)
    }

    function handleDelete(){
        deletePoll(poll.id);
        setShow(false)
    }

    function handleUpdate(){
        console.log(newFinished, poll.finished_at)
        if(newFinished!==poll.finished_at){
            if (new Date(newFinished) < new Date()){
                setErrorDate(true);
                return
            }
            updatePoll(poll.id, newQuestion, newFinished, newIsClosed, setIsPollChecked)
        }
        else {
            updatePoll(poll.id, newQuestion, null, newIsClosed, setIsPollChecked)
        }
    }

    function toResult(){
        setActivePollElement('PollResultItem')
        setPoll(poll)
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
                        <div className={'text-black fw-4 fs-20 dialog-style'}>Ви впевнені, що хочете видалити опитування?</div>
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
                <>
                    <Checkbox
                        name="announcement_check_box"
                        id='announcement_check_box'
                        checked={isPollChecked}
                        size="large"
                        style={{color:'var(--blue-color)'}}
                        onChange={()=>{setIsPollChecked(!isPollChecked)}}
                    />
                </>
            </div>
            <div className='poll-item'>
                {!isPollChecked &&
                    <div className="inner-wrap">
                        <div className='flex flex-sb flex-wrap'>
                            <div className='polls-item-info-item fw-7 fs-26'>
                                <span>{newQuestion} | {newIsClosed && <span>Закрите</span>}{!newIsClosed && <span>Відкрите</span>}</span>
                            </div>
                            <div className='polls-item-info-item fw-6 fs-15'>
                                <span> {format(poll.created_at, 'MMMM do yyyy, hh:mm:ss a')}</span>
                            </div>
                        </div>
                        <div className='flex flex-sb flex-wrap'>
                            <div className='polls-item-info-item fw-7 fs-15'>
                                <span>Завершується - {format(poll.finished_at, 'MMMM do yyyy, hh:mm:ss a')}</span>
                            </div>
                        </div>
                    </div>}
                {isPollChecked &&
                    <div className="inner-wrap">
                        <div className='flex flex-sb flex-wrap'>
                            <label className='polls-item-info-item fw-7 fs-20' form={'question'}>Нове запитання
                                <input maxLength={256}
                                       minLength={2}
                                       required={true}
                                       name="poll_update_question"
                                       placeholder=""
                                       type='text'
                                       onChange={e=>setNewQuestion(e.target.value)}
                                       value={newQuestion}
                                       id='poll_update_question'/>
                            </label>
                        </div>
                        <div className={'polls-item-info-item fw-7 fs-20'}>
                            <label form={'finished_at'}>Новий час завершення
                                <input name="poll_update_finished_at"
                                       placeholder=""
                                       value={newFinished}
                                       onChange={e=>setNewFinished(e.target.value)}
                                       type='datetime-local'
                                       step="1"
                                       id='poll_update_finished_at'/>
                            </label>
                        </div>
                        <div className={'polls-item-info-item fw-7 fs-20'}>
                            <label form={'is_closed'}>
                                <input name="poll_update_is_closed"
                                       checked={newIsClosed}
                                       onChange={e=>setNewIsClosed(!newIsClosed)}
                                       type='checkbox'
                                       className={'m-5'}
                                       id='poll_update_is_closed'/>
                                Завершити
                            </label>
                        </div>
                        {errorDate &&
                            <div className={'error'}>
                                Дата завершення опитування повинна бути більша за поточну
                            </div>
                        }
                    </div>
                }
                {isPollChecked &&
                    <button className='button poll_button_form' style={{marginRight:'5px'}} type="submit" onClick={()=>handleUpdate()} name="update_annpuncement">
                        <span className="button_content poll_button_content_form">Оновити</span>
                    </button>
                }
                {poll.test_answer.length !== 0 &&  <TestAnswerAdminList answers={poll.test_answer} pollID={poll.id}/>}
                <button className='button poll_button_form' style={{marginRight:'5px'}} type="submit" onClick={()=>toResult()} name="to_result">
                    <span className="button_content poll_button_content_form">Результати</span>
                </button>
                <button className='button announcement_button' type="submit" onClick={()=>handleClickDelete()} name="update_annpuncement">
                    <span className="button_content announcement_button_content">Видалити</span>
                </button>
            </div>
        </div>

    )
}

export default PollAdminItem