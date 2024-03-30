import React, {useContext, useState} from "react";
import {jwtDecode} from "jwt-decode";


const AppContext = React.createContext('light');

const AppProvider = ({children}: { children?: React.ReactNode }) => {
    const [activeOSBBElement, setActiveOSBBElement] = useState("OSBBList");
    const [activePollElement, setActivePollElement] = useState("1");
    const [token, setToken] = useState("");
    const [isLogin, setIsLogin] = useState(()=>{
        let token = localStorage.getItem('token') || '';
        if(token !==''){
            let decodedToken = jwtDecode(token);
            let currentDate = new Date();

            // @ts-ignore
            return decodedToken.exp * 1000 >= currentDate.getTime();
        }
        else {
            return false
        }
    });
    const [osbbID, setOsbbID] = useState(0);
    const [poll, setPoll] = useState<any>(null);
    return (
        // @ts-ignore
        <AppContext.Provider value = {{
            activeOSBBElement, setActiveOSBBElement, activePollElement, setActivePollElement,
            osbbID, setOsbbID, isLogin, setIsLogin, token, setToken, poll, setPoll
        }}>
            {children}
        </AppContext.Provider>
    )
}

export const useAppContext = () => {
    return useContext(AppContext);
}

export {AppContext, AppProvider};