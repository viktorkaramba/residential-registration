import Navbar from "./Navbar";
import React from "react";

const WelcomeBlock = () =>{
    return(
        <div className='header-content flex flex-c text-center text-black'>
            <h2 className='header-title text-capitalize'>Зручний сервіс для керування ОСББ</h2><br/>
            <p className='header-text fs-18 fw-3'>Додавайте ОСББ та комунікуйте з мешканцями за допомогою оголошень та опитувань, а також відстежуйте їхню активність</p>
        </div>
    )
}

export default WelcomeBlock