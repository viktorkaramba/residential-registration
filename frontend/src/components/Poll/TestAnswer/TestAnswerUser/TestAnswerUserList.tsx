import config from "../../../../config";
import {useState} from "react";
import {useOSBBContext} from "../../../OSBB/OSBBContext";
import TestAnswerUserItem from "./TestAnswerUserItem";

const TestAnswerUserList = ({answers, pollID}:any) =>{
    // @ts-ignore
    const {osbbID} = useOSBBContext()
    const [testAnswers, setTestAnswers] = useState(answers);
    function addTestAnswer(testAnswerID:any){
        const requestOptions = {
            method: 'POST',
            headers:config.headers,
            body: JSON.stringify({test_answer_id: testAnswerID}),
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls/'+pollID+'/answers-test', requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                // @ts-ignore
                setTestAnswers(currentTestAnswer => {
                    return [
                        ...currentTestAnswer,
                        {test_answer_id:testAnswerID}
                    ]
                })
            });
    }


    return(
        <ul>
            {testAnswers.map((answer: {id:any, content:any})=>{
                return(
                   <TestAnswerUserItem {...answer} key={answer.id} addTestAnswer={addTestAnswer}/>
                )
            })}
        </ul>
    )
}

export default TestAnswerUserList