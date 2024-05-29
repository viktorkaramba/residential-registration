import React, {useCallback, useEffect, useState} from "react";
import config from "../../utils/config";
import {useAppContext} from "../../utils/AppContext";
import err from "../../utils/err";
import {useNavigate} from "react-router-dom";
import Checkbox from "@mui/material/Checkbox";
import logo from "../../images/photo.png";


const OSBBDescription = () => {
    // @ts-ignore
    const {osbbID, setOsbbID, user} = useAppContext();
    // @ts-ignore
    const {token} = useAppContext();
    const navigate = useNavigate();
    const [errorEDRPOU, setErrorEDRPOU] = useState(false);
    const [errorIBAN, setErrorIBAN] = useState(false);
    const [osbb, setOSBB] = useState<any>(null);
    const [isChecked, setIsChecked] = useState(false);
    const [newName, setNewName] = useState(osbb?.name);
    const [newEDRPOU, setNewEDRPOU] = useState(osbb?.edrpou);
    const [newIBAN, setNewIBAN] = useState(osbb?.iban);
    const [newRent, setNewRent] = useState(osbb?.rent);
    const [newPostAddress, setNewPostAddress] = useState(osbb?.building.Address);
    const [newAddress, setNewAddress] = useState(osbb?.building.Address);
    const [newPhoto, setNewPhoto] = useState(osbb?.photo);

    const fetchOSBB = useCallback(async() => {
        try{
            const requestOptions = {
                method: 'GET',
                headers:{ 'Content-Type': 'application/json',
                    'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            };
            fetch(config.apiUrl+'osbb/profile', requestOptions)
                .then(response => response.json())
                .then(data =>{

                    const {error}:any = data;
                    if(error){
                        err.HandleError({errorMsg:error, func:fetchOSBB, navigate:navigate});
                    }else {
                        if(data){
                            const {announcements, building, createdAt, edrpou, iban, id, name, osbb_head, rent, photo, updatedAt}:any = data;
                            const newOSBB = {
                                announcements: announcements,
                                building: building,
                                createdAt: createdAt,
                                edrpou: edrpou,
                                iban: iban,
                                id : id,
                                name: name,
                                osbb_head: osbb_head,
                                rent: rent,
                                photo:photo,
                                updatedAt: updatedAt
                            };
                            setOSBB(newOSBB);
                            setOsbbID(newOSBB.id);
                            setNewName(newOSBB.name);
                            setNewEDRPOU(newOSBB.edrpou);
                            setNewIBAN(newOSBB.iban);
                            setNewRent(newOSBB.rent);
                            setNewAddress(newOSBB.building.Address)
                            setNewPostAddress(newOSBB.building.Address)
                            setNewPhoto(newOSBB.photo)
                        }else {
                            setOSBB(null);
                        }
                    }
                });
        } catch(error){
            console.log(error);
        }
    }, [token]);
  
    useEffect(() => {
        fetchOSBB();
    }, []);

    function updateOSBBInfo({name, edrpou, iban, rent,address, photo}:any){
        setErrorEDRPOU(false)
        setErrorIBAN(false)
        let body = null;
        if(name != null){
            body = JSON.stringify({name: name});
        }
        if(edrpou != null){
            body = JSON.stringify({edrpou: parseInt(edrpou)});
        }
        if(iban != null){
            body = JSON.stringify({iban: iban});
        }
        if(rent != null){
            body = JSON.stringify({rent: parseFloat(rent)});
        }
        if(address != null){
            body = JSON.stringify({address: address});
        }
        if(photo != null){
            body = JSON.stringify({photo: photo});
        }
        console.log(body)
        const requestOptions = {
            method: 'PUT',
            headers:{ 'Content-Type': 'application/json',
                'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
            body: body,
        }

        fetch(config.apiUrl+'osbb/profile', requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data)
                const {error}:any = data;
                if(error){
                    if(error.includes(err.errorsMessages.osbbAlreadyExist)) {
                        setErrorEDRPOU(true);
                    }
                    else if(error.includes(err.errorsMessages.ibanAlreadyExist)) {
                        setErrorIBAN(true);
                    }
                    else {
                        err.HandleError({errorMsg:error, func:updateOSBBInfo, navigate:navigate});
                    }
                }else {
                    if(data){
                        setIsChecked(false);
                    }
                }
            });
    }

    return (
        <div className={'profile flex flex-c align-items-center'}>
            <div className={'flex bg-light-black'} style={{flexGrow:'1'}}>
            </div>
            <div className="card flex align-items-stretch flex-wrap">
                <div className="left-container flex flex-column align-items-center  align-self-center">
                    <img src={newPhoto !== null ? newPhoto: logo}
                         alt="OSBB Profile Image"/>
                    <h2>{newName}</h2>
                    <p>{newPostAddress}</p>
                </div>
                <div className="right-container flex flex-sb align-items-start">
                    <div className={'flex flex-column align-items-center  align-self-center'} style={{flexGrow:'1'}}>
                        <h3>Профіль ОСББ</h3>
                        {!isChecked && <table>
                            <tr>
                                <td>Назва ОСББ :</td>
                                <td>{newName}</td>
                            </tr>
                            <tr>
                                <td>ЕДРПОУ :</td>
                                <td>{newEDRPOU}</td>
                            </tr>
                            <tr>
                                <td>IBAN :</td>
                                <td>{newIBAN}</td>
                            </tr>
                            <tr>
                                <td>Ціна за м^2 :</td>
                                <td>{newRent} UAH</td>
                            </tr>
                            <tr>
                                <td>Адреса :</td>
                                <td>{newAddress}</td>
                            </tr>
                        </table>}
                        {isChecked && <table>
                            <tr>
                                <td>Назва ОСББ :</td>
                                <td className={'form'}>
                                    <div className="inner-wrap">
                                        <input maxLength={256}
                                               minLength={2}
                                               name="name_update_content"
                                               placeholder="Нове ім'я"
                                               type='text'
                                               onChange={e=>setNewName(e.target.value)}
                                               value={newName}
                                               id='name_update_content'/>
                                        <button className='button' onClick={()=>updateOSBBInfo({name:newName})}>
                                            <span className="button_content"> Оновити</span>
                                        </button>
                                    </div>
                                </td>
                            </tr>
                            <tr >
                                <td>ЕДРПОУ :</td>
                                <td className={'form'}>
                                    <div className="inner-wrap">
                                        <input
                                               name="edrpou_update_content"
                                               placeholder="Новий ЕДРПОУ"
                                               type='number'
                                               onChange={e=>setNewEDRPOU(e.target.value)}
                                               value={newEDRPOU}
                                               id='edrpou_update_content'/>
                                        {errorEDRPOU &&
                                            <div className={'error'}>
                                                ОСББ із таким ЕДРПОУ уже додане!
                                            </div>
                                        }
                                        <button className='button' onClick={()=>updateOSBBInfo({edrpou:newEDRPOU})}>
                                            <span className="button_content"> Оновити</span>
                                        </button>
                                    </div>
                                </td>
                            </tr>
                            <tr >
                                <td>IBAN :</td>
                                <td className={'form'}>
                                    <div className="inner-wrap">
                                        <input
                                            maxLength={256}
                                            minLength={2}
                                            name="iban_update_content"
                                            placeholder="Новий IBAN"
                                            type='text'
                                            onChange={e=>setNewIBAN(e.target.value)}
                                            value={newIBAN}
                                            id='iban_update_content'/>
                                        {errorIBAN &&
                                            <div className={'error'}>
                                                ОСББ із таким IBAN уже додане!
                                            </div>
                                        }
                                        <button className='button' onClick={()=>updateOSBBInfo({iban:newIBAN})}>
                                            <span className="button_content"> Оновити</span>
                                        </button>
                                    </div>
                                </td>
                            </tr>
                            <tr >
                                <td>Ціна за м^2 :</td>
                                <td className={'form'}>
                                    <div className="inner-wrap">
                                        <input
                                               name="rent_update_content"
                                               type='number'
                                               placeholder="Нова ціна"
                                               onChange={e=>setNewRent(e.target.value)}
                                               value={newRent}
                                               id='rent_update_content'/>
                                        <button className='button' onClick={()=>updateOSBBInfo({rent:newRent})}>
                                            <span className="button_content"> Оновити</span>
                                        </button>
                                    </div>
                                </td>
                            </tr>
                            <tr >
                                <td>Адреса :</td>
                                <td className={'form'}>
                                    <div className="inner-wrap">
                                        <input maxLength={256}
                                               minLength={2}
                                               name="adress_update_content"
                                               type='text'
                                               placeholder="Нова адреса"
                                               onChange={e=>setNewAddress(e.target.value)}
                                               value={newAddress}
                                               id='adress_update_content'/>
                                        <button className='button' onClick={()=>updateOSBBInfo({address:newAddress})}>
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
                                        <button className='button' onClick={()=>updateOSBBInfo({photo:newPhoto})}>
                                            <span className="button_content"> Оновити</span>
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        </table>}
                    </div>
                    <div>
                        {user?.role === "osbb_head" &&  <Checkbox
                            name="profile_check_box"
                            id='profile_check_box'
                            checked={isChecked}
                            size="large"
                            style={{color:'var(--blue-color)'}}
                            onChange={()=>{setIsChecked(!isChecked)}}
                        />}
                    </div>
                </div>
            </div>
            <div className={'flex'} style={{flexGrow:'1'}}>
            </div>
        </div>
    )
}

export default OSBBDescription