import React, {useCallback, useEffect, useState} from "react";
import config from "../../utils/config";
import Header from "../../components/Header/Header";
import {useAppContext} from "../../utils/AppContext";
import err from "../../utils/err";
import {useNavigate} from "react-router-dom";
import OSBBProfileMenu from "../../components/Menu/OSBBProfileMenu";
import OSBBDescription from "../../components/OSBB/OSBBDescription";


const OSBBProfile = () => {
    return (
        <div>
            <OSBBProfileMenu/>
        </div>
    )
}

export default OSBBProfile