import React, {useCallback, useContext, useEffect, useState} from "react";
import {jwtDecode} from "jwt-decode";
import config from "./config";
import err from "./err";
import {useNavigate} from "react-router-dom";


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
    const [user, setUser] = useState<any>(null);
    const fetchUserProfile = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers:{ 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            };
            fetch(config.apiUrl+'osbb/' + osbbID + '/inhabitants/profile', requestOptions)
                .then(response => response.json())
                .then(data =>{
                    console.log(data)
                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchUserProfile});
                    }else {
                        if(data){
                            const {osbbid, apartment, full_name, phone_number, role, is_approved}:any = data;
                            const userInfo = {
                                osbbid: osbbid,
                                apartment: apartment,
                                full_name: full_name,
                                phone_number: phone_number,
                                role : role,
                                is_approved: is_approved
                            };
                            setUser(userInfo);
                        }else {
                            setUser(null);
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, [token]);


    useEffect(() => {
        if(user === null) {
            fetchUserProfile();
        }
    }, []);

    return (
        // @ts-ignore
        <AppContext.Provider value = {{
            activeOSBBElement, setActiveOSBBElement, activePollElement, setActivePollElement,
            osbbID, setOsbbID, isLogin, setIsLogin, token, setToken, poll, setPoll, user, setUser
        }}>
            {children}
        </AppContext.Provider>
    )
}

export const    useAppContext = () => {
    return useContext(AppContext);
}

export {AppContext, AppProvider};