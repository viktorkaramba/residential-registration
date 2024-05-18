import config from "./config";
import err from "./err";
import {Navigate, useLocation, useNavigate} from "react-router-dom";
import {useAppContext} from "./AppContext";
import React from "react";

function RefreshToken():any{
    let oldToken = localStorage.getItem('token') || '{}'
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ token:oldToken })
    }
    return fetch(config.apiUrl+'refresh-token', requestOptions)
        .then(response => response.json())
        .then(data => {
            const {token}:any = data
            localStorage.setItem('token', token);
        });
}

function Logout(navigate:any):any{
    const requestOptions = {
        method: 'GET',
        headers: { 'Content-Type': 'application/json',
            'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
    }
    return fetch(config.apiUrl+'auth/logout', requestOptions)
        .then(response => response.json())
        .then(data => {
            console.log(data)
            const {error}:any = data;
            if(error){
                err.HandleError({errorMsg:error, func:Logout, navigate:navigate});
            }else {
                if(data){
                    localStorage.removeItem("token");
                }
            }
        });
}

    function RequireAuth({ children }:any) {
    // @ts-ignore
    const { isLogin, setPrevPage } = useAppContext();
    const location = useLocation();
    setPrevPage(location.pathname)
    return isLogin === true ? (
        children
    ) : (
        <Navigate to="/login" replace state={{ path: location.pathname }} />
    );
}

export default {RefreshToken, Logout, RequireAuth}