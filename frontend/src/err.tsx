import RefreshToken from "./auth";
import {useNavigate} from "react-router-dom";
import React from "react";
import auth from "./auth";


const errorsMessages ={
    tokenExpired :'token is expired by',
    tokenIsRevoked: 'token is revoked',
    tokenNotFound: 'token not found',
    incorrectPassword:'incorrect password',
    osbbAlreadyExist:'idx_osbbs_edrpou',
    phoneNumberAlreadyExist:'user with this number already exist',
}

const HandleError = ({errorMsg,  func, argument, navigate}:any) => {
    if(errorMsg.includes(errorsMessages.tokenExpired)){
            auth.RefreshToken().then(func(argument))
    }else if(errorMsg.includes(errorsMessages.tokenIsRevoked) || errorMsg.includes(errorsMessages.tokenNotFound)) {
        localStorage.removeItem('token')
        navigate("/login");
    }else {
        console.log(errorMsg);
    }
}

export default {errorsMessages, HandleError};