import React, {useCallback, useEffect, useState} from "react";
import {format} from "date-fns";
import '../../Poll.css';
import config from "../../../../utils/config";
import err from "../../../../utils/err";
import {useNavigate} from "react-router-dom";
import {useAppContext} from "../../../../utils/AppContext";

import OpenAnswer from "./OpenAnswer";
import TestAnswer from "./TestAnswer";


const PollResultItem = () => {

    // @ts-ignore
    const {osbbID, poll} = useAppContext()
    const navigate = useNavigate();
    const [pollResult, setPollResult] = useState<any>(null);
    const fetchPollResult = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers:{ 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            }
            fetch(config.apiUrl+'osbb/'+ osbbID+ '/polls/'+poll.id +'/answers-results', requestOptions)
                .then(response => response.json())
                .then(data => {
                    console.log(poll.id)
                    console.log(data)
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchPollResult, navigate:navigate});
                    }else {
                        if(data){
                            const {answers, count_of_answers, count_of_test_answers,}:any = data;
                            const newPollResult = {
                                answers: answers,
                                count_of_answers: count_of_answers,
                                count_of_test_answers: count_of_test_answers,
                            }
                            setPollResult(newPollResult);
                        }else {
                            setPollResult(null);
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, [osbbID]);

    useEffect(() => {
        fetchPollResult();
    }, []);
    return(
        <section className='announcements-list'>
            <div className='container'>
                <div className='announcements-content grid'>
                    <div className='announcements-item'>
                        <div className='flex flex-sb flex-wrap'>
                            <div className='announcements-item-info-item fw-7 fs-26'>
                                Запитання: <span>{poll?.question}</span>
                            </div>
                            <div className='announcements-item-info-item fw-6 fs-15'>
                                <span>{format(poll?.created_at, 'MMMM do yyyy, hh:mm:ss a')}</span>
                            </div>
                        </div>
                        {poll.type === "open_answer" &&  pollResult?.answers.map((answer:{id:any, pollID:any, UserID:any,
                            test_answer_id:any, content:any, created_at:any, updated_at:any}) => {
                            return (
                                <OpenAnswer
                                    answer={answer}
                                    key={answer.id}
                                />
                            )
                        })}
                        {poll.type === "test" && pollResult !== null &&
                            <TestAnswer
                                result={pollResult}
                                testAnswers={poll.test_answer}
                            />}
                    </div>
                </div>
            </div>
        </section>
    )
}

export default PollResultItem