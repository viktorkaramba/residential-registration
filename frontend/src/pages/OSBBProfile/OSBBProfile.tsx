import React, {useCallback, useEffect, useState} from "react";
import config from "../../config";
import Header from "../../components/Header/Header";
import AnnouncementList from "../../components/Announcements/AnnouncementList/AnnouncementList";
import {useOSBBContext} from "../../components/OSBB/OSBBContext";
import PollTestForm from "../../components/Poll/PollForm/PollTestForm";
import PollAdminList from "../../components/Poll/PollAdmin/PollAdminList";
import PollForm from "../../components/Poll/PollForm/PollForm";
import PollUserList from "../../components/Poll/PollUser/PollUserList";
import errorHandling from "../../error";
import RefreshToken from "../../auth";


const OSBBProfile = () => {
    // @ts-ignore
    const {setOsbbID} = useOSBBContext();
    const [osbb, setOSBB] = useState<any>(null);
    const fetchOSBB = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers:config.headers,
            };
            fetch(config.apiUrl+'osbb/profile', requestOptions)
                .then(response => response.json())
                .then(data =>{
                    const {error}:any = data;
                    if(error){
                        console.log(error);
                        if(error.includes(errorHandling.tokenIsRevoked) || error.includes(errorHandling.tokenExpired)){
                            if(RefreshToken()){
                                
                            }
                        }
                    }
                });

            // if(data){
            //     const {announcements, building, createdAt, edrpou, id, name, osbb_head, rent, updatedAt}:any = data;
            //     const newOSBB = {
            //         announcements: announcements,
            //         building: building,
            //         createdAt: createdAt,
            //         edrpou: edrpou,
            //         id : id,
            //         name: name,
            //         osbb_head: osbb_head,
            //         rent: rent,
            //         updatedAt: updatedAt
            //     };
            //     console.log(newOSBB);
            //     setOSBB(newOSBB);
            //     setOsbbID(newOSBB.id);
            // }else {
            //     setOSBB(null);
            // }
        } catch(error){
            console.log(error);
        }
    }, []);
    useEffect(() => {
        fetchOSBB();
    }, []);


    
    return (
        <div>
            <Header/>
            <div>{osbb?.id}</div>
            <div>{osbb?.name}</div>
            <br />
            <div>{osbb?.osbb_head?.full_name?.first_name}</div>
            <br/>
            {/*<AnnouncementList key={osbb?.id}/>*/}
            {/*<br/>*/}
            <div style={{background:"peru"}}>
                <PollForm/>
            </div>
            <div style={{background:"bisque"}}>
                <PollTestForm/>
            </div>
            {/*<div style={{background:"green"}}>*/}
            {/*    <PollUserList/>*/}
            {/*</div>*/}
        </div>
    )
}

export default OSBBProfile