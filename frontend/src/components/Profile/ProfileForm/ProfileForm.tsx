import React, {useCallback, useEffect, useState} from "react";
import config from "../../../utils/config";
import {useAppContext} from "../../../utils/AppContext";
import err from "../../../utils/err";
import {useNavigate} from "react-router-dom";

const Profile = ({profile_user}:any) => {
    // @ts-ignore
    const navigate = useNavigate();
    // @ts-ignore
    const {osbbID, setOsbbID} = useAppContext();

    const [isChecked, setIsChecked] = useState(false);
    const [newFirstName, setNewFirstName] = useState(profile_user?.full_name.first_name);
    const [newSurname, setNewSurname] = useState(profile_user?.full_name.surname);
    const [newPatronymic, setNewPatronymic] = useState(profile_user?.full_name.patronymic);
    const [newPhoneNumber, setNewPhoneNumber] = useState(profile_user?.phone_number);
    const [newApartmentNumber , setNewApartmentNumber ] = useState(profile_user?.apartment.number);
    const [newApartmentArea    , setNewApartmentArea    ] = useState(profile_user?.apartment.area);


    function updateUserInfo({apartment_number, apartment_area, first_name, surname, patronymic, phone_number}:any){

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
        const requestOptions = {
            method: 'PUT',
            headers:config.headers,
            body: body,
        }

        fetch(config.apiUrl+'osbb/'+osbbID+'/inhabitants', requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:updateUserInfo, navigate:navigate});
                }else {
                    if(data){
                        setIsChecked(false);
                    }
                }
            });
    }

    return (
        <ul className={'flex flex-wrap align-items-start bg-dark-grey form'}>
            <input
                type="checkbox"
                checked={isChecked}
                onChange={e=>setIsChecked(!isChecked)}
            />
            <li>
                <label>
                    {!isChecked && newFirstName}
                    {isChecked &&
                        <div>
                            <input maxLength={256}
                                   minLength={2}
                                   required={true}
                                   name="first_name_update_content"
                                   placeholder=""
                                   type='text'
                                   onChange={e=>setNewFirstName(e.target.value)}
                                   value={newFirstName}
                                   id='first_name_update_content'/>
                            <button onClick={()=>updateUserInfo({first_name: newFirstName})}>Оновити</button>
                        </div>
                    }
                </label>
            </li>
            <li>
                <label>
                    {!isChecked && newSurname}
                    {isChecked &&
                        <div>
                            <input maxLength={256}
                                   minLength={2}
                                   required={true}
                                   name="surname_update_content"
                                   placeholder=""
                                   type='text'
                                   onChange={e=>setNewSurname(e.target.value)}
                                   value={newSurname}
                                   id='surname_update_content'/>
                            <button onClick={()=>updateUserInfo({surname:newSurname})}>Оновити</button>
                        </div>
                    }
                </label>
            </li>
            <li>
                <label>
                    {!isChecked && newPatronymic}
                    {isChecked &&
                        <div>
                            <input maxLength={256}
                                   minLength={2}
                                   required={true}
                                   name="patronymic_update_content"
                                   placeholder=""
                                   type='text'
                                   onChange={e=>setNewPatronymic(e.target.value)}
                                   value={newPatronymic}
                                   id='patronymic_update_content'/>
                            <button onClick={()=>updateUserInfo({patronymic:newPatronymic})}>Оновити</button>
                        </div>
                    }
                </label>
            </li>
            <li>
                <label>
                    {!isChecked && newPhoneNumber}
                    {isChecked &&
                        <div>
                            <input
                                required={true}
                                name="phone_number_update_content"
                                placeholder=""
                                type='tel'
                                onChange={e=>setNewPhoneNumber(e.target.value)}
                                value={newPhoneNumber}
                                id='phone_number_update_content'/>
                            <button onClick={()=>updateUserInfo({phone_number:newPhoneNumber})}>Оновити</button>
                        </div>
                    }
                </label>
            </li>
            <li>
                <label>
                    {!isChecked && newApartmentNumber}
                    {isChecked &&
                        <div>
                            <input
                                required={true}
                                name="apartment_number_update_content"
                                placeholder=""
                                type='number'
                                onChange={e=>setNewApartmentNumber(e.target.value)}
                                value={newApartmentNumber}
                                id='apartment_number_update_content'/>
                            <button onClick={()=>updateUserInfo({apartment_number:newApartmentNumber})}>Оновити</button>
                        </div>
                    }
                </label>
            </li>
            <li>
                <label>
                    {!isChecked && newApartmentArea}
                    {isChecked &&
                        <div>
                            <input
                                required={true}
                                name="apartment_area_update_content"
                                placeholder=""
                                type='number'
                                onChange={e=>setNewApartmentArea(e.target.value)}
                                value={newApartmentArea}
                                id='apartment_area_update_content'/>
                            <button onClick={()=>updateUserInfo({apartment_area:newApartmentArea})}>Оновити</button>
                        </div>
                    }
                </label>
            </li>
        </ul>
    )
}

export default Profile