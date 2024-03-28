import {useCallback, useEffect, useState} from "react";
import AnnouncementUserItem from "./AnnouncementUserItem";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import "../Announcements.css"

const AnnouncementUserList = () =>{
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

    useEffect(() => {
        fetchAnnouncements();
    }, [fetchAnnouncements]);
    return(
        <section className='announcements-list'>
            <div className='container'>
                <div className='announcements-content grid'>
                    {
                        announcements.map((item, index) => {
                            return (
                                // @ts-ignore
                                <AnnouncementUserItem key = {index}{...item}/>
                            )
                        })
                    }
                </div>
            </div>
        </section>
    )
}

export default AnnouncementUserList