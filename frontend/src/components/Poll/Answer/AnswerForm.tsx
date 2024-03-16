import React, {useState} from "react";
import config from "../../../config";
import {useAppContext} from "../../../AppContext";
import err from "../../../err";

const AnswerForm = ({pollID, userAnswers, isExist, setIsExist}:any) =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [newAnswer, setNewAnswer] = useState('');
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
            {isExist && <div>
                Content
            </div>}
        </form>
    )
}

export default AnswerForm