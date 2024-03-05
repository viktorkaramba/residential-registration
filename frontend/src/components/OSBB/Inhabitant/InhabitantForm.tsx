import React from "react";

const InhabitantForm = () =>{
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
            <label form={'apartment-number'}>
                Номер квартири
            </label>
            <input name="apartment-number" placeholder="" type='number' id='apartment-number'/>
            <label form={'apartment-area'}>
                Площа квартири
            </label>
            <input name="apartment-area" placeholder="" type='number' id='apartment-area'/>
            <label form={'address'}>
                Адреса
            </label>
            <input placeholder="" type='submit' id='rent' value={'Submit'}/>
        </form>
    )
}

export default InhabitantForm