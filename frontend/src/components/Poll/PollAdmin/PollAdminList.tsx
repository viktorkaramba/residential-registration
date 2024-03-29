import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import PollAdminItem from "./PollAdminItem";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import '../Poll.css'

const PollAdminList = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [polls, setPolls] = useState([]);
    const navigate = useNavigate();
    const fetchPolls = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers:{ 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            }
           fetch(config.apiUrl+'osbb/'+ osbbID+ '/polls', requestOptions)
               .then(response => response.json())
               .then(data => {
                   console.log(data)
                   const {error}:any = data;
                   if(error){
                       err.HandleError({errorMsg:error, func:fetchPolls, navigate:navigate});
                   }else {
                       if(data){
                           const polls = data.map(
                               (pollSingle: { id:any, question: any; test_answer: any; type: any; created_at: any;
                                   finished_at: any; createdAt: any; updatedAt: any; is_closed:any}) => {
                                   const {id, question, test_answer, type, created_at, finished_at, createdAt, updatedAt, is_closed} = pollSingle;
                                   return {
                                       id: id,
                                       question: question,
                                       test_answer: test_answer,
                                       type: type,
                                       created_at: created_at,
                                       finished_at: finished_at,
                                       createdAt: createdAt,
                                       updatedAt: updatedAt,
                                       is_closed: is_closed
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

    function updatePoll(id:any, question:any, finished_at:any, isClosed:any, setIsPollChecked:any){
        if(finished_at!==""){
            finished_at = new Date(finished_at);
        }else {
            finished_at = null;
        }
        let questionJSON = null
        let finishedAtJSON = null
        let isClosedJSON = null
        if(question !== "" ){
            questionJSON = {question: question};
        }
        if(finished_at != null){
            finishedAtJSON={finished_at:finished_at};
        }
        if(isClosed != null) {
            isClosedJSON = {is_closed:isClosed};
        }
        let body = {...questionJSON, ...finishedAtJSON, ...isClosedJSON};
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
                    console.log(error)
                    err.HandleError({errorMsg:error, func:updatePoll, navigate:navigate});
                }else {
                    setPolls((currentPoll:any) => {
                        return currentPoll.map((poll:any)=>{
                            if(poll.id === id){
                                return {...poll, question, finished_at, isClosed}
                            }
                            return poll
                        })
                    })
                    setIsPollChecked(false);
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
                    error.HandleError({errorMsg:error, func:deletePoll, navigate:navigate});
                }else {
                    setPolls((currentPoll: any) => {
                        return currentPoll.filter((poll:any) => poll.id !== id)
                    })
                }
            });
    }

    return(
        <section className='poll_list'>
            <div className='container'>
                <div className='poll_content grid'>
                    {polls.length === 0 && "Немає опитувань"}
                    {
                        polls.map((poll:{id:any, question:any, finished_at:any, is_closed:any}) => {
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
                </div>
            </div>
        </section>
    )
}

export default PollAdminList