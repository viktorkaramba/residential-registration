import React, {useState} from "react";
import {useNavigate} from "react-router-dom";
import Alert from "@mui/material/Alert";
import {Stack} from "@mui/material";
import config from "../../utils/config";
import err from "../../utils/err";
import {useAppContext} from "../../utils/AppContext";

const ApartmentForm = ({addApartment}:any) =>{
    // @ts-ignore
    const {osbbID} = useAppContext()
    const navigate = useNavigate();
    const handleAddApartment = (event: any) => {
        event.preventDefault();

        // 👇️ access input values using name prop
        const number = event.target.number.value;
        const area = event.target.area.value;
        const requestOptions = {
            method: 'POST',
            headers:{ 'Content-Type': 'application/json',
                'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            body: JSON.stringify({ number: number, area: area})
        }
        fetch(config.apiUrl+'osbb/'+osbbID+'/apartments', requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data)
                const {error}:any = data;
                if(error){
                    err.HandleError({errorMsg:error, func:handleAddApartment, navigate:navigate});
                }else {
                    const {id}:any = data
                    const newApartment = {
                        id: id,
                        number: number,
                        area: area,
                    };
                    addApartment(newApartment)
                }
            });
        // // 👇️ clear all input values in the form
        // event.target.reset();
    };

    return(
        <form method='post'  onSubmit={handleAddApartment}>
            <div className={'flex flex-wrap align-items-start bg-dark-grey m-5'}>
                <div className="form announcement_form">
                    <h1>Номер та площа квартири</h1>
                    <div className="inner-wrap">
                        <label form={'number'}>Номер квартири
                            <input required={true} name="number" placeholder="Номер" type='number' id='number'/>
                        </label>
                        <label form={'area'}>Площа квартири
                            <input required={true} name="area" placeholder="Площа" type='number' id='area'/>
                        </label>
                    </div>
                    <div className={'flex flex-c'}>
                        <button className='button add_osbb' type="submit" name="submit_osbb">
                            <span className="button_content add_osbb_content">Додати Квартиру</span>
                        </button>
                    </div>
                </div>
            </div>
        </form>
    )
}

export default ApartmentForm