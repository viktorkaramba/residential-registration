import TestAnswerAdminItem from "./TestAnswerAdminItem";
import config from "../../../../config";
import {useState} from "react";
import {useOSBBContext} from "../../../OSBB/OSBBContext";

const TestAnswerAdminList = ({answers, pollID}:any) =>{
    // @ts-ignore
    const {osbbID} = useOSBBContext()
    const [testAnswers, setTestAnswers] = useState(answers);


    function updateTestAnswer(id:any, content:any){

        let body = null;
        if(content != null){
            body = JSON.stringify({content: content});
        }else {

        }

        const requestOptions = {
            method: 'PUT',
            headers:config.headers,
            body: body,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls-test/'+id, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                const {error}:any = data;
                if(error){
                    error.HandleError({error, updateTestAnswer});
                }else {
                    setTestAnswers((currentAnswer: any[]) => {
                        return currentAnswer.map((answer:any)=>{
                            if(answer.id === id){
                                return {...answer}
                            }
                            return answer
                        })
                    });
                }
            });
    }

    function deleteTestAnswer(id:any){
        const requestOptions = {
            method: 'DELETE',
            headers:config.headers,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls-test/'+id, requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    error.HandleError({error, deleteTestAnswer});
                }else {
                    setTestAnswers((currentAnswer: any[]) => {
                        return currentAnswer.filter(answer => answer.id !== id)
                    })
                }
            });
    }
    return(
        <ul>
            {testAnswers.map((answer: {id:any, content:any})=>{
                return(
                    <TestAnswerAdminItem
                        {...answer}
                        updateTestAnswer={updateTestAnswer}
                        deleteTestAnswer={deleteTestAnswer}
                        key={answer.id}
                    />
                )
            })}
        </ul>
    )
}

export default TestAnswerAdminList