import React, {useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import Checkbox from "@mui/material/Checkbox";
import './Answer.css'

const AnswerForm = ({pollID, userAnswer, isExist, setIsExist, updateAnswer, deleteAnswer}:any) =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [isChecked, setIsChecked] = useState(false);
    const [newAnswer, setNewAnswer] = useState(userAnswer !== null ? userAnswer.content : '');
   const addOpenAnswer = (event: any) => {
        event.preventDefault();
        // 👇️ access input values using name prop
        const content = event.target.answer.value;
        const requestOptions = {
            method: 'POST',
            headers:config.headers,
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
            setNewAnswer(userAnswer.content)
            setIsExist(true);
        }
    }, [userAnswer]);

    function handleDelete(){
        deleteAnswer(pollID);
        setNewAnswer('')
        setIsExist(false);
    }

    function handleUpdate(){
        updateAnswer(pollID, newAnswer, setIsChecked);
        setIsChecked(false);
    }


    return(
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
                                    <span className="button_content" onClick={()=>handleDelete()}>Видалити</span>
                                </button>
                            </div>
                        </div>
                      }
            {isExist && isChecked &&
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
                        <div className={'poll_answer_update'}>
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
                        </div>
                    </label>
                    <div className={'flex'}>
                        <button className='button'>
                            <span className="button_content" onClick={()=>handleDelete()}>Видалити</span>
                        </button>
                    </div>
                </div>
            }
        </form>
    )
}

export default AnswerForm