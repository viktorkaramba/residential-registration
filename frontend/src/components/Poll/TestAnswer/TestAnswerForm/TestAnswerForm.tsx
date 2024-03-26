import React, {useState} from "react";
import {Stack} from "@mui/material";
import Alert from "@mui/material/Alert";

const TestAnswerForm = (props:any) =>{
    const [newAnswer, setAnswer] = useState("");
    const handleAddTestAnswer = (event: any) => {
        event.preventDefault();

        if(newAnswer === "")return

        props.addTestAnswer(newAnswer);
        setAnswer("");

        // // 👇️ clear all input values in the form
        // event.target.reset();
    };

    return(
      <>
          <div className="inner-wrap">
              <label form={'test_answer'}>Нова відповідь
                  <input maxLength={256} minLength={2}
                         value={newAnswer}
                         onChange={e=>setAnswer(e.target.value)}
                         name="test_answer" placeholder="" type='text' id='test_answer'/> </label>
          </div>
          <div className={'flex'}>
              <button className='button' type="submit" name="submit_poll" onClick={handleAddTestAnswer}>
                  <span className="button_content">Додати Відповідь</span>
              </button>
          </div>
      </>
    )
}

export default TestAnswerForm

