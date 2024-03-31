import React from "react";
import {useAppContext} from "../../../utils/AppContext";
import "../OSBB.css"
import logo from "../../../images/photo.png";
const OSBBListElement = ((osbb:any) => {

    // @ts-ignore
    const {isLogin, setOsbbID, setActiveOSBBElement, user} = useAppContext();
    const handleConnectOSBB = (id:number, element: React.SetStateAction<string>) => {
        setActiveOSBBElement(element);
        setOsbbID(id);
    };
    return(
        <div className='osbb-item flex flex-column flex-sb flex-wrap'>
            <div className='osbb-item-info text-center'>
                <div className='osbb-item-img'>
                    <img src={osbb.photo !== undefined ? osbb.photo: logo}
                         alt="OSBB Profile Image"/>
                </div>
                <div className='osbb-item-info-item fw-7 fs-18'>
                    <span>{osbb.name}</span>
                </div>
                <div className='osbb-item-info-item fs-15'>
                    <span className='text-capitalize fw-7'>Голова: </span>
                    <span className={'m-5'}>{osbb.osbb_head.full_name.first_name}</span>
                    <span className={'m-5'}>{osbb.osbb_head.full_name.surname}</span>
                    <span className={'m-5'}>{osbb.osbb_head.full_name.patronymic}</span>
                </div>
                <div className='osbb-item-info-item fs-15'>
                    <span className='text-capitalize fw-7'>Адреса: </span>
                    <span>{osbb.building.Address}</span>
                </div>
            </div>
            {(!isLogin || !user.is_approved) &&
                <button className='button' onClick={() => handleConnectOSBB(osbb.id, 'InhabitantForm')}>
                    <span className="button_content">Приєднатися до ОСББ</span>
            </button>}
        </div>
        )
});

export default OSBBListElement