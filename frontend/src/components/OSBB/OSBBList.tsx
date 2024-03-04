import {useCallback, useState} from "react";
import OSBB from "./OSBB";

const OSBBList = () =>{
    const [osbbs, setOSBBS] = useState([]);
    const [resultTitle, setResultTitle] = useState("");
    const fetchOSBBS = useCallback(async() => {
        try{
            const response = await fetch(`${URL}+'osbb'`);
            const data = await response.json();
            const {docs} = data;
            console.log(data)
            if(docs){
                const newOSBBS = docs.slice(0, 20).map(
                    (osbbSingle: { key: any; author_name: any; cover_i: any; edition_count: any; first_publish_year: any; title: any; }) => {
                    const {key, author_name, cover_i, edition_count, first_publish_year, title} = osbbSingle;

                    return {
                        id: key,
                        author: author_name,
                        cover_id: cover_i,
                        edition_count: edition_count,
                        first_publish_year: first_publish_year,
                        title: title
                    }
                });

                setOSBBS(newOSBBS);

                if(newOSBBS.length > 1){
                } else {
                    setResultTitle("No OSBB Found!")
                }
            } else {
                setOSBBS([]);
                setResultTitle("No OSBB Found!")
            }
        } catch(error){
            console.log(error);
        }
    }, []);
    fetchOSBBS();
    return(
        <div>
            <section className='booklist'>
                <div className='container'>
                    <div className='section-title'>
                        <h2>{resultTitle}</h2>
                    </div>
                    <div className='booklist-content grid'>
                        {
                            osbbs.slice(0, 30).map((item, index) => {
                                return (
                                    <OSBB key = {index}/>
                                )
                            })
                        }
                    </div>
                </div>
            </section>
        </div>
    )
}

export default OSBBList