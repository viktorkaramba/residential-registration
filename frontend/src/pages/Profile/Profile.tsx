import React, {useCallback, useEffect, useState} from "react";
import config from "../../utils/config";
import {useAppContext} from "../../utils/AppContext";
import err from "../../utils/err";
import {useNavigate} from "react-router-dom";
import ProfileNavbar from "../../components/Navbar/ProfileNavbar";
import ProfileMenu from "../../components/Menu/ProfileMenu";
import ProfileForm from "../../components/Profile/ProfileForm/ProfileForm";


const Profile = () => {
    return (
        <ProfileMenu/>
    )
}

export default Profile