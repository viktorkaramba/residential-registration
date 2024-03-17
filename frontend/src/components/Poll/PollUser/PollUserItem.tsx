import React, {useCallback, useEffect, useState} from "react";
import TestAnswerUserList from "../TestAnswer/TestAnswerUser/TestAnswerUserList";
import AnswerForm from "../Answer/AnswerForm";
import config from "../../../config";
import err from "../../../err";
import {useAppContext} from "../../../AppContext";
import {useNavigate} from "react-router-dom";


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

    function updateAnswer(pollID:any, content:any){
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
                        console.log(data)
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
        <li>
            <label>
                {poll.question}
            </label>
            {poll.test_answer.length !== 0 && <TestAnswerUserList answers={poll.test_answer} pollID={poll.id}
                                                                   userAnswer={userAnswer} deleteAnswer={deleteAnswer}/>}
            {poll.test_answer.length === 0 && <AnswerForm pollID={poll.id} userAnswer={userAnswer} isExist={isExist}
                                                           setIsExist={setIsExist} deleteAnswer={deleteAnswer}
                                                          updateAnswer={updateAnswer}/>}
        </li>
    )
}

export default PollUserItem