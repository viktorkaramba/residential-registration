import React, {useState, useEffect, useRef} from 'react';
import config from "../../utils/config";
import Header from "../../components/Header/Header";
import err from "../../utils/err";
import {useAppContext} from "../../utils/AppContext";
import {useNavigate} from "react-router-dom";
import './Login.css'
import {IoEyeOffOutline, IoEyeOutline} from "react-icons/io5";

const LoginPage = () => {

    const [phone_number, setPhoneNumber] = useState('');
    const [password, setPassword] = useState('');
    const [errorNotFound, setNotFound] = useState(false);
    const [errorIncorrectPassword, setIncorrectPassword] = useState(false);
    const [errorNotApproved, setNotApproved] = useState(false);
    const [errorWaitApprove, setWaitApprove] = useState(false);
    const [visiblePassword, setPasswordVisible] = useState(false);

    // @ts-ignore
    const {token, setToken, prevPage} = useAppContext();
    const inputRef = useRef<HTMLInputElement | null>(null);
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
                    if(error.includes(err.errorsMessages.incorrectPassword)){
                        setIncorrectPassword(true);
                    }else if(error.includes(err.errorsMessages.userNotApproved)) {
                        setNotApproved(true);
                    }else if(error.includes(err.errorsMessages.userNotFound)){
                        setNotFound(true)
                    }else if(error.includes(err.errorsMessages.userWaitApprove)){
                        setWaitApprove(true)
                    } else {
                        err.HandleError({errorMsg:error, func:handleLogin, navigate:navigate});
                    }
                }else {
                    if(data){
                        const { token } = data;
                        setToken(token);
                    }
                }
            });
    };
    function handleHide(){
        if (inputRef.current) {
            inputRef.current.focus();
            const length = inputRef.current.value.length;
            requestAnimationFrame(() => {
                inputRef.current?.setSelectionRange(length, length);
            });
        }
        setPasswordVisible(!visiblePassword);
    }

    useEffect(() => {
        if (token) {
            localStorage.setItem("token", token);
            setIsLogin(true);
            setIsLoggedIn(true);
        }
    }, [token]);

    useEffect(() => {
        if (isLoggedIn) {
            navigate(prevPage);
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
                        <div className={'flex'}>
                            <input
                                type={visiblePassword ? "text": "password"}
                                id="password"
                                name="password"
                                value={password}
                                ref={inputRef}
                                required={true}
                                onChange={handlePasswordChange}
                            />
                            <span>Пароль</span>
                            <i id={"platform"}></i>
                            <div className={'p-15 flex flex-end align-items-end align-self-end hide_password'} style={{zIndex:2}} onClick={()=>handleHide()}>
                                {visiblePassword ? <IoEyeOutline className={'text-blue'}/>:<IoEyeOffOutline className={'text-blue'}/>}
                            </div>
                        </div>
                    </div>
                    <div className={'links'}>
                        <a href={"#"}>Забули Пароль</a>
                    </div>
                    {errorNotFound &&
                        <div className={'error login_error'}>
                            Невірний номер телефону
                        </div>
                    }
                    {errorIncorrectPassword &&
                        <div className={'error login_error'}>
                            Невірний пароль!
                        </div>
                    }
                    {errorWaitApprove &&
                        <div className={'error login_error'}>
                            Дочекайтись поки приймуть ваш запит
                        </div>
                    }
                    {errorNotApproved &&
                        <div className={'error login_error'}>
                            На жаль, ваш запит був відхилений
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
