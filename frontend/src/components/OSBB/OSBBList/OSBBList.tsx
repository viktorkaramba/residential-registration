import {useCallback, useEffect, useState} from "react";
import OSBBListElement from "./OSBBListElement";
import config from "../../../config";

const OSBBList = () =>{
    const [osbbs, setOSBBS] = useState([]);
    const fetchOSBBS = useCallback(async() => {
        try{
            const response = await fetch(config.apiUrl+'osbb/');
            const data = await response.json();
            const newOSBBS = data.slice(0, 20).map(
                    (osbbSingle: { id: any; building: any; announcements: any; osbb_head: any; name: any; edrpou: any; }) => {
                        const {id, building, announcements, osbb_head, name, edrpou} = osbbSingle;
                        return {
                            id: id,
                            building: building,
                            announcements: announcements,
                            osbb_head: osbb_head,
                            name: name,
                            edrpou: edrpou
                        }
                    });
            setOSBBS(newOSBBS)
        } catch(error){
            console.log(error);
        }
    }, []);
    useEffect(() => {
        fetchOSBBS();
    }, [fetchOSBBS]);
    return(
        <div>
            <section className='booklist'>
                <div className='container'>
                    <div className='booklist-content grid'>
                        {
                            osbbs.slice(0, 30).map((item, index) => {
                                return (
                                    // @ts-ignore
                                    <OSBBListElement key = {index}{...item}/>
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