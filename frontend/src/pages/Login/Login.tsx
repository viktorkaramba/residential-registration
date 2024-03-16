import React, { useState, useEffect } from 'react';
import config from "../../config";
import Header from "../../components/Header/Header";
import err from "../../err";
import {useAppContext} from "../../AppContext";
import {useNavigate} from "react-router-dom";

const LoginPage = () => {

    const [phone_number, setPhoneNumber] = useState('');
    const [password, setPassword] = useState('');
    // @ts-ignore
    const {token, setToken} = useAppContext();
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    // @ts-ignore
    const {setIsLogin} = useAppContext();
    const navigate = useNavigate();
    const handleLogin = (event: any) => {
        event.preventDefault();
        fetch(config.apiUrl+'auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ phone_number, password }),
        })
            .then(response => response.json())
            .then(data =>{
                console.log(data)
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:handleLogin, navigate:navigate});
                }else {
                    if(data){
                        const { token } = data;
                        setToken(token);
                    }
                }
            });
    };

    useEffect(() => {
        if (token) {
            localStorage.setItem("token", token);
            setIsLogin(true);
            setIsLoggedIn(true);
        }
    }, [token]);

    useEffect(() => {
        if (isLoggedIn) {
            navigate('/osbbs/profile');
        }
    }, [isLoggedIn]);

    const handlePhoneNumberChange = (event: { target: { value: React.SetStateAction<string>; }; }) => {
        setPhoneNumber(event.target.value);
    };

    const handlePasswordChange = (event: { target: { value: React.SetStateAction<string>; }; }) => {
        setPassword(event.target.value);
    };

    return (
        <div>
            <Header/>
            <h1>Логін</h1>
            <form method={'post'} onSubmit={handleLogin}>
                <label form={"phoneNumber"}>Номер телефону:</label>
                <input
                    type="tel"
                    id="phoneNumber"
                    value={phone_number}
                    required={true}
                    name="phoneNumber"
                    onChange={handlePhoneNumberChange}
                />
                <label form={"password"}>Пароль:</label>
                <input
                    type="password"
                    id="password"
                    name="password"
                    value={password}
                    required={true}
                    onChange={handlePasswordChange}
                />
                <button type="submit">Увійти</button>
            </form>
        </div>
    );
};

export default LoginPage;
