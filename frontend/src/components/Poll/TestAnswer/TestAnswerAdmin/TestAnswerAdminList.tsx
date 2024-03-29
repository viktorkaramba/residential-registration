import TestAnswerAdminItem from "./TestAnswerAdminItem";
import config from "../../../../utils/config";
import {useState} from "react";
import {useAppContext} from "../../../../utils/AppContext";
import err from "../../../../utils/err";
import {useNavigate} from "react-router-dom";

const TestAnswerAdminList = ({answers, pollID}:any) =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [testAnswers, setTestAnswers] = useState(answers);
    const navigate = useNavigate();

    function updateTestAnswer(id:any, content:any, setIsChecked:any){

        let body = null;
        if(content != null){
            body = JSON.stringify({content: content});
        }else {

        }

        const requestOptions = {
            method: 'PUT',
            headers:{ 'Content-Type': 'application/json',
                'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            body: body,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls/' + pollID +'/tests/'+id, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:updateTestAnswer, navigate:navigate});
                }else {
                    setTestAnswers((currentAnswer: any[]) => {
                        return currentAnswer.map((answer:any)=>{
                            if(answer.id === id){
                                return {...answer}
                            }
                            return answer
                        })
                    });
                    setIsChecked(false);
                }
            });
    }

    function deleteTestAnswer(id:any){
        const requestOptions = {
            method: 'DELETE',
            headers:{ 'Content-Type': 'application/json',
                'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls/' + pollID + '/tests/'+id, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:deleteTestAnswer, navigate:navigate});
                }else {
                    setTestAnswers((currentAnswer: any[]) => {
                        return currentAnswer.filter(answer => answer.id !== id)
                    })
                }
            });
    }

    return(
        <ul className={'test_answer_list'} >
            {testAnswers.map((answer: {id:any, content:any})=>{
                return(
                    <TestAnswerAdminItem
                        {...answer}
                        count_test_answers={testAnswers.length}
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