
const TestAnswerFormItem = ({content, index, deleteTestAnswer}:any) =>{
    return(
        <li>
            <label>
                {content}
            </label>
            <button
                onClick={()=>deleteTestAnswer(index)}
            >
                Delete
            </button>
        </li>
    )
}

export default TestAnswerFormItem