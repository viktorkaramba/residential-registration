import React, {useEffect, useState} from "react";

const TestAnswerUserItem = ({content, id, pollID, userAnswer, addTestAnswer, selectedValue, setSelectedValue, deleteAnswer}:any) =>{
    const handleChange = (e:any) => {
        setSelectedValue(e);
        addTestAnswer(e);
    };


    function handleDelete(){
        deleteAnswer(pollID)
        setSelectedValue(undefined)
    }

    useEffect(() => {
        if(userAnswer != null){
            if(userAnswer.test_answer_id === id){
                setSelectedValue(id)
            }else {
            }
        }
    }, [userAnswer]);

    return(
        <li className={'flex flex-sb'}>
            <label  className={'flex'}>
                <input
                    type="radio"
                    checked={selectedValue==id}
                    value={id}
                    onChange={e=>handleChange(e.target.value)}
                />
                <span className={'m-5'}>{content}</span>
            </label>
            {selectedValue ==id &&  <button className='button' type="submit" name="delete_test_answer"   onClick={()=>handleDelete()}>
                <span className="button_content">Видалити</span>
            </button>}
        </li>
    )
}

export default TestAnswerUserItem