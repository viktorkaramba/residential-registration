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
    const [userAnswers, setUserAnswers] = useState([]);
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
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchUserAnswers, navigate:navigate});
                    }else {
                        if(data){
                            const answers = data.map(
                                (answerSingle: { id:any, pollID: any; userID: any; test_answer_id: any; content: any;
                                    created_at: any; updated_at: any }) => {
                                    const {id, pollID, userID, test_answer_id, content, created_at, updated_at} = answerSingle;
                                    return {
                                        id: id,
                                        pollID: pollID,
                                        userID: userID,
                                        test_answer_id: test_answer_id,
                                        content: content,
                                        created_at: created_at,
                                        updated_at: updated_at,
                                    }
                                });
                            setUserAnswers(answers)
                        }else {
                            setUserAnswers([])
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, []);

    useEffect(() => {
        fetchUserAnswers();
    }, []);

    useEffect(() => {
       setIsExist(userAnswers.some((answer:any) => answer.content.length!==0))
    }, [userAnswers]);

    return(
        <li>
            <label>
                {poll.question}
            </label>
            {poll.test_answer.length !== 0 && <TestAnswerUserList answers={poll.test_answer} pollID={poll.id}
                                                                   userAnswers={userAnswers}/>}
            {poll.test_answer.length === 0 && <AnswerForm pollID={poll.id} userAnswers={userAnswers} isExist={isExist}
                                                           setIsExist={setIsExist}/>}
        </li>
    )
}

export default PollUserItem