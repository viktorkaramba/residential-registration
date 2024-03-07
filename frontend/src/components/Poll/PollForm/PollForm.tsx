import React from "react";
import config from "../../../config";
import {useOSBBContext} from "../../OSBB/OSBBContext";

const PollForm = () =>{
    // @ts-ignore
    const {osbbID} = useOSBBContext()
    const handleSubmit = (event: any) => {
        console.log('handleSubmit ran');
        event.preventDefault();

        // 👇️ access input values using name prop
        const question = event.target.question.value;
        const finished_at = event.target.finished_at.value;
        const requestOptions = {
            method: 'POST',
            headers:config.headers,
            body: JSON.stringify({ question: question, finished_at: finished_at})
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/polls', requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
            });
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };
    return(
        <form className='form' method='post'  onSubmit={handleSubmit}>
            <label form={'question'}>
                Запитання
            </label>
            <input maxLength={256} minLength={2} required={true} name="question" placeholder="" type='text' id='question'/>
            <label form={'finished_at'}>
               Дата завершення
            </label>
            <input required={true} name="finished_at" placeholder="" type='date' id='finished_at'/>
            <button type="submit">Submit form</button>
        </form>
    )
}

export default PollForm