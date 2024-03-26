import React from "react";

const TestAnswerFormItem = ({content, index, deleteTestAnswer}:any) =>{
    return(
        <li className={'flex flex-sb'}>
            {content}
            <button className='button' type="submit" name="delete_test_answer"  onClick={()=>deleteTestAnswer(index)}>
                <span className="button_content">Видалити</span>
            </button>
        </li>
    )
}

export default TestAnswerFormItem