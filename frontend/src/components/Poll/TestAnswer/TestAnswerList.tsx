import TestAnswerItem from "./TestAnswerItem";

const TestAnswerList = ({answers, deleteAnswer}:any) =>{
    return(
        <ul>
            {answers.length === 0 && "Немає відповідей"}
            {answers.map((answer: {content:any}, index:any)=>{
                return(
                  <TestAnswerItem
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

export default TestAnswerList