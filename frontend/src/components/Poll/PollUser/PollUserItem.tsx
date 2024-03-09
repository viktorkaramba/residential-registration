import React, {useState} from "react";
import TestAnswerAdminList from "../TestAnswer/TestAnswerAdmin/TestAnswerAdminList";
import TestAnswerUserList from "../TestAnswer/TestAnswerUser/TestAnswerUserList";


const PollUserItem = ({poll}:any) => {

    return(
        <li>
            <label>
                {poll.question}
            </label>
            {poll.test_answer.length !== 0 &&  <TestAnswerUserList answers={poll.test_answer} pollID={poll.id}/>}
        </li>
    )
}

export default PollUserItem