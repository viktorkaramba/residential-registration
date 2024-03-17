import React, {useState} from "react";

const TestAnswerAdminItem = ({content, id, updateTestAnswer, deleteTestAnswer}:any) =>{
    const [isChecked, setIsChecked] = useState(false);
    const [newContent, setNewContent] = useState(content);

    return(
        <li>
            <label>
                <input
                    type="checkbox"
                    checked={isChecked}
                    onChange={e=>setIsChecked(!isChecked)}
                />
                {!isChecked && newContent}
                {isChecked &&
                <div>
                    <input maxLength={256}
                           minLength={2}
                           required={true}
                           name="test_answer_update_content"
                           placeholder=""
                           type='text'
                           onChange={e=>setNewContent(e.target.value)}
                           value={newContent}
                           id='test_answer_update_content'/>
                    <button onClick={()=>updateTestAnswer(id, newContent, setIsChecked)}>Оновити</button>
                </div>
                }
            </label>
            <button onClick={()=>deleteTestAnswer(id)}>Видалити</button>
        </li>
    )
}

export default TestAnswerAdminItem