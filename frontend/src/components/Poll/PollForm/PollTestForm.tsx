import React, {useEffect, useState} from "react";
import config from "../../../config";
import {useOSBBContext} from "../../OSBB/OSBBContext";
import TestAnswerForm from "../TestAnswer/TestAnswerForm/TestAnswerForm";
import TestAnswerFormList from "../TestAnswer/TestAnswerForm/TestAnswerFormList";

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

    function makeRequest({question,finished_at }:any){
        const requestOptions = {
            method: 'POST',
            headers:config.headers,
            body: JSON.stringify({ question: question, test_answer:answers, finished_at: finished_at.toISOString() })
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/polls-test', requestOptions)
            .then(response =>response.json())
            .then(data => {
                console.log(data)
                const {error}:any = data;
                if(error){
                    let argument = { question, finished_at};
                    error.HandleError({error, makeRequest, argument});
                }else{
                    if(data){
                        localStorage.removeItem("TestAnswers")
                    }
                }
            });
    }
    const handleAddPollTest = (event: any) => {
        console.log('handleSubmit ran');
        event.preventDefault();

        if(answers.length < 2){
            console.log("<2")
            return
        }
        // 👇️ access input values using name prop
        const question = event.target.question.value;
        const finished_at = new Date(event.target.finished_at.value);

        makeRequest({question, finished_at});
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };
    return(
        <div>
        <form className='form' onSubmit={handleAddPollTest}>
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
            <TestAnswerFormList answers={answers} deleteAnswer={deleteAnswer}/>
            <button>Add poll test</button>
        </form>
        </div>
    )
}

export default PollTestForm