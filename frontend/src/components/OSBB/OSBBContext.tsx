import React, {useContext, useState} from "react";


const OSBBContext = React.createContext('light');

const OSBBProvider = ({children}: { children?: React.ReactNode }) => {
    const [activeOSBBElement, setActiveOSBBElement] = useState("1");
    const [osbbID, setOsbbID] = useState(0);

    return (
        // @ts-ignore
        <OSBBContext.Provider value = {{
            activeOSBBElement, setActiveOSBBElement, osbbID, setOsbbID
        }}>
            {children}
        </OSBBContext.Provider>
    )
}

export const useOSBBContext = () => {
    return useContext(OSBBContext);
}

export {OSBBContext, OSBBProvider};