import React, {useCallback, useEffect, useState} from "react";
import config from "../../../config";
import {useAppContext} from "../../../AppContext";
import error from "../../../err";
import err from "../../../err";
import {useNavigate} from "react-router-dom";
import AnnouncementAdminItem from "./AnnouncementAdminItem";
import AnnouncementForm from "../AnnouncementForm/AnnouncementForm";

const AnnouncementAdminList = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [announcements, setAnnouncements] = useState([]);
    const navigate = useNavigate();

    const fetchAnnouncements = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: { 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            }
            fetch(config.apiUrl+'osbb/'+ osbbID+ '/announcements', requestOptions)
                .then(response => response.json())
                .then(data=>{
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchAnnouncements,
                            navigate:navigate});
                    }else {
                        if(data){
                            const announcements = data.slice(0, 20).map(
                                (announcementSingle: { ID: any; UserID: any; OSBBID: any; Title: any; Content: any; CreatedAt: any; updatedAt: any }) => {
                                    const {ID, UserID, OSBBID, Title, Content, CreatedAt, updatedAt} = announcementSingle;
                                    return {
                                        ID: ID,
                                        UserID: UserID,
                                        OSBBID: OSBBID,
                                        Title: Title,
                                        Content: Content,
                                        CreatedAt: CreatedAt,
                                        updatedAt: updatedAt
                                    }
                                });
                            setAnnouncements(announcements);
                        }else {
                            setAnnouncements([]);
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, []);

    function updateAnnouncement(announcementID:any, title:any, content:any){
        const requestOptions = {
            method: 'PUT',
            headers:config.headers,
            body: JSON.stringify({ title:title, content: content})
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/announcements/'+announcementID, requestOptions)
            .then(response => response.json())
            .then(data => {
                if(data){
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:updateAnnouncement});
                    }else {
                        setAnnouncements((currentAnnouncement:any) => {
                            return currentAnnouncement.map((announcement:any)=>{
                                if(announcement.ID === announcementID){
                                    return {...announcement, title, content}
                                }
                                return announcement
                            })
                        })
                    }
                }
            });
    }

    function deleteAnnouncement(announcementID:any){
        const requestOptions = {
            method: 'DELETE',
            headers:config.headers,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/announcements/'+announcementID, requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:deleteAnnouncement});
                }else {
                    setAnnouncements((currentAnnouncements: any) => {
                        return currentAnnouncements.filter((announcement:any) => announcement.ID !== announcementID)
                    })
                }
            });
    }

    useEffect(() => {
        fetchAnnouncements();
    }, [fetchAnnouncements]);

    return(
        <div>
            <ul>
                {
                    announcements.map((announcement:{ID:any, Title:any, Content:any})=> {
                        return (
                            <AnnouncementAdminItem key={announcement.ID}
                                                   announcement={announcement}
                                                   deleteAnnouncement={deleteAnnouncement}
                                                   updateAnnouncement={updateAnnouncement}/>
                        )
                    })
                }
            </ul>
            <AnnouncementForm/>
        </div>
    )
}

export default AnnouncementAdminList