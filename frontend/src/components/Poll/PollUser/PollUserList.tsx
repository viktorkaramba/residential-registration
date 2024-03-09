import {useCallback, useEffect, useState} from "react";
import config from "../../../config";
import {useOSBBContext} from "../../OSBB/OSBBContext";
import PollUserItem from "./PollUserItem";

const PollUserList = () =>{
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
            setPolls(polls)
        } catch(error){
            console.log(error);
        }
    }, [osbbID]);

    useEffect(() => {
        fetchPolls();
    }, []);



    return(
        <ul>
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
        </ul>
    )
}

export default PollUserList