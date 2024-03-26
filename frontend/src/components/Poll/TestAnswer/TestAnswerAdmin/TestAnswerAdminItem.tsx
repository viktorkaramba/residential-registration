import React, {useState} from "react";
import Checkbox from "@mui/material/Checkbox";

const TestAnswerAdminItem = ({content, id, updateTestAnswer, deleteTestAnswer}:any) =>{
    const [isChecked, setIsChecked] = useState(false);
    const [newContent, setNewContent] = useState(content);

    return(
        <li className={'flex flex-sb'}>
            <label  className={'flex'}>
                <Checkbox
                    name="announcement_check_box"
                    id='announcement_check_box'
                    checked={isChecked}
                    size="medium"
                    style={{color:'var(--blue-color)'}}
                    onChange={e=>setIsChecked(!isChecked)}
                />
                {!isChecked && <span className={'m-5'}>{newContent}</span>}
                {isChecked &&
                    <div className={'poll_update'}>
                        <input maxLength={256}
                               minLength={2}
                               required={true}
                               name="test_answer_update_content"
                               placeholder=""
                               type='text'
                               onChange={e=>setNewContent(e.target.value)}
                               value={newContent}
                               className='m-5'
                               id='test_answer_update_content'/>
                            <button className='button m-5' type="submit" name="delete_test_answer"   onClick={()=>updateTestAnswer(id, newContent, setIsChecked)}>
                                <span className="button_content">Оновити</span>
                            </button>
                    </div>
                }
            </label>
            <button className='button' type="submit" name="delete_test_answer"   onClick={()=>deleteTestAnswer(id)}>
                <span className="button_content">Видалити</span>
            </button>
        </li>
    )
}

export default TestAnswerAdminItem