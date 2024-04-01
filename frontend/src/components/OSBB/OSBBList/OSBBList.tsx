import {useCallback, useEffect, useState} from "react";
import OSBBListElement from "./OSBBListElement";
import config from "../../../utils/config";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import "../OSBB.css"

const OSBBList = () =>{
    const [osbbs, setOSBBS] = useState([]);
    const navigate = useNavigate();
    const fetchOSBBS = useCallback(async() => {
        try{
            fetch(config.apiUrl+'osbb/')
                .then(response => response.json())
                .then(data => {
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchOSBBS, navigate:navigate});
                    }else {
                        if(data){
                            const newOSBBS = data.slice(0, 20).map(
                                (osbbSingle: { id: any; building: any; announcements: any; osbb_head: any; name: any; edrpou: any; photo: any}) => {
                                    const {id, building, announcements, osbb_head, name, edrpou, photo} = osbbSingle;
                                    return {
                                        id: id,
                                        building: building,
                                        announcements: announcements,
                                        osbb_head: osbb_head,
                                        name: name,
                                        edrpou: edrpou,
                                        photo:photo
                                    }
                                });
                            setOSBBS(newOSBBS);
                            console.log(newOSBBS)
                        }
                        else {
                            setOSBBS([]);
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, []);

    useEffect(() => {
        fetchOSBBS();
    }, [fetchOSBBS]);

    return(
        <section className='osbblist'>
            <div className='container'>
                <div className='osbblist-content grid'>
                    {
                        osbbs.map((item, index) => {
                            return (
                                // @ts-ignore
                                <OSBBListElement key = {index}{...item}/>
                            )
                        })
                    }
                </div>
            </div>
        </section>
    )
}

export default OSBBList