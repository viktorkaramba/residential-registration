import {useCallback, useEffect, useState} from "react";
import config from "../../../config";
import {useAppContext} from "../../../AppContext";
import PollAdminItem from "./PollAdminItem";
import err from "../../../err";

const PollAdminList = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [polls, setPolls] = useState([]);

    const fetchPolls = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: config.headers,
            }
           fetch(config.apiUrl+'osbb/'+ osbbID+ '/polls', requestOptions)
               .then(response => response.json())
               .then(data => {
                   const {error}:any = data;
                   if(error){
                       error.HandleError({error, fetchPolls});
                   }else {
                       if(data){
                           const polls = data.slice(0, 20).map(
                               (pollSingle: { id:any, question: any; test_answer: any; type: any; created_at: any;
                                   finished_at: any; createdAt: any; updatedAt: any }) => {
                                   const {id, question, test_answer, type, created_at, finished_at, createdAt, updatedAt} = pollSingle;
                                   return {
                                       id: id,
                                       question: question,
                                       test_answer: test_answer,
                                       type: type,
                                       created_at: created_at,
                                       finished_at: finished_at,
                                       createdAt: createdAt,
                                       updatedAt: updatedAt
                                   }
                               });
                           setPolls(polls);
                       }else {
                           setPolls([]);
                       }
                   }
               });
        } catch(error){
            console.log(error);
        }
    }, [osbbID]);

    useEffect(() => {
        fetchPolls();
    }, []);

    function updatePoll(id:any, question:any, isOpen:any, finished_at:any){
        if(finished_at!==""){
            finished_at = new Date(finished_at);
        }else {
            finished_at = null;
        }
        let questionJSON = null
        let finishedAtJSON = null
        let isOpenJSON = null
        if(question !== "" ){
            questionJSON = {question: question};
        }
        if(finished_at != null){
            finishedAtJSON={finished_at:finished_at};
        }
        if(isOpen != null) {
            isOpenJSON = {is_open:isOpen};
        }
        let body = {...questionJSON, ...finishedAtJSON, ...isOpenJSON};
        const requestOptions = {
            method: 'PUT',
            headers:config.headers,
            body:  JSON.stringify(body),
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls/'+id, requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    err.HandleError({error, updatePoll});
                }else {
                    setPolls((currentPoll:any) => {
                        return currentPoll.map((poll:any)=>{
                            if(poll.id === id){
                                return {...poll, question, finished_at}
                            }
                            return poll
                        })
                    })
                }
            });
    }

    function deletePoll(id:any){
        const requestOptions = {
            method: 'DELETE',
            headers:config.headers,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls/'+id, requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    error.HandleError({error, deletePoll});
                }else {
                    setPolls((currentPoll: any) => {
                        return currentPoll.filter((poll:any) => poll.id !== id)
                    })
                }
            });
    }

    return(
        <ul>
            {
                polls.map((poll:{id:any, question:any, finished_at:any}) => {
                    return (
                        <PollAdminItem
                            poll={poll}
                            updatePoll={updatePoll}
                            deletePoll={deletePoll}
                            key={poll.id}
                    />
                    )
                })
            }
        </ul>
    )
}

export default PollAdminList