import React from "react";
import config from "../../../config";
import {useAppContext} from "../../../AppContext";
import error from "../../../err";
import RefreshToken from "../../../auth";
import err from "../../../err";

const AnnouncementForm = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const addAnnouncement = (event: any) => {
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
                    err.HandleError({errorMsg:error, func:addAnnouncement});
                }
            });
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };
    return(
        <form className='form' method='post'  onSubmit={addAnnouncement}>
            <label form={'title'}>
               Назва оголошення
            </label>
            <input maxLength={256} minLength={2} required={true} name="title" placeholder="" type='text' id='title'/>
            <label form={'content'}>
                Оголошення
            </label>
            <input maxLength={256} minLength={2} required={true} name="content" placeholder="" type='text' id='content'/>
            <button type="submit">Submit form</button>
        </form>
    )
}

export default AnnouncementForm