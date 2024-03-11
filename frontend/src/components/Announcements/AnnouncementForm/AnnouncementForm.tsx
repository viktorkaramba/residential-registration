import React from "react";
import config from "../../../config";
import {useOSBBContext} from "../../OSBB/OSBBContext";

const AnnouncementForm = () =>{
    // @ts-ignore
    const {osbbID} = useOSBBContext()
    const handleSubmit = (event: any) => {
        console.log('handleSubmit ran');
        event.preventDefault();

        // 👇️ access input values using name prop
        const title = event.target.title.value;
        const content = event.target.content.value;
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json',
                'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            body: JSON.stringify({ title: title, content: content})
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/announcements', requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
            });
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };
    return(
        <form className='form' method='post'  onSubmit={handleSubmit}>
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