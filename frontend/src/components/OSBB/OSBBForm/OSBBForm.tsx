import React from "react";
import config from "../../../config";

const OSBBForm = () =>{
    const handleSubmit = (event: any) => {
        console.log('handleSubmit ran');
        event.preventDefault();

        // 👇️ access input values using name prop
        const firstName = event.target.first_name.value;
        const surname = event.target.surname.value;
        const patronymic = event.target.patronymic.value;
        const password = event.target.password.value;
        const phone_number = event.target.phone_number.value;
        const name = event.target.name.value;
        const edrpou = parseInt(event.target.edrpou.value);
        const address = event.target.address.value;
        const rent = parseFloat(event.target.rent.value);
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ first_name: firstName, surname: surname, patronymic:patronymic,
                password: password, phone_number:phone_number, name:name,
                edrpou:edrpou, address:address, rent:rent
            })
        }
        fetch(config.apiUrl+'osbb/', requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
            });
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };
    return(
        <form className='form' method='post'  onSubmit={handleSubmit}>
            <label form={'first_name'}>
                Ім'я
            </label>
            <input maxLength={256} minLength={2} required={true} name="first_name" placeholder="" type='text' id='first_name'/>
            <label form={'surname'}>
                Прізвище
            </label>
            <input maxLength={256} minLength={2} required={true} name="surname" placeholder="" type='text' id='surname'/>
            <label form={'patronymic'}>
                По батькові
            </label>
            <input maxLength={256} minLength={2} required={true} name="patronymic" placeholder="" type='text' id='patronymic'/>
            <label form={'password'}>
                Пароль
            </label>
            <input name="password" minLength={8} required={true} placeholder="" type='password' id='password'/>
            <label form={'phone_number'}>
                Номер телефону
            </label>
            <input name="phone_number" required={true} placeholder="" type='tel' id='phone_number'/>
            <label form={'name'}>
                Назва ОСББ
            </label>
            <input maxLength={256} minLength={2} required={true} name="name" placeholder="" type='text' id='name'/>
            <label form={'edrpou'}>
                ЕДРПОУ
            </label>
            <input name="edrpou" required={true} placeholder="" type='number' id='edrpou'/>
            <label form={'address'}>
                Адреса
            </label>
            <input maxLength={256} minLength={2} required={true} name="address" placeholder="" type='text' id='address'/>
            <label form={'rent'}>
                Плата за м^2
            </label>
            <input name="rent" required={true} placeholder="" type='text' id='rent'/>
            <button type="submit">Submit form</button>
        </form>
    )
}

export default OSBBForm