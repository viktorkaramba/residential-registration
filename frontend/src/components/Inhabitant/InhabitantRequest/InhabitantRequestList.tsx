import {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import InhabitantRequestItem from "./InhabitantRequestItem";
import AnnouncementUserItem from "../../Announcements/AnnouncementListUser/AnnouncementUserItem";

const InhabitantRequestList = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()

    const [inhabitantsWaitApprove, setInhabitantsWaitApprove] = useState([]);
    const navigate = useNavigate();
    const fetchInhabitantWaitApprove = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers: { 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            }
            fetch(config.apiUrl+'osbb/'+ osbbID+ '/inhabitants/wait-approve', requestOptions)
                .then(response => response.json())
                .then(data => {
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchInhabitantWaitApprove, navigate:navigate});
                    }else {
                        if(data){
                            const inhabitantWaitApproveList = data.map(
                                (inhabitantWaitApproveList: { id:any, osbbid: any; apartment: any; full_name: any; phone_number: any;
                                    role: any; is_approved: any; createdAt: any}) => {
                                    const {id, osbbid, apartment, full_name, phone_number, role, is_approved, createdAt} = inhabitantWaitApproveList;
                                    return {
                                        id: id,
                                        osbbid: osbbid,
                                        apartment: apartment,
                                        full_name: full_name,
                                        phone_number: phone_number,
                                        role: role,
                                        is_approved: is_approved,
                                        createdAt:createdAt
                                    }
                                });
                            setInhabitantsWaitApprove(inhabitantWaitApproveList);
                            console.log(data)
                        }else {
                            setInhabitantsWaitApprove([]);
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, []);

    useEffect(() => {
        fetchInhabitantWaitApprove();
    }, []);

    function updateStatus(userID:any, answer:any){
        const requestOptions = {
            method: 'POST',
            headers:config.headers,
            body: JSON.stringify({ userID: userID, answer:answer})
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/inhabitants/approve', requestOptions)
            .then(response => response.json())
            .then(data => {
                if(data){
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:updateStatus, navigate:navigate});
                    }else {
                        setInhabitantsWaitApprove((currentInhabitant:any) => {
                            return currentInhabitant.filter((inhabitant:any) => inhabitant.id !== userID)
                        })
                    }
                }
            });
    }

    return(
        <section className='announcements-list'>
            <div className='container'>
                <div className='announcements-content grid'>
                    {inhabitantsWaitApprove.length===0 && <h1 style={{color:"white"}}>Немає запитів</h1>}
                    {
                        inhabitantsWaitApprove.map((inhabitantWaitApprove:{id:any, osbbid:any, apartment:any, full_name:any,
                            phone_number:any, role:any, is_approved:any}) => {
                            return (
                                <InhabitantRequestItem
                                    inhabitantWaitApprove={inhabitantWaitApprove}
                                    updateStatus={updateStatus}
                                    key={inhabitantWaitApprove.id}
                                />
                            )
                        })
                    }
                </div>
            </div>
        </section>
    )
}

export default InhabitantRequestList