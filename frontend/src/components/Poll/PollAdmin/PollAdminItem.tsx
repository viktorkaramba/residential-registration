import React, {useState} from "react";
import TestAnswerAdminList from "../TestAnswer/TestAnswerAdmin/TestAnswerAdminList";


const PollAdminItem = ({poll, updatePoll, deletePoll}:any) => {
    const [isPollChecked, setIsPollChecked] = useState(false);
    const [newQuestion, setNewQuestion] = useState(poll.question);
    const [newFinished, setNewFinished] = useState("");
    const [newIsClosed, setNewIsClosed] = useState(poll.is_closed);

    function handleDelete(){
        deletePoll(poll.id);
    }

    function handleUpdate(){
        updatePoll(poll.id, newQuestion, newFinished, newIsClosed, setIsPollChecked)
    }

    return(
        <li>
            <label>
                <input
                    type="checkbox"
                    checked={isPollChecked}
                    onChange={()=>setIsPollChecked(!isPollChecked)}
                />
                {!isPollChecked &&
                    <div>
                        {newQuestion}
                        <br/>
                        {newIsClosed && <p>Закрите</p>}
                        {!newIsClosed && <p>Відкрите</p>}
                    </div>}
                {isPollChecked &&
                    <div>
                        <input maxLength={256}
                               minLength={2}
                               required={true}
                               name="poll_update_question"
                               placeholder=""
                               type='text'
                               onChange={e=>setNewQuestion(e.target.value)}
                               value={newQuestion}
                               id='poll_update_question'/>
                        <input name="poll_update_finished_at"
                               placeholder=""
                               value={newFinished}
                               onChange={e=>setNewFinished(e.target.value)}
                               type='datetime-local'
                               step="1"
                               id='poll_update_finished_at'/>
                        <label>
                            <input name="poll_update_is_closed"
                                   checked={newIsClosed}
                                   onChange={e=>setNewIsClosed(!newIsClosed)}
                                   type='checkbox'
                                   id='poll_update_is_closed'/>
                            Завершити
                        </label>
                        <button onClick={()=>handleUpdate()}>Оновити</button>
                    </div>
                }
            </label>
            <button onClick={()=>handleDelete()}>Видалити</button>
            {poll.test_answer.length !== 0 &&  <TestAnswerAdminList answers={poll.test_answer}/>}
        </li>
    )
}

export default PollAdminItem