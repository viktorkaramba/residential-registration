import React, {useEffect, useState} from "react";
import config from "../../../config";
import {useAppContext} from "../../../AppContext";
import err from "../../../err";

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
        <form className='form' method='post'  onSubmit={addOpenAnswer}>
            <label form={'answer'}>
                Відповідь
            </label>
            {!isExist &&
                <div>
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
                    <button type="submit">Add answer</button>
                </div>
            }
            {isExist && !isChecked && <div style={{background:"bisque"}}>
                <input
                    type="checkbox"
                    checked={isChecked}
                    onChange={e=>setIsChecked(!isChecked)}
                />
                {newAnswer}
                <button onClick={()=>handleDelete()}>Delete answer</button>
            </div>}
            {isExist && isChecked && <div style={{background:"bisque"}}>
                <input
                    type="checkbox"
                    checked={isChecked}
                    onChange={e=>setIsChecked(!isChecked)}
                />
                <div>
                    <input maxLength={256}
                           minLength={2}
                           required={true}
                           name="answer_update_content"
                           placeholder=""
                           type='text'
                           onChange={e=>setNewAnswer(e.target.value)}
                           value={newAnswer}
                           id='answer_update_content'/>
                    <button onClick={()=>handleUpdate()}>Оновити</button>
                </div>
                <button onClick={()=>handleDelete()}>Delete answer</button>
            </div>}
        </form>
    )
}

export default AnswerForm