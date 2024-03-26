import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import AnnouncementAdminItem from "./AnnouncementAdminItem";
import AnnouncementForm from "../AnnouncementForm/AnnouncementForm";
import {MdAddCircle} from "react-icons/md";
import {IoListCircleSharp} from "react-icons/io5";

const AnnouncementAdminList = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()

    const [announcements, setAnnouncements] = useState([]);
    const navigate = useNavigate();
    const [isAddedChecked, setIsAddedChecked] = useState(false);

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

    function updateAnnouncement(announcementID:any, title:any, content:any, setIsAnnouncementChecked:any){
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
                        err.HandleError({errorMsg:error, func:updateAnnouncement, navigate:navigate});
                    }else {
                        setAnnouncements((currentAnnouncement:any) => {
                            return currentAnnouncement.map((announcement:any)=>{
                                if(announcement.ID === announcementID){
                                    return {...announcement, title, content}
                                }
                                return announcement
                            })
                        })
                        setIsAnnouncementChecked(false);
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
                    err.HandleError({errorMsg:error, func:deleteAnnouncement, navigate:navigate});
                }else {
                    setAnnouncements((currentAnnouncements: any) => {
                        return currentAnnouncements.filter((announcement:any) => announcement.ID !== announcementID)
                    })
                }
            });
    }

    function addAnnouncement (announcement: any){
        // @ts-ignore
        setAnnouncements(currentAnnouncement => {
            return [
                ...currentAnnouncement,
                announcement,
            ]
        })
    }

    useEffect(() => {
        fetchAnnouncements();
    }, [fetchAnnouncements]);

    return(
        <section className='announcements-list'>
            <div className='container'>
                <div className={'flex flex-end m-5'}>
                    {!isAddedChecked && <MdAddCircle fontSize={'40px'} style={{color:'var(--blue-color)'}} onClick={()=>setIsAddedChecked(!isAddedChecked)}/>}
                    {isAddedChecked &&
                        <IoListCircleSharp fontSize={'40px'} style={{color:'var(--blue-color)'}} onClick={()=>setIsAddedChecked(!isAddedChecked)}/>}

                </div>
                {isAddedChecked &&  <AnnouncementForm addAnnouncement={addAnnouncement}/>}
                {!isAddedChecked &&  <div className='announcements-content grid'>
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
                </div>}
            </div>
        </section>
    )
}

export default AnnouncementAdminList