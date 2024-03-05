import React from "react";
import {useOSBBContext} from "../OSBBContext";

const OSBBListElement = ((osbb:any) => {

    // @ts-ignore
    const {setOsbbID, setActiveOSBBElement} = useOSBBContext();
    const handleConnectOSBB = (id:number, element: React.SetStateAction<string>) => {
        setActiveOSBBElement(element);
        setOsbbID(id);
    };
    return(
        <div>
            <br></br>
            {osbb.name}
            <br></br>
            {osbb.osbb_head.phone_number}
            <br></br>
            <div className='text-block' onClick={() => handleConnectOSBB(osbb.id, '3')}>
                Приєднатися до ОСББ
            </div>
        </div>
        )
});

export default OSBBListElement