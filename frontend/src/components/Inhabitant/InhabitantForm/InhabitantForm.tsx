import React, {useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import {IoEyeOffOutline, IoEyeOutline} from "react-icons/io5";
import {Stack} from "@mui/material";
import Alert from "@mui/material/Alert";

const InhabitantForm = () =>{
    // @ts-ignore
    const {osbbID} = useAppContext();
    const [errorPhoneNumber, setErrorPhoneNumber] = useState(false);
    const [errorPassword, setErrorPassword] = useState(false);
    const [errorUserAlreadyExist, setErrorUserAlreadyExist] = useState(false);
    const [errorWaitApprove, setWaitApprove] = useState(false);
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    // @ts-ignore
    const {token, setToken} = useAppContext();
    const [visiblePassword, setPasswordVisible] = useState(false);
    const [visibleConfirmPassword, setConfirmVisible] = useState(false);
    // @ts-ignore
    const {setIsLogin} = useAppContext();
    const navigate = useNavigate();
    const addInhabitant = (event: any) => {
        event.preventDefault();

        // 👇️ access input values using name prop
        const firstName = event.target.first_name.value;
        const surname = event.target.surname.value;
        const patronymic = event.target.patronymic.value;
        const photo =  event.target.photo.value;
        const password = event.target.password.value;
        const confirm_password = event.target.confirm_password.value;
        if(password!==confirm_password){
            setErrorPassword(true)
            return
        }else {
            setErrorPassword(false)
        }
        const phone_number = event.target.phone_number.value;
        const apartment_number = parseInt(event.target.apartment_number.value);
        const apartment_area = parseInt(event.target.apartment_area.value);
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ first_name: firstName, surname: surname, patronymic:patronymic,
                password: password, phone_number:phone_number, apartment_number:apartment_number,
                apartment_area:apartment_area, photo:photo
            })
        }
        fetch(config.apiUrl+'osbb/' + osbbID + '/inhabitants', requestOptions)
            .then(response => response.json())
            .then(data => {
                const {error}:any = data;
                if(error){
                    if(error.includes(err.errorsMessages.phoneNumberAlreadyExist)) {
                        setErrorPhoneNumber(true);
                    }else if(error.includes(err.errorsMessages.userWaitApprove)){
                        setWaitApprove(true);
                    }else if(error.includes(err.errorsMessages.userAlreadyExist)){
                        setErrorUserAlreadyExist(true);
                    } else {
                        err.HandleError({errorMsg:error, func:addInhabitant, navigate:navigate});
                    }
               }else {
                    if(data){
                        const {token}:any = data
                        setToken(token);
                    }
                }
            });
        // 👇️ clear all input values in the form
        event.target.reset();
    };

    useEffect(() => {
        if (token) {
            localStorage.setItem("token", token);
            setIsLogin(true);
            setIsLoggedIn(true);
        }
    }, [token]);

    return(
        <form method='post' onSubmit={addInhabitant}>
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
                        <label form={'photo'}>Фото
                            <input name="photo" placeholder="" type='url' id='photo'/>
                        </label>
                        {errorPhoneNumber &&
                            <div className={'error'}>
                                Користувач з таким номером телефона уже зареєстрований!
                            </div>
                        }
                    </div>
                </div>
                <div className="form">
                    <div className="section"><span>2</span>Інформація про квартиру</div>
                    <div className="inner-wrap">
                        <label form={'apartment_number'}>Номер квартири
                            <input name="apartment_number" min={1} placeholder="" type='number' id='apartment_number' required={true}/>
                        </label>
                        <label form={'apartment_area'}>Площа квартири
                            <input name="apartment_area" min={1} placeholder="" type='number' id='apartment_area' required={true}/>
                        </label>

                    </div>
                    <div className={'flex flex-c'}>
                        <button className='button add_osbb' type="submit" name="submit_osbb">
                            <span className="button_content add_osbb_content">Додати ОСББ</span>
                        </button>

                    </div>
                    {errorWaitApprove &&
                        <div className={'error login_error m-15 text-center'}>
                            Ви вже відправили запит на приєднання
                        </div>
                    }
                    {errorUserAlreadyExist &&
                        <div className={'error login_error m-15 text-center'}>
                            Ви вже приєднанні до ОСББ
                        </div>
                    }
                    {isLoggedIn &&  <Stack sx={{margin: '10px'}} spacing={2}>
                        <Alert variant={'filled'} severity="success" style={{fontSize:'15px'}}>Запита успішно надісланий, очікуйте підтвердження</Alert>
                    </Stack>}
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

export default InhabitantForm