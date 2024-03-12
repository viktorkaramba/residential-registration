import React, {useState} from "react";

const TestAnswerUserItem = ({content, id, addTestAnswer}:any) =>{
    const [selectedValue, setSelectedValue] = useState('');

    const handleChange = (e:any) => {
        console.log(e)
        setSelectedValue(e);
        addTestAnswer(e);
    };

    return(
        <li>
            <label>
                <input
                    type="radio"
                    checked={selectedValue===id}
                    value={id}
                    onChange={e=>handleChange(e.target.value)}
                />
                {content}
            </label>
        </li>
    )
}

export default TestAnswerUserItem