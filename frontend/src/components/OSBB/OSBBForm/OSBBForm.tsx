import React, {useEffect, useRef, useState} from "react";
import config from "../../../utils/config";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import {useAppContext} from "../../../utils/AppContext";
import {IoEyeOutline, IoEyeOffOutline} from "react-icons/io5";
import "./OSBBForm.css"
import "../OSBB.css"

const OSBBForm = () =>{
    const [errorPhoneNumber, setErrorPhoneNumber] = useState(false);
    const [errorEDRPOU, setErrorEDRPOU] = useState(false);
    const [errorPassword, setErrorPassword] = useState(false);
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [visiblePassword, setPasswordVisible] = useState(false);
    const [visibleConfirmPassword, setConfirmVisible] = useState(false);
    const focusRef = useRef(null);
    // @ts-ignore
    const {token, setToken} = useAppContext();
    // @ts-ignore
    const {setIsLogin} = useAppContext();
    const navigate = useNavigate();
    const addOSBB = (event: any) => {

        event.preventDefault();
        setErrorEDRPOU(false);
        setErrorPhoneNumber(false);
        // 👇️ access input values using name prop
        const firstName = event.target.first_name.value;
        const surname = event.target.surname.value;
        const patronymic = event.target.patronymic.value;
        const password = event.target.password.value;
        const confirm_password = event.target.confirm_password.value;
        if(password!==confirm_password){
            setErrorPassword(true)
            return
        }else {
            setErrorPassword(false)
        }
        const phone_number = event.target.phone_number.value;
        const name = event.target.name.value;
        const edrpou = parseInt(event.target.edrpou.value);
        const address = event.target.address.value;
        const rent = parseFloat(event.target.rent.value);
        const photo =  event.target.photo.value;
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ first_name: firstName, surname: surname, patronymic:patronymic,
                password: password, phone_number:phone_number, name:name,
                edrpou:edrpou, address:address, rent:rent, photo:photo,
            })
        }
        fetch(config.apiUrl+'osbb/', requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    if(error.includes(err.errorsMessages.osbbAlreadyExist)){
                        setErrorEDRPOU(true);
                    }else if(error.includes(err.errorsMessages.phoneNumberAlreadyExist)) {
                        setErrorPhoneNumber(true);
                    }else {
                        err.HandleError({errorMsg:error, func:addOSBB, navigate:navigate});
                    }
                }else {
                    if(data){
                        const {token}:any = data
                        setToken(token);
                        event.target.reset();
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
        // @ts-ignore
        focusRef.current.scrollIntoView({ behavior: 'smooth' });
    }, []);

    useEffect(() => {
        if (isLoggedIn) {
            navigate('/osbbs/profile');
        }
    }, [isLoggedIn]);

    return(
        <form method='post'  onSubmit={addOSBB} ref={focusRef}>
            <div className={'flex flex-wrap align-items-start bg-dark-grey'}>
                <div className="form">
                    <div className="section"><span>1</span>ПІБ та Номер Телефону</div>
                    <div className="inner-wrap">
                        <label form={'first_name'}>Ім'я
                            <input maxLength={256} minLength={2} required={true} name="first_name" placeholder="" type='text' id='first_name'/>
                        </label>
                        <label  form={'surname'}>Прізвище
                            <input maxLength={256} minLength={2} required={true} name="surname" placeholder="" type='text' id='surname'/>
                        </label>
                        <label form={'patronymic'}>По Батькові
                            <input maxLength={256} minLength={2} required={true} name="patronymic" placeholder="" type='text' id='patronymic'/>
                        </label>
                        <label form={'phone_number'}>Номер Телефону
                            <input name="phone_number" required={true} placeholder="" type='tel' id='phone_number'/>
                        </label>
                        {errorPhoneNumber &&
                            <div className={'error'}>
                                Користувач з таким номером телефона уже зареєстрований!
                            </div>
                        }
                    </div>
                </div>
                <div className="form">
                    <div className="section"><span>2</span>Інформація про ОСББ</div>
                    <div className="inner-wrap">
                        <label form={'name'}>Назва ОСББ
                            <input maxLength={256} minLength={2} required={true} name="name" placeholder="" type='text' id='name'/>
                        </label>
                        <label form={'edrpou'}>ЕДРПОУ
                            <input name="edrpou" required={true} placeholder="" type='number' id='edrpou'/>
                        </label>
                        <label form={'address'}>Адреса
                            <input maxLength={256} minLength={2} required={true} name="address" placeholder="" type='text' id='address'/>
                        </label>
                        <label form={'rent'}>Плата за м^2
                            <input name="rent" required={true} placeholder="" type='number' id='rent'/>
                        </label>
                        <label form={'photo'}>Фото
                            <input name="photo" placeholder="" type='url' id='photo'/>
                        </label>
                        {errorEDRPOU &&
                            <div className={'error'}>
                                ОСББ із таким ЕДРПОУ уже додане!
                            </div>
                        }
                    </div>

                    <div className={'flex flex-c'}>
                        <button className='button add_osbb' type="submit" name="submit_osbb">
                            <span className="button_content add_osbb_content">Додати ОСББ</span>
                        </button>
                    </div>
                </div>
                <div className="form">
                    <div className="section"><span>3</span>Пароль</div>
                    <div className="inner-wrap">
                        <label form={'password'}>Пароль
                            <div className={'flex'}>
                                <input name="password" minLength={8} required={true} placeholder="" type={visiblePassword ? "text": "password"} id='password'/>
                                <div className={'p-5'} onClick={()=>setPasswordVisible(!visiblePassword)}>
                                    {visiblePassword ? <IoEyeOutline/>:<IoEyeOffOutline/>}
                                </div>
                            </div>
                        </label>
                        <label form={'confirm_password'}>Підтвердження пароля
                            <div className={'flex'}>
                                <input name="confirm_password" minLength={8} required={true} placeholder="" type={visibleConfirmPassword ? "text": "password"} id='confirm_password'/>
                                <div className={'p-5'} onClick={()=>setConfirmVisible(!visibleConfirmPassword)}>
                                    {visibleConfirmPassword ? <IoEyeOutline/>:<IoEyeOffOutline/>}
                                </div>
                            </div>
                        </label>
                        {errorPassword &&
                            <div className={'error'}>
                                Паролі не співпадають!
                            </div>
                        }
                    </div>
                </div>
            </div>
        </form>
    )
}

export default OSBBForm