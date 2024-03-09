import {useCallback, useEffect, useState} from "react";
import config from "../../../config";
import {useOSBBContext} from "../../OSBB/OSBBContext";
import PollAdminItem from "./PollAdminItem";

const PollAdminList = () =>{
    // @ts-ignore
    const {osbbID} = useOSBBContext()
    const [polls, setPolls] = useState([]);

    const fetchPolls = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: config.headers,
            }
            const response = await fetch(config.apiUrl+'osbb/'+ osbbID+ '/polls', requestOptions);
            const data = await response.json();
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
            console.log(polls)
            setPolls(polls)
        } catch(error){
            console.log(error);
        }
    }, [osbbID]);
    useEffect(() => {
        fetchPolls();
    }, []);
    function updatePoll(id:any, question:any, finished_at:any){
        console.log('handleSubmit ran');
        console.log(question, finished_at);
        if(finished_at!==""){
            finished_at = new Date(finished_at);
        }else {
            finished_at = null;
        }
        let body = null;
        if(question !== "" && finished_at != null){
            body = JSON.stringify({question: question, finished_at: finished_at.toISOString()});
        }
        if(question == null){
            body = JSON.stringify({finished_at: finished_at.toISOString()});
        }
        if(finished_at == null) {
            body = JSON.stringify({question: question});
        }

        const requestOptions = {
            method: 'PUT',
            headers:config.headers,
            body: body,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/polls/'+id, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);

                setPolls((currentPoll:any) => {
                    return currentPoll.map((poll:any)=>{
                        if(poll.id === id){
                            return {...poll, question, finished_at}
                        }
                        return poll
                    })
                })
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
                console.log(data);
                setPolls((currentPoll: any) => {
                    return currentPoll.filter((poll:any) => poll.id !== id)
                })
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