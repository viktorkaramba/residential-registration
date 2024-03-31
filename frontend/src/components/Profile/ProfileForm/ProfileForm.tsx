import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";
import './ProfileFrom.css'
import Checkbox from "@mui/material/Checkbox";
import ApartmentForm from "../../Apartments/ApartmentForm";

const Profile = ({profile_user}:any) => {
    // @ts-ignore
    const navigate = useNavigate();
    // @ts-ignore
    const {osbbID} = useAppContext();
    const [errorPhoneNumber, setErrorPhoneNumber] = useState(false);

    const [isChecked, setIsChecked] = useState(false);
    const [isAddApartment, setIsAddApartment] = useState(false);
    const [newFirstName, setNewFirstName] = useState(profile_user?.full_name.first_name);
    const [newSurname, setNewSurname] = useState(profile_user?.full_name.surname);
    const [newPatronymic, setNewPatronymic] = useState(profile_user?.full_name.patronymic);
    const [newPhoneNumber, setNewPhoneNumber] = useState(profile_user?.phone_number);
    const [newPhoto, setNewPhoto] = useState(profile_user?.photo);
    const [newApartmentNumber , setNewApartmentNumber ] = useState(profile_user?.apartment.number);
    const [newApartmentArea    , setNewApartmentArea    ] = useState(profile_user?.apartment.area);

    function addApartment (apartment: any){
        setNewApartmentNumber(apartment.number)
        setNewApartmentArea(apartment.area)
        setIsAddApartment(false)
    }

    function updateUserInfo({apartment_number, apartment_area, first_name, surname, patronymic, phone_number, photo}:any){

        let body = null;
        if(apartment_number != null){
            body = JSON.stringify({apartment_number: apartment_number});
        }
        if(apartment_area != null){
            body = JSON.stringify({apartment_area: apartment_area});
        }
        if(first_name != null){
            body = JSON.stringify({first_name: first_name});
        }
        if(surname != null){
            body = JSON.stringify({surname: surname});
        }
        if(patronymic != null){
            body = JSON.stringify({patronymic: patronymic});
        }
        if(phone_number != null){
            body = JSON.stringify({phone_number: phone_number});
        }
        if(photo != null){
            body = JSON.stringify({photo: photo});
        }
        const requestOptions = {
            method: 'PUT',
            headers:{ 'Content-Type': 'application/json',
                'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            body: body,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/inhabitants', requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                const {error}:any = data;
                if(error){
                    if(error.includes(err.errorsMessages.phoneNumberAlreadyExist)) {
                        setErrorPhoneNumber(true);
                    }else {
                        err.HandleError({errorMsg:error, func:updateUserInfo, navigate:navigate});
                    }
                }else {
                    if(data){
                        setIsChecked(false);
                    }
                }
            });
    }

    function handle(){
        console.log('add')
        setIsAddApartment(true)
    }
    return (
        <div className={'profile flex flex-c align-items-center'}>
         <div className={'flex bg-light-black'} style={{flexGrow:'1'}}>
         </div>
        <div className="card flex align-items-stretch flex-wrap">
            <div className="left-container flex flex-column align-self-center">
                <img src={newPhoto}
                     alt="Profile Image"/>
                    <h2>{newFirstName} {newSurname} {newPatronymic}</h2>
                {profile_user?.role === "osbb_head" &&   <p>Голова ОСББ</p>}
                {profile_user?.role === "inhabitant" &&   <p>Мешканець</p>}
            </div>
            <div className="right-container flex flex-sb align-items-start">
                <div className={'flex flex-column align-items-center align-self-center'} style={{flexGrow:'1'}}>
                    <h3>Профіль</h3>
                    {!isChecked && <table>
                        <tr>
                            <td>ПІБ :</td>
                            <td>{newFirstName} {newSurname} {newPatronymic}</td>
                        </tr>
                        <tr>
                            <td>Номер телефону :</td>
                            <td>{newPhoneNumber}</td>
                        </tr>
                        {newApartmentNumber !== 0 && <>
                            <tr>
                                <td>Номер квартири :</td>
                                <td>{newApartmentNumber}</td>
                            </tr>
                            <tr>
                                <td>Площа квартири :</td>
                                <td>{newApartmentArea}</td>
                            </tr></>}
                        {newApartmentNumber === 0 && <tr>
                            <td>
                                {!isAddApartment && <button className='button add_apartment' onClick={()=>handle()}>
                                    <span className="button_content add_apartment_content">Додати квартиру</span>
                                </button>}
                                {isAddApartment && <button className='button add_apartment' onClick={()=>setIsAddApartment(false)}>
                                    <span className="button_content add_apartment_content">Закрити</span>
                                </button>}
                            </td>
                        </tr>}
                        <tr>
                            <td>
                            </td>
                        </tr>
                    </table>}
                    {isAddApartment && <ApartmentForm addApartment={addApartment}/>}
                    {isChecked && <table>
                        <tr>
                            <td>ПІБ :</td>
                            <td className={'form'}>
                                <div className="inner-wrap">
                                    <input maxLength={256}
                                           minLength={2}
                                           name="first_name_update_content"
                                           placeholder=""
                                           type='text'
                                           onChange={e=>setNewFirstName(e.target.value)}
                                           value={newFirstName}
                                           id='first_name_update_content'/>
                                    <input maxLength={256}
                                           minLength={2}
                                           name="surname_update_content"
                                           placeholder=""
                                           type='text'
                                           onChange={e=>setNewSurname(e.target.value)}
                                           value={newSurname}
                                           id='surname_update_content'/>
                                    <input maxLength={256}
                                           minLength={2}
                                           name="patronymic_update_content"
                                           placeholder=""
                                           type='text'
                                           onChange={e=>setNewPatronymic(e.target.value)}
                                           value={newPatronymic}
                                           id='patronymic_update_content'/>
                                    <button className='button' onClick={()=>updateUserInfo({first_name: newFirstName, surname: newSurname, patronymic:newPatronymic})}>
                                        <span className="button_content"> Оновити</span>
                                    </button>
                                </div>
                            </td>
                        </tr>
                        <tr >
                            <td>Номер телефону :</td>
                            <td className={'form'}>
                                <div className="inner-wrap">
                                    <input
                                        name="phone_number_update_content"
                                        placeholder=""
                                        type='tel'
                                        onChange={e=>setNewPhoneNumber(e.target.value)}
                                        value={newPhoneNumber}
                                        id='phone_number_update_content'/>
                                    {errorPhoneNumber &&
                                        <div className={'error m-5'}>
                                            Користувач з таким номером телефона уже зареєстрований!
                                        </div>
                                    }
                                    <button className='button' onClick={()=>updateUserInfo({phone_number:newPhoneNumber})}>
                                        <span className="button_content"> Оновити</span>
                                    </button>
                                </div>
                            </td>
                        </tr>
                        <tr >
                            <td>Фото :</td>
                            <td className={'form'}>
                                <div className="inner-wrap">
                                    <input name="photo_update_content"
                                           type='url'
                                           placeholder="Нове фото"
                                           onChange={e=>setNewPhoto(e.target.value)}
                                           value={newPhoto}
                                           id='photo_update_content'/>
                                    <button className='button' onClick={()=>updateUserInfo({photo:newPhoto})}>
                                        <span className="button_content"> Оновити</span>
                                    </button>
                                </div>
                            </td>
                        </tr>
                        {newApartmentNumber !== 0 && <>
                            <tr>
                                <td>Номер квартири :</td>
                                <td className={'form'}>
                                    <div className="inner-wrap">
                                        <input
                                            name="apartment_number_update_content"
                                            placeholder=""
                                            type='number'
                                            onChange={e=>setNewApartmentNumber(e.target.value)}
                                            value={newApartmentNumber}
                                            id='apartment_number_update_content'/>
                                        <button className='button' onClick={()=>updateUserInfo({apartment_number:newApartmentNumber})}>
                                            <span className="button_content"> Оновити</span>
                                        </button>
                                    </div>
                                </td>
                            </tr>
                            <tr>
                                <td>Площа квартири :</td>
                                <td className={'form'}>
                                    <div className="inner-wrap">
                                        <input
                                            name="apartment_area_update_content"
                                            placeholder=""
                                            type='number'
                                            onChange={e=>setNewApartmentArea(e.target.value)}
                                            value={newApartmentArea}
                                            id='apartment_area_update_content'/>
                                        <button className='button' onClick={()=>updateUserInfo({apartment_area:newApartmentArea})}>
                                            <span className="button_content"> Оновити</span>
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        </>}
                        {newApartmentNumber === 0 && <tr>
                            <td>
                                <button className='button add_apartment'>
                                    <span className="button_content add_apartment_content">Додати квартиру</span>
                                </button>
                            </td>
                        </tr>}
                    </table>}
                </div>
                <div>
                    <Checkbox
                        name="profile_check_box"
                        id='profile_check_box'
                        checked={isChecked}
                        size="large"
                        style={{color:'var(--blue-color)'}}
                        onChange={()=>{setIsChecked(!isChecked)}}
                    />
                </div>
            </div>
        </div>
            <div className={'flex'} style={{flexGrow:'1'}}>
            </div>
        </div>
    )
}

export default Profile