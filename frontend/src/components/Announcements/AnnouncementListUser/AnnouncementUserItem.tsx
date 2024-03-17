import React from "react";


const AnnouncementUserItem = ((announcement:any) => {

    return(
        <div>
            <br></br>
            {announcement.Title}
            <br></br>
            {announcement.Content}
            <br></br>
        </div>
    )
});

export default AnnouncementUserItem