import TestAnswerFormItem from "./TestAnswerFormItem";
import '../TestAnswer.css'

const TestAnswerFormList = ({answers, deleteTestAnswer}:any) =>{
    return(
        <ul className={'test_answer_list'}>
            {answers.length === 0 && "Немає відповідей"}
            {answers.map((answer: {content:any}, index:any)=>{
                return(
                  <TestAnswerFormItem
                      {...answer}
                      index={index}
                      key={index}
                      deleteTestAnswer={deleteTestAnswer}
                  />
                )
            })}
        </ul>
    )
}

export default TestAnswerFormList