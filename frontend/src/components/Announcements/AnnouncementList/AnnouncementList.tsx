import {useCallback, useEffect, useState} from "react";
import AnnouncementListElement from "./AnnouncementListElement";
import config from "../../../config";
import {useOSBBContext} from "../../OSBB/OSBBContext";

const AnnouncementList = () =>{
    // @ts-ignore
    const {osbbID} = useOSBBContext()
    const [announcements, setAnnouncements] = useState([]);
    const fetchAnnouncements = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: config.headers,
            }
            console.log("osbbID " + osbbID)
            const response = await fetch(config.apiUrl+'osbb/'+ osbbID+ '/announcements', requestOptions);
            const data = await response.json();
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
            setAnnouncements(announcements)
        } catch(error){
            console.log(error);
        }
    }, []);
    useEffect(() => {
        fetchAnnouncements();
    }, [fetchAnnouncements]);
    return(
        <div>
            <section className='booklist'>
                <div className='container'>
                    <div className='booklist-content grid'>
                        {
                            announcements.slice(0, 30).map((item, index) => {
                                return (
                                    // @ts-ignore
                                    <AnnouncementListElement key = {index}{...item}/>
                                )
                            })
                        }
                    </div>
                </div>
            </section>
        </div>
    )
}

export default AnnouncementList