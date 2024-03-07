import React, {useEffect, useState} from "react";
import config from "../../../config";
import {useOSBBContext} from "../../OSBB/OSBBContext";
import TestAnswerForm from "../TestAnswer/TestAnswerForm";
import TestAnswerList from "../TestAnswer/TestAnswerList";

const PollTestForm = () =>{
    // @ts-ignore
    const {osbbID} = useOSBBContext();
    const [answers, setAnswers] = useState(()=>{
        const localValue = localStorage.getItem("TestAnswers")
        if(localValue==null)return[]
        return JSON.parse(localValue)
    });

    useEffect(()=>{
        localStorage.setItem("TestAnswers", JSON.stringify(answers));
    }, [answers])

    function addTestAnswer(content:any){
        // @ts-ignore
        setAnswers(currentAnswer => {
            return [
                ...currentAnswer,
                {content: content},
            ]
        })
    }
    function deleteAnswer(index:any){
        setAnswers((currentAnswer: any[]) => {
            return currentAnswer.filter((answer, i) => index !== i)
        })
    }
    const handleSubmit = (event: any) => {
        console.log('handleSubmit ran');
        event.preventDefault();

        if(answers.length < 2){
            console.log("<2")
            return
        }
        // 👇️ access input values using name prop
        const question = event.target.question.value;
        const finished_at = new Date(event.target.finished_at.value);
        const requestOptions = {
            method: 'POST',
            headers:config.headers,
            body: JSON.stringify({ question: question, test_answer:answers, finished_at: finished_at.toISOString() })
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/polls-test', requestOptions)
            .then(response => {
                console.log(response.json())
            })
            .then(data => {
                console.log(data);
                localStorage.removeItem("TestAnswers")
            });
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };
    return(
        <div>
        <form className='form' onSubmit={handleSubmit}>
            <label form={'question'}>
                Запитання
            </label>
            <input maxLength={256} minLength={2} required={true} name="question" placeholder="" type='text' id='question'/>
            <label form={'finished_at'}>
                Дата завершення
            </label>
            <input required={true} name="finished_at" placeholder="" type='datetime-local' step="1"
                   id='finished_at'/>
            <TestAnswerForm addTestAnswer={addTestAnswer}/>
            <h1>Тестові відповіді</h1>
            <TestAnswerList answers={answers} deleteAnswer={deleteAnswer}/>
            <button>Submit form</button>
        </form>
        </div>
    )
}

export default PollTestForm