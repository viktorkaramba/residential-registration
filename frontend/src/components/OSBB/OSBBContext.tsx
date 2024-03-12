import React, {useContext, useEffect, useState} from "react";


const OSBBContext = React.createContext('light');

const OSBBProvider = ({children}: { children?: React.ReactNode }) => {
    const [activeOSBBElement, setActiveOSBBElement] = useState("1");
    const [isLogin] = useState(()=>{
        let token = localStorage.getItem('token') || '';
        if(token!==''){
            return true;
        }else {
            return false;
        }
    });
    const [osbbID, setOsbbID] = useState(0);

    return (
        // @ts-ignore
        <OSBBContext.Provider value = {{
            activeOSBBElement, setActiveOSBBElement, osbbID, setOsbbID, isLogin
        }}>
            {children}
        </OSBBContext.Provider>
    )
}

export const useOSBBContext = () => {
    return useContext(OSBBContext);
}

export {OSBBContext, OSBBProvider};