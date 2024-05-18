import React, {useEffect, useState} from "react";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import {Button} from "@mui/material";

const TestAnswerUserItem = ({content, id, pollID, userAnswer, addTestAnswer, selectedValue, setSelectedValue, deleteAnswer}:any) =>{
    const handleChange = (e:any) => {
        setSelectedValue(e);
        addTestAnswer(e);
    };
    const [show, setShow] = useState(false)

    function handleClickDelete(){
        setShow(true)
    }
    const handleClose = ()=>{
        setShow(false)
    }

    function handleDelete(){
        deleteAnswer(pollID)
        setSelectedValue(undefined)
        setShow(false)
    }

    useEffect(() => {
        if(userAnswer != null){
            if(userAnswer.test_answer_id === id){
                setSelectedValue(id)
            }else {
            }
        }
    }, [userAnswer]);

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
                        <div className={'text-black fw-4 fs-20 dialog-style'}>Ви впевнені, що хочете видалити відповідь?</div>
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}><span className={'fw-4 fs-16 dialog-style'}>Відмінити</span></Button>
                    <Button onClick={handleDelete} autoFocus>
                        <span className={'fw-4 fs-16 dialog-style'}>Видалити</span>
                    </Button>
                </DialogActions>
            </Dialog>
            <li className={'flex flex-sb'}>
                <label  className={'flex'}>
                    <input
                        type="radio"
                        checked={selectedValue==id}
                        value={id}
                        onChange={e=>handleChange(e.target.value)}
                    />
                    <span className={'m-5'}>{content}</span>
                </label>
                {selectedValue ==id &&  <button className='button' type="submit" name="delete_test_answer"   onClick={()=>handleClickDelete()}>
                    <span className="button_content">Видалити</span>
                </button>}
            </li>
        </div>

    )
}

export default TestAnswerUserItem