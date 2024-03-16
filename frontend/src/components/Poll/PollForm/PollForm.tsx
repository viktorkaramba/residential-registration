import React from "react";
import config from "../../../config";
import {useAppContext} from "../../../AppContext";

const PollForm = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const addPoll = (event: any) => {
        console.log('handleSubmit ran');
        event.preventDefault();

        // 👇️ access input values using name prop
        const question = event.target.question.value;
        const finished_at = new Date(event.target.finished_at.value);
        const requestOptions = {
            method: 'POST',
            headers:config.headers,
            body: JSON.stringify({ question: question, finished_at: finished_at.toISOString()})
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/polls', requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    error.HandleError({error, addPoll});
                }
            });
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };
    return(
        <form className='form' method='post'  onSubmit={addPoll}>
            <label form={'question'}>
                Запитання
            </label>
            <input maxLength={256} minLength={2} required={true} name="question" placeholder="" type='text' id='question'/>
            <label form={'finished_at'}>
               Дата завершення
            </label>
            <input required={true} name="finished_at" placeholder="" type='datetime-local' step="1" id='finished_at'/>
            <button type="submit">Add poll</button>
        </form>
    )
}

export default PollForm