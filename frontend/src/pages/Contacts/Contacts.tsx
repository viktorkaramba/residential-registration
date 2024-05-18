import React, {useState, useEffect, useRef} from 'react';
import Header from "../../components/Header/Header";
import './Contacts.css'
const ContactsPage = () => {


    return (
        <main>
            <Header withWelcomeBlock={false}/>
                <div className="flex m-15">
                    <div className="container">
                        <div className="row">
                            <div className="col-lg-6 d-flex align-items-center">
                                <div className="contact-info">
                                    <h2 className="contact-title">Є будь-які запитання?</h2>
                                    <p>Якщо у вас виникли будь-які запитання або вам потрібна додаткова інформація, будь ласка,
                                        звертайтеся до нас за наведеними контактними даними. Ми завжди раді допомогти!</p>
                                    <ul className="contact-info">
                                        <li>
                                            <div className="info-left">
                                                <i className="fas fa-mobile-alt"></i>
                                            </div>
                                            <div className="info-right">
                                                <h4>+380961533469</h4>
                                            </div>
                                        </li>
                                        <li>
                                            <div className="info-left">
                                                <i className="fas fa-at"></i>
                                            </div>
                                            <div className="info-right">
                                                <h4>osbbonline@gmail.com</h4>
                                            </div>
                                        </li>
                                        <li>
                                            <div className="info-left">
                                                <i className="fas fa-map-marker-alt"></i>
                                            </div>
                                            <div className="info-right">
                                                <h4>проспект Незалежності, 27, Нетішин, Хмельницька область, 30100</h4>
                                            </div>
                                        </li>
                                    </ul>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
        </main>
    )
}

export default ContactsPage;
