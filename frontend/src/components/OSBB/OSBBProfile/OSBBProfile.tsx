import React from "react";
import {useOSBBContext} from "../OSBBContext";
import Header from "../../Header/Header";

const OSBBProfile = () => {

    // @ts-ignore
    const {osbbID} = useOSBBContext();
    return(
        <div>
            <Header/>
            {osbbID}
        </div>
    )
}

export default OSBBProfile