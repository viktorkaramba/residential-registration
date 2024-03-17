import React, {useContext, useState} from "react";


const AppContext = React.createContext('light');

const AppProvider = ({children}: { children?: React.ReactNode }) => {
    const [activeOSBBElement, setActiveOSBBElement] = useState("OSBBList");
    const [activePollElement, setActivePollElement] = useState("1");
    const [token, setToken] = useState("");
    const [isLogin, setIsLogin] = useState(()=>{
        let token = localStorage.getItem('token') || '';
        return token !== '';
    });
    const [osbbID, setOsbbID] = useState(0);

    return (
        // @ts-ignore
        <AppContext.Provider value = {{
            activeOSBBElement, setActiveOSBBElement, activePollElement, setActivePollElement,
            osbbID, setOsbbID, isLogin, setIsLogin, token, setToken
        }}>
            {children}
        </AppContext.Provider>
    )
}

export const useAppContext = () => {
    return useContext(AppContext);
}

export {AppContext, AppProvider};