import React, {useState} from "react";
import 'chart.js/auto';
import {Bar} from "react-chartjs-2";

const TestAnswer = ({result, testAnswers}:any) => {

    const [labels] = useState(()=>{
        return testAnswers.map((answer: {id:any, content:any}, index:any) => index + 1);
    })
    const [data] = useState(()=>{
        return result.count_of_test_answers.map((answer: {test_answer_id:any, count:any}) => answer.count);
    })
    return(
        <div className={'answer'}>
            <Bar
                data={{
                    labels:labels,
                    datasets: [
                        {
                            label:'Відповіді',
                            data: data
                        },
                    ]
                }}
                options={{
                    plugins: {
                        title: {
                            text: 'Результати опитування'
                        }
                    },
                }}
            />
            <ul className={'test_answer_list'}>
                {testAnswers.map((answer: {id:any, content:any}, index:any)=>{
                    return(
                        <li className={'flex flex-sb'}>
                            <span className={'m-5'}>{answer.content}</span>
                            <div>{index + 1}</div>
                        </li>
                    )
                })}
            </ul>
        </div>
    )
}

export default TestAnswer