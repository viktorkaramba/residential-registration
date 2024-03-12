import RefreshToken from "./auth";
import {useNavigate} from "react-router-dom";
import React from "react";

const errorsMessages ={
    tokenExpired :'token is expired by',
    tokenIsRevoked: 'token is revoked',
    incorrectPassword:'incorrect password',
    osbbAlreadyExist:'idx_osbbs_edrpou',
    phoneNumberAlreadyExist:'user with this number already exist',
}

const HandleError = ({errorMsg,  func, argument, navigate}:any) => {
    if(errorMsg.includes(errorsMessages.tokenExpired)){
            RefreshToken().then(func(argument))
    }else if(errorMsg.includes(errorsMessages.tokenIsRevoked)) {
        navigate("/login");
    }else {
        console.log(errorMsg);
    }
}

export default {errorsMessages, HandleError};