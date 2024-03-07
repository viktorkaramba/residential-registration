import TestAnswerFormItem from "./TestAnswerFormItem";

const TestAnswerFormList = ({answers, deleteAnswer}:any) =>{
    return(
        <ul>
            {answers.length === 0 && "Немає відповідей"}
            {answers.map((answer: {content:any}, index:any)=>{
                return(
                  <TestAnswerFormItem
                      {...answer}
                      index={index}
                      key={index}
                      deleteAnswer={deleteAnswer}
                  />
                )
            })}
        </ul>
    )
}

export default TestAnswerFormList