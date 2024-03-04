import React from "react";

const OSBBForm = () =>{
    return(
        <form className='form' method='post'>
            <label form={'first-name'}>
                Ім'я
            </label>
            <input maxLength={256} name="first-name" placeholder="" type='text' id='first-name'/>
            <label form={'surname'}>
                Прізвище
            </label>
            <input maxLength={256} name="surname" placeholder="" type='text' id='surname'/>
            <label form={'patronymic'}>
                По батькові
            </label>
            <input maxLength={256} name="patronymic" placeholder="" type='text' id='patronymic'/>
            <label form={'password'}>
                Пароль
            </label>
            <input name="password" placeholder="" type='password' id='password'/>
            <label form={'phone-number'}>
                Номер телефону
            </label>
            <input name="phone-number" placeholder="" type='tel' id='phone-number'/>
            <label form={'osbb-name'}>
                Назва ОСББ
            </label>
            <input maxLength={256} name="osbb-name" placeholder="" type='text' id='osbb-name'/>
            <label form={'edrpou'}>
                ЕДРПОУ
            </label>
            <input name="edrpou" placeholder="" type='text' id='edrpou'/>
            <label form={'address'}>
                Адреса
            </label>
            <input maxLength={256} name="address" placeholder="" type='text' id='address'/>
            <label form={'rent'}>
                Плата за м^2
            </label>
            <input name="rent" placeholder="" type='number' id='rent'/>
            <input placeholder="" type='submit' id='rent' value={'Submit'}/>
        </form>
    )
}

export default OSBBForm