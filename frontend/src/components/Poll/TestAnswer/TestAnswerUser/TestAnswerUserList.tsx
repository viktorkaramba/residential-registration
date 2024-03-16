import config from "../../../../config";
import React, {useState} from "react";
import {useAppContext} from "../../../../AppContext";
import TestAnswerUserItem from "./TestAnswerUserItem";
import err from "../../../../err";
import {useNavigate} from "react-router-dom";

const TestAnswerUserList = ({answers, pollID, userAnswer, deleteAnswer}:any) =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [testAnswers] = useState(answers);
    const navigate = useNavigate();
    const [selectedValue, setSelectedValue] = useState();


    function addTestAnswer(testAnswerID:any){
        const requestOptions = {
            method: 'POST',
            headers:config.headers,
            body: JSON.stringify({test_answer_id: parseInt(testAnswerID)}),
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls/'+pollID+'/answers-test', requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:addTestAnswer, navigate:navigate});
                }
                console.log(data)
            });
    }

    function handleDelete(){
        deleteAnswer(pollID)
        setSelectedValue(undefined)
    }
    return(
        <ul>
            {testAnswers.map((answer: {id:any, content:any})=>{
                return(
                   <TestAnswerUserItem {...answer} key={answer.id} userAnswer={userAnswer} addTestAnswer={addTestAnswer}
                                       selectedValue={selectedValue} setSelectedValue={setSelectedValue}
                                       deleteAnswer={deleteAnswer}/>
                )
            })}
            <button
                onClick={()=>handleDelete()}
            >
                Delete
            </button>
        </ul>
    )
}

export default TestAnswerUserList