import TestAnswerFormItem from "./TestAnswerFormItem";

const TestAnswerFormList = ({answers, deleteTestAnswer}:any) =>{
    return(
        <ul>
            {answers.length === 0 && "Немає відповідей"}
            {answers.map((answer: {content:any}, index:any)=>{
                return(
                  <TestAnswerFormItem
                      {...answer}
                      index={index}
                      key={index}
                      deleteAnswer={deleteTestAnswer}
                  />
                )
            })}
        </ul>
    )
}

export default TestAnswerFormList