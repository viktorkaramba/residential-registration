import React, {useEffect, useState} from "react";

const TestAnswerUserItem = ({content, id, userAnswers, addTestAnswer, selectedValue, setSelectedValue, deleteAnswer}:any) =>{

    const handleChange = (e:any) => {
        setSelectedValue(e);
        addTestAnswer(e);
    };

    useEffect(() => {
        if(userAnswers.some((answer:any) => answer.test_answer_id === id)){
            setSelectedValue(id)
        }
    }, [userAnswers]);

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
                <button
                    onClick={()=>deleteAnswer(id)}
                >
                    Delete
                </button>
            </label>
        </li>
    )
}

export default TestAnswerUserItem