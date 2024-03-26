import React, {useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import Alert from "@mui/material/Alert";
import {Stack} from "@mui/material";

const AnnouncementForm = ({addAnnouncement}:any) =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const [isSuccess, setIsSuccess]= useState(false);
    const navigate = useNavigate();
    const handleAddAnnouncement = (event: any) => {
        event.preventDefault();

        // 👇️ access input values using name prop
        const title = event.target.title.value;
        const content = event.target.content.value;
        const requestOptions = {
            method: 'POST',
            headers: config.headers,
            body: JSON.stringify({ title: title, content: content})
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/announcements', requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:handleAddAnnouncement, navigate:navigate});
                }else {
                    const {id}:any = data
                    const newAnnouncement = {
                        ID: id,
                        Title: title,
                        Content: content,
                        CreatedAt:new Date()
                    };
                    addAnnouncement(newAnnouncement)
                    setIsSuccess(true);
                }
            });
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };

    return(
        <form method='post'  onSubmit={handleAddAnnouncement}>
            <div className={'flex flex-wrap align-items-start bg-dark-grey'}>
                    <div className="form announcement_form">
                        <div className="section">Заголовок та вміст оголошення</div>
                        <div className="inner-wrap">
                            <label form={'title'}>Заголовок оголошення
                                <input maxLength={256} minLength={2} required={true} name="title" placeholder="" type='text' id='title'/>
                            </label>
                            <label form={'content'}>Оголошення
                                <textarea minLength={2} required={true} name="content" placeholder="" id='content'/>
                            </label>
                        </div>
                        <div className={'flex flex-c'}>
                            <button className='button add_osbb' type="submit" name="submit_osbb">
                                <span className="button_content add_osbb_content">Додати ОСББ</span>
                            </button>
                        </div>
                        {isSuccess &&
                            <Stack sx={{margin: '10px'}} spacing={2}>
                                <Alert variant={'filled'} severity="success" style={{fontSize:'15px'}}>Оголошення успішно додане!</Alert>
                            </Stack>}
                    </div>
                </div>
        </form>
    )
}

export default AnnouncementForm