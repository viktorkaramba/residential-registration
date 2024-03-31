import React, {useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import {Stack} from "@mui/material";
import Alert from "@mui/material/Alert";
import "../Poll.css"

const PollForm = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const navigate = useNavigate();
    const [isSuccess, setIsSuccess]= useState(false);
    const [errorDate, setErrorDate]= useState(false);
    const addPoll = (event: any) => {
        console.log('handleSubmit ran');
        event.preventDefault();

        // 👇️ access input values using name prop
        const question = event.target.question.value;
        const finished_at = new Date(event.target.finished_at.value);
        if(finished_at < new Date()){
            setErrorDate(!errorDate);
            return
        }
        const requestOptions = {
            method: 'POST',
            headers:{ 'Content-Type': 'application/json',
                'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            body: JSON.stringify({ question: question, finished_at: finished_at.toISOString()})
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/polls', requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data)
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:addPoll, navigate:navigate});
                }else {
                    // const {id}:any = data
                    // const newAnnouncement = {
                    //     ID: id,
                    //     Title: title,
                    //     Content: content,
                    //     CreatedAt:new Date()
                    // };
                    // addAnnouncement(newAnnouncement)
                    setIsSuccess(true);
                }
            });
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };
    return(
        <form method='post'  onSubmit={addPoll}>
            <div className={'flex flex-wrap align-items-start bg-dark-grey'}>
                <div className="form poll_form">
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

export default PollForm