import React, {useCallback, useEffect, useState} from "react";
import TestAnswerUserList from "../TestAnswer/TestAnswerUser/TestAnswerUserList";
import AnswerForm from "../Answer/AnswerForm";
import config from "../../../utils/config";
import err from "../../../utils/err";
import {useAppContext} from "../../../utils/AppContext";
import {useNavigate} from "react-router-dom";
import {format} from "date-fns";

const PollUserItem = ({poll}:any) => {
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [userAnswer, setUserAnswer] = useState<any>(null);
    const navigate = useNavigate();
    const [isExist, setIsExist] = useState(false);

    const fetchUserAnswers = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: { 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            }
            fetch(config.apiUrl+'osbb/'+ osbbID+ '/polls/' + poll.id + '/user-answers', requestOptions)
                .then(response => response.json())
                .then(data => {
                        if(data){
                            const {error}:any = data;
                            if(error){
                                err.HandleError({errorMsg:error, func:fetchUserAnswers, navigate:navigate});
                            }else {
                                const { id, pollID, userID, test_answer_id, content, created_at, updated_at }:any =data
                                const newAnswer ={
                                    id: id,
                                    pollID: pollID,
                                    userID: userID,
                                    test_answer_id: test_answer_id,
                                    content: content,
                                    created_at: created_at,
                                    updated_at: updated_at,
                                }
                                console.log(newAnswer)
                                setUserAnswer(newAnswer)
                            }
                        }else {
                            setUserAnswer(null)
                        }
                });
        } catch(error){
            console.log(error);
        }
    }, []);

    function updateAnswer(pollID:any, content:any, setIsChecked:any){
        const requestOptions = {
            method: 'PUT',
            headers:config.headers,
            body: JSON.stringify({ content: content})
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls/'+pollID + '/answers', requestOptions)
            .then(response => response.json())
            .then(data => {
                if(data){
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:updateAnswer});
                    }else {
                        setIsChecked(false);
                    }
                }

            });
    }

    function deleteAnswer(pollID:any){
        const requestOptions = {
            method: 'DELETE',
            headers:config.headers,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls/'+pollID + '/answers', requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:deleteAnswer});
                }
            });
    }

    useEffect(() => {
        fetchUserAnswers();
    }, []);

    useEffect(() => {
        if(userAnswer != null){
            setIsExist(userAnswer.content.length!==0)
        }
    }, [userAnswer]);

    return(
        <div className='poll-item'>
            <div className="inner-wrap">
                <div className='flex flex-sb flex-wrap'>
                    <div className='polls-item-info-item fw-7 fs-26'>
                        <span>{poll.question}</span>
                    </div>
                    <div className='polls-item-info-item fw-6 fs-15'>
                        <span> {format(poll.created_at, 'MMMM do yyyy, hh:mm:ss a')}</span>
                    </div>
                </div>
                <div className='flex flex-sb flex-wrap'>
                    <div className='polls-item-info-item fw-7 fs-15'>
                        <span>Завершується - {format(poll.finished_at, 'MMMM do yyyy, hh:mm:ss a')}</span>
                    </div>
                </div>
            </div>
            {poll.test_answer.length !== 0 && <TestAnswerUserList answers={poll.test_answer} pollID={poll.id}
                                                                  userAnswer={userAnswer} deleteAnswer={deleteAnswer}/>}
            {poll.test_answer.length === 0 && <AnswerForm pollID={poll.id} userAnswer={userAnswer} isExist={isExist}
                                                          setIsExist={setIsExist} deleteAnswer={deleteAnswer}
                                                          updateAnswer={updateAnswer}/>}
        </div>
    )
}

export default PollUserItem