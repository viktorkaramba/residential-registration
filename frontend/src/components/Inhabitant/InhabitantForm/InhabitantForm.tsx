import React from "react";
import config from "../../../config";
import {useOSBBContext} from "../../OSBB/OSBBContext";

const InhabitantForm = () =>{
    // @ts-ignore
    const {osbbID} = useOSBBContext();
    const handleSubmit = (event: any) => {
        console.log('handleSubmit ran');
        event.preventDefault();

        // 👇️ access input values using name prop
        const firstName = event.target.first_name.value;
        const surname = event.target.surname.value;
        const patronymic = event.target.patronymic.value;
        const password = event.target.password.value;
        const phone_number = event.target.phone_number.value;
        const apartment_number = parseInt(event.target.apartment_number.value);
        const apartment_area = parseInt(event.target.apartment_area.value);
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ first_name: firstName, surname: surname, patronymic:patronymic,
                password: password, phone_number:phone_number, apartment_number:apartment_number,
                apartment_area:apartment_area
            })
        }
        fetch(config.apiUrl+'osbb/' + osbbID + '/inhabitants', requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
            });
        // 👇️ clear all input values in the form
        event.target.reset();
    };

    return(
        <form className='form' method='post' onSubmit={handleSubmit}>
            <label form={'first_name'}>
                Ім'я
            </label>
            <input maxLength={256} minLength={2} name="first_name" placeholder="" type='text' id='first_name' required={true}/>
            <label form={'surname'}>
                Прізвище
            </label>
            <input maxLength={256} minLength={2} name="surname" placeholder="" type='text' id='surname' required={true}/>
            <label form={'patronymic'}>
                По батькові
            </label>
            <input maxLength={256} minLength={2} name="patronymic" placeholder="" type='text' id='patronymic' required={true}/>
            <label form={'password'}>
                Пароль
            </label>
            <input name="password" minLength={8} placeholder="" type='password' id='password' required={true}/>
            <label form={'phone_number'}>
                Номер телефону
            </label>
            <input name="phone_number" placeholder="" type='tel' id='phone_number' required={true}/>
            <label form={'apartment_number'}>
                Номер квартири
            </label>
            <input name="apartment_number" min={1} placeholder="" type='number' id='apartment_number' required={true}/>
            <label form={'apartment_area'}>
                Площа квартири
            </label>
            <input name="apartment_area" min={1} placeholder="" type='number' id='apartment_area' required={true}/>
            <button type="submit">Submit form</button>
        </form>
    )
}

export default InhabitantForm