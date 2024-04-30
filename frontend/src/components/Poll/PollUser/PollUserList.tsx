import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import PollUserItem from "./PollUserItem";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import PollAdminItem from "../PollAdmin/PollAdminItem";

const PollUserList = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()

    const [polls, setPolls] = useState([]);
    const navigate = useNavigate();
    const fetchPolls = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: { 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            }
            fetch(config.apiUrl+'osbb/'+ osbbID+ '/polls', requestOptions)
                .then(response => response.json())
                .then(data => {
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchPolls, navigate:navigate});
                    }else {
                        if(data){
                            const polls = data.map(
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
                            setPolls(polls)
                        }else {
                            setPolls([])
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, []);

    useEffect(() => {
        fetchPolls();
    }, []);

    return(
        <section className='poll_list'>
            <div className='container'>
                <div className='poll_content grid'>
                    {polls.length === 0 && <h1 style={{color:"white"}}>Немає опитувань</h1>}
                    {
                        polls.map((poll:{id:any, question:any, finished_at:any}) => {
                            return (
                                <PollUserItem
                                    poll={poll}
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

export default PollUserList