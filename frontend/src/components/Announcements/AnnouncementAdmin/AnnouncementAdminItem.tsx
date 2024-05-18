import React, {useEffect, useState} from "react";
import {format} from "date-fns";
import Checkbox from '@mui/material/Checkbox';
import AnnouncementForm from "../AnnouncementForm/AnnouncementForm";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import {Button} from "@mui/material";

const AnnouncementAdminItem = ({announcement, deleteAnnouncement, updateAnnouncement}:any) => {
    const [newTitle, setNewTitle] = useState(announcement.Title);
    const [newContent, setNewContent] = useState(announcement.Content);
    const [isAnnouncementChecked, setIsAnnouncementChecked] = useState(false);
    const [isAddedChecked, setIsAddedChecked] = useState(false);
    const [show, setShow] = useState(false)

    function handleClickDelete(){
        setShow(true)
    }

    function handleDelete(){
        deleteAnnouncement(announcement.ID);
    }
    function handleUpdate(){
        updateAnnouncement(announcement.ID, newTitle, newContent, setIsAnnouncementChecked)
    }

    const handleClose = ()=>{
        setShow(false)
    }
    return(
        <div>
            <Dialog
                open={show}
                onClose={handleClose}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                <DialogTitle id="alert-dialog-title">
                    <div className={'text-black fw-7 fs-24 dialog-style'}>{"Попередження"}</div>
                </DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        <div className={'text-black fw-4 fs-20 dialog-style'}>Ви впевнені, що хочете видалити оголошення?</div>
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}><span className={'fw-4 fs-16 dialog-style'}>Відмінити</span></Button>
                    <Button onClick={handleDelete} autoFocus>
                        <span className={'fw-4 fs-16 dialog-style'}>Видалити</span>
                    </Button>
                </DialogActions>
            </Dialog>
            <div className={'flex flex-end m-5'}>
                {!isAddedChecked &&
                    <>
                        <Checkbox
                            name="announcement_check_box"
                            id='announcement_check_box'
                            checked={isAnnouncementChecked}
                            size="large"
                            style={{color:'var(--blue-color)'}}
                            onChange={()=>{setIsAnnouncementChecked(!isAnnouncementChecked)}}
                        />
                    </>}
            </div>
            {isAddedChecked &&  <AnnouncementForm/>}
            {!isAddedChecked &&
                <div className='announcements-item'>
                    {!isAnnouncementChecked &&
                        <div className="inner-wrap">
                  <div className='flex flex-sb flex-wrap'>
                      <div className='announcements-item-info-item fw-7 fs-26'>
                          <span> {newTitle}</span>
                      </div>
                      <div className='announcements-item-info-item fw-6 fs-15'>
                          <span> {format(announcement.CreatedAt, 'MMMM do yyyy, hh:mm:ss a')}</span>
                      </div>
                  </div>
                  <div className={'announcements-item-info'}>
                      <div className='announcements-item-info-item fw-6 fs-18'>
                          <span>{newContent}</span>
                      </div>
                  </div>
              </div>}
                    {isAnnouncementChecked &&
                        <div className="inner-wrap">
                            <div className='flex flex-sb flex-wrap'>
                                <label className='announcements-item-info-item fw-7 fs-26' form={'title'}>Новий заголовок оголошення
                                    <input maxLength={256}
                                           minLength={2}
                                           name="announcement_update_title"
                                           placeholder=""
                                           type='text'
                                           onChange={e=>setNewTitle(e.target.value)}
                                           value={newTitle}
                                           id='announcement_update_title'/>
                                </label>
                            </div>
                            <div className={'announcements-item-info'}>
                                <label form={'content'}>Новий зміст оголошення
                                    <textarea minLength={2} onChange={e=>{setNewContent(e.target.value)}} required={true} disabled={false} name="content" placeholder="" id='content' value={newContent}/>
                                </label>
                            </div>
                        </div>
                    }
                    {isAnnouncementChecked &&
                        <button className='button announcement_button' type="submit" onClick={()=>handleUpdate()} name="update_annpuncement">
                            <span className="button_content announcement_button_content">Оновити</span>
                        </button>
                    }
                    <button className='button announcement_button' type="submit" onClick={()=>handleClickDelete()} name="update_annpuncement">
                        <span className="button_content announcement_button_content">Видалити</span>
                    </button>
                </div>
            }
        </div>
    )
}

export default AnnouncementAdminItem