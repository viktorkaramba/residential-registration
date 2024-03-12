import React from "react";
import config from "../../../config";
import {useOSBBContext} from "../../OSBB/OSBBContext";
import error from "../../../error";
import RefreshToken from "../../../auth";

const AnnouncementForm = () =>{
    // @ts-ignore
    const {osbbID} = useOSBBContext()
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
                // if(data.toString().includes(error.tokenExpire)){
                //     if (RefreshToken()){
                //         addAnnouncement(event);
                //     }
                // }
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