import React, { useState, useEffect } from 'react';
import config from "../../utils/config";
import Header from "../../components/Header/Header";
import err from "../../utils/err";
import {useAppContext} from "../../utils/AppContext";
import {useNavigate} from "react-router-dom";
import './Login.css'

const LoginPage = () => {

    const [phone_number, setPhoneNumber] = useState('');
    const [password, setPassword] = useState('');
    const [errorIncorrectData, setIncorrectData] = useState(false);
    const [errorNotApproved] = useState(false);
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
                    setIncorrectData(true)
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
        <div className={'login'}>
            <div className={'box'}>
                <form method={'post'} onSubmit={handleLogin}>
                    <h2>Вхід</h2>
                    <div className={'inputBox'}>
                        <input
                            type="tel"
                            id="phoneNumber"
                            value={phone_number}
                            required={true}
                            name="phoneNumber"
                            onChange={handlePhoneNumberChange}
                        />
                        <span>Номер Телефону</span>
                        <i></i>
                    </div>
                    <div className={'inputBox'}>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            value={password}
                            required={true}
                            onChange={handlePasswordChange}
                        />
                        <span>Пароль</span>
                        <i></i>
                    </div>
                    <div className={'links'}>
                        <a href={"#"}>Забули Пароль</a>
                    </div>
                    {errorIncorrectData &&
                        <div className={'error login_error'}>
                            Невірний номер телефону або пароль!
                        </div>
                    }

                    <button className='button login_button' type="submit" name="submit_osbb">
                        <span className="button_content login_button_content">Увійти</span>
                    </button>
                </form>
            </div>
        </div>

    );
};

export default LoginPage;
