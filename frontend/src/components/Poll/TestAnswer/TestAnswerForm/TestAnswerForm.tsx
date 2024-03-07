import React, {useState} from "react";

const TestAnswerForm = (props:any) =>{
    const [newAnswer, setAnswer] = useState("");

    const handleAddTestAnswer = (event: any) => {
        event.preventDefault();
        console.log('handleSubmit ran', newAnswer);

        if(newAnswer === "")return

        props.addTestAnswer(newAnswer);
        setAnswer("");
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };

    return(
        <div>
            <label form={'test_answer'}>
                Нова відповідь
            </label>
            <input maxLength={256} minLength={2}
                   value={newAnswer}
                   onChange={e=>setAnswer(e.target.value)}
                   name="test_answer" placeholder="" type='text' id='test_answer'/>
            <button onClick={handleAddTestAnswer}>Add answer</button>
        </div>
    )
}

export default TestAnswerForm

