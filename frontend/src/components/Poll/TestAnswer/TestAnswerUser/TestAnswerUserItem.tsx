import React, {useEffect, useState} from "react";

const TestAnswerUserItem = ({content, id, userAnswer, addTestAnswer, selectedValue, setSelectedValue, deleteAnswer}:any) =>{

    const handleChange = (e:any) => {
        setSelectedValue(e);
        addTestAnswer(e);
    };



    useEffect(() => {
        if(userAnswer != null){
            if(userAnswer.test_answer_id === id){
                setSelectedValue(id)
            }
        }
    }, [userAnswer]);

    return(
        <li>
            <label>
                <div>
                    <input
                        type="radio"
                        checked={selectedValue==id}
                        value={id}
                        onChange={e=>handleChange(e.target.value)}
                    />
                    {content}
                </div>
            </label>
        </li>
    )
}

export default TestAnswerUserItem