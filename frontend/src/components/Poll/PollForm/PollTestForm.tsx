import React, {useEffect, useRef, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import TestAnswerForm from "../TestAnswer/TestAnswerForm/TestAnswerForm";
import TestAnswerFormList from "../TestAnswer/TestAnswerForm/TestAnswerFormList";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import {Stack} from "@mui/material";
import Alert from "@mui/material/Alert";

const PollTestForm = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext();
    const navigate = useNavigate();
    const [isSuccess, setIsSuccess]= useState(false);
    const [errorDate, setErrorDate]= useState(false);
    const [errorCount, setErrorCount]= useState(false);
    const [answers, setAnswers] = useState(()=>{
        const localValue = localStorage.getItem("TestAnswers")
        if(localValue==null)return[]
        return JSON.parse(localValue)
    });

    const focusRef = useRef(null);
    useEffect(() => {
        // @ts-ignore
        focusRef.current.scrollIntoView({behavior: 'smooth'});
    }, []);

    useEffect(()=>{
        localStorage.setItem("TestAnswers", JSON.stringify(answers));
    }, [answers])

    function addTestAnswer(content:any){
        setErrorCount(false)
        // @ts-ignore
        setAnswers(currentAnswer => {
            return [
                ...currentAnswer,
                {content: content},
            ]
        })
    }

    function deleteTestAnswer(index:any){
        setAnswers((currentAnswer: any[]) => {
            return currentAnswer.filter((answer, i) => index !== i)
        })
    }

    function makeRequest({question,finished_at }:any){
        if(finished_at < new Date()){
            setErrorDate(!errorDate);
            return
        }
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
                    err.HandleError({errorMsg:error, func:makeRequest, argument:argument,
                        navigate:navigate});
                }else{
                    if(data){
                        setIsSuccess(true);
                        localStorage.removeItem("TestAnswers")
                    }
                }
            });
    }

    const handleAddPollTest = (event: any) => {
        console.log('handleSubmit ran');
        event.preventDefault();

        if(answers.length < 2){
            setErrorCount(true)
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
        <form method='post'  onSubmit={handleAddPollTest}>
            <div className={'flex flex-wrap align-items-start bg-dark-grey'} >
                <div className="form poll_form"  ref={focusRef} >
                    <h1>Форма для додання відкритого опитування</h1>
                    <div className="inner-wrap">
                        <label form={'question'}>Запитання
                            <input maxLength={256} minLength={2} required={true} name="question" placeholder="" type='text' id='question'/>
                        </label>
                        <label form={'finished_at'}>Дата завершення
                            <input required={true} name="finished_at" placeholder="" type='datetime-local' step="1" id='finished_at'/>
                        </label>
                        {errorDate &&
                            <div className={'error'}>
                                Дата завершення опитування повинна бути більша за поточну
                            </div>
                        }
                    </div>
                    <div className="section">Тестові відповіді</div>
                    <TestAnswerForm addTestAnswer={addTestAnswer}/>
                    <TestAnswerFormList answers={answers} deleteTestAnswer={deleteTestAnswer}/>
                    {errorCount &&
                        <div className={'error'} style={{marginBottom:'10px'}}>
                            Тестових відповідей повинно бут 2 або більше
                        </div>
                    }
                    <div className={'flex flex-c'}>
                        <button className='button poll_button' type="submit" name="submit_poll">
                            <span className="button_content poll_button_content">Додати Опитування</span>
                        </button>
                    </div>
                    {isSuccess &&
                        <Stack sx={{margin: '10px'}} spacing={2}>
                            <Alert variant={'filled'} severity="success" style={{fontSize:'15px'}}>Оголошення успішно додане!</Alert>
                        </Stack>}
                </div>
            </div>
        </form>
    )
}

export default PollTestForm