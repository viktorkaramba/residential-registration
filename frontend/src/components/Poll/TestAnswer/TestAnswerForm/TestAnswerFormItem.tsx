
const TestAnswerFormItem = ({content, index, deleteAnswer}:any) =>{
    return(
        <li>
            <label>
                {content}
            </label>
            <button
                onClick={()=>deleteAnswer(index)}
            >
                Delete
            </button>
        </li>
    )
}

export default TestAnswerFormItem