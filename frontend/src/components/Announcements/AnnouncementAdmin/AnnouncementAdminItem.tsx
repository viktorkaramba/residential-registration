import React, {useEffect, useState} from "react";


const AnnouncementAdminItem = ({announcement, deleteAnnouncement, updateAnnouncement}:any) => {
    const [newTitle, setNewTitle] = useState(announcement.Title);
    const [newContent, setNewContent] = useState(announcement.Content);
    const [isAnnouncementChecked, setIsAnnouncementChecked] = useState(false);

    function handleDelete(){
        deleteAnnouncement(announcement.ID);
    }

    function handleUpdate(){
        updateAnnouncement(announcement.ID, newTitle, newContent, setIsAnnouncementChecked)
    }

    return(
        <li>
            <label>
                <input
                    name="announcement_check_box"
                    id='announcement_check_box'
                    type="checkbox"
                    checked={isAnnouncementChecked}
                    onChange={()=>{setIsAnnouncementChecked(!isAnnouncementChecked)}}
                />
                {!isAnnouncementChecked && <div>
                    {newTitle}
                    <br/>
                    {newContent}
                </div>}
                {isAnnouncementChecked &&
                    <div>
                        <input maxLength={256}
                               minLength={2}
                               name="announcement_update_title"
                               placeholder=""
                               type='text'
                               onChange={e=>setNewTitle(e.target.value)}
                               value={newTitle}
                               id='announcement_update_title'/>
                        <input
                               maxLength={256}
                               minLength={2}
                               name="announcement_update_content"
                               placeholder=""
                               value={newContent}
                               onChange={e=>setNewContent(e.target.value)}
                               type='text'
                               id='announcement_update_content'/>
                        <button id="announcement_update_button" onClick={handleUpdate}>Оновити</button>
                    </div>
                }
            </label>
            <button onClick={()=>handleDelete()}>Видалити</button>
        </li>
    )
}

export default AnnouncementAdminItem