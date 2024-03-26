import React, {useCallback, useEffect, useState} from "react";
import config from "../../utils/config";
import Header from "../../components/Header/Header";
import {useAppContext} from "../../utils/AppContext";
import err from "../../utils/err";
import {useNavigate} from "react-router-dom";
import OSBBProfileMenu from "../../components/Menu/OSBBProfileMenu";


const OSBBDescription = () => {
    // @ts-ignore
    const {setOsbbID} = useAppContext();
    // @ts-ignore
    const {token} = useAppContext();
    const navigate = useNavigate();
    const [osbb, setOSBB] = useState<any>(null);
    const fetchOSBB = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers:{ 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            };
            fetch(config.apiUrl+'osbb/profile', requestOptions)
                .then(response => response.json())
                .then(data =>{
                    console.log(data);
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchOSBB, navigate:navigate});
                    }else {
                        if(data){
                            const {announcements, building, createdAt, edrpou, id, name, osbb_head, rent, updatedAt}:any = data;
                            const newOSBB = {
                                announcements: announcements,
                                building: building,
                                createdAt: createdAt,
                                edrpou: edrpou,
                                id : id,
                                name: name,
                                osbb_head: osbb_head,
                                rent: rent,
                                updatedAt: updatedAt
                            };
                            setOSBB(newOSBB);
                            setOsbbID(newOSBB.id);
                        }else {
                            setOSBB(null);
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, [token]);

    useEffect(() => {
        fetchOSBB();
    }, []);

    return (
        <div>
            <div>{osbb?.name}</div>
            <br />
            <div>{osbb?.osbb_head?.full_name?.first_name}</div>
            <div>{osbb?.osbb_head?.full_name?.surname}</div>
            <div>{osbb?.osbb_head?.full_name?.patronymic}</div>
            {/*<br/>*/}
            {/*<AnnouncementUserList key={osbb?.id}/>*/}
            {/*<br/>*/}
            {/*<div style={{background:"peru"}}>*/}
            {/*    <PollForm/>*/}
            {/*</div>*/}
            {/*<div style={{background:"bisque"}}>*/}
            {/*    <PollTestForm/>*/}
            {/*</div>*/}
            {/*<div style={{background:"green"}}>*/}
            {/*    <PollUserList/>*/}
            {/*</div>*/}
        </div>
    )
}

export default OSBBDescription