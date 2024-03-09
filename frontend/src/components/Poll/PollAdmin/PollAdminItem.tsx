import React, {useState} from "react";
import TestAnswerAdminList from "../TestAnswer/TestAnswerAdmin/TestAnswerAdminList";


const PollAdminItem = ({poll, updatePoll, deletePoll}:any) => {
    const [isChecked, setIsChecked] = useState(false);
    const [newQuestion, setNewQuestion] = useState(poll.question);
    const [newFinished, setNewFinished] = useState("");

    return(
        <li>
            <label>
                <input
                    type="checkbox"
                    checked={isChecked}
                    onChange={e=>setIsChecked(!isChecked)}
                />
                {!isChecked && poll.question}
                {isChecked &&
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
                        <button onClick={()=>updatePoll(poll.id, newQuestion, newFinished)}>Оновити</button>
                    </div>
                }
            </label>
            <button onClick={()=>deletePoll(poll.id)}>Видалити</button>
            {poll.test_answer.length !== 0 &&  <TestAnswerAdminList answers={poll.test_answer}/>}
        </li>
    )
}

export default PollAdminItem