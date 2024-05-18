import React, {useState} from "react";
import Checkbox from "@mui/material/Checkbox";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import {Button} from "@mui/material";

const TestAnswerAdminItem = ({content, id, count_test_answers, updateTestAnswer, deleteTestAnswer}:any) =>{
    const [isChecked, setIsChecked] = useState(false);
    const [newContent, setNewContent] = useState(content);
    const [errorDelete, setErrorDelete] = useState(false);
    const [show, setShow] = useState(false)
    const [pollID, setPollID] = useState(null)

    function handleClickDelete(id:any){
        setShow(true)
        setPollID(id)
    }
    const handleClose = ()=>{
        setShow(false)
    }

    function handleDelete(){
        if(count_test_answers<3){
            setErrorDelete(true)
        }else {
            setErrorDelete(false)
            deleteTestAnswer(pollID);
        }
    }

    return(
        <li className={'flex flex-sb'}>
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
                        <div className={'text-black fw-4 fs-20 dialog-style'}>Ви впевнені, що хочете видалити тестову відповідь?</div>
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}><span className={'fw-4 fs-16 dialog-style'}>Відмінити</span></Button>
                    <Button onClick={handleDelete} autoFocus>
                        <span className={'fw-4 fs-16 dialog-style'}>Видалити</span>
                    </Button>
                </DialogActions>
            </Dialog>
            <label  className={'flex'}>
                <Checkbox
                    name="announcement_check_box"
                    id='announcement_check_box'
                    checked={isChecked}
                    size="medium"
                    style={{color:'var(--blue-color)'}}
                    onChange={e=>setIsChecked(!isChecked)}
                />
                {!isChecked && <span className={'m-5'}>{newContent}</span>}
                {isChecked &&
                    <div className={'poll_update'}>
                        <input maxLength={256}
                               minLength={2}
                               required={true}
                               name="test_answer_update_content"
                               placeholder=""
                               type='text'
                               onChange={e=>setNewContent(e.target.value)}
                               value={newContent}
                               className='m-5'
                               id='test_answer_update_content'/>
                            <button className='button m-5' type="submit" name="delete_test_answer"   onClick={()=>updateTestAnswer(id, newContent, setIsChecked)}>
                                <span className="button_content">Оновити</span>
                            </button>
                    </div>
                }
            </label>
            {errorDelete &&
                <div className={'error login_error'}>
                    Кількість відповідей повинна бути більше ніж 2
                </div>
            }
            <button className='button' type="submit" name="delete_test_answer"   onClick={()=>handleClickDelete(id)}>
                <span className="button_content">Видалити</span>
            </button>
        </li>
    )
}

export default TestAnswerAdminItem