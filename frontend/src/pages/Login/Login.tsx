import React, { useState, useEffect } from 'react';
import config from "../../config";

const LoginPage = () => {
    const [phone_number, setPhoneNumber] = useState('');
    const [password, setPassword] = useState('');
    const [token, setToken] = useState('');

    const handleLogin = async () => {
        // В цьому місці ви можете використати вашу логіку для відправки запиту на сервер для перевірки номера телефону та паролю
        // Наприклад, використовуючи fetch або axios
        try {
            const response = await fetch(config.apiUrl+'auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ phone_number, password }),
            });
            const data = await response.json();
            const { token } = data;
            setToken(token);
        } catch (error) {
            console.error('Помилка під час логіну:', error);
        }
    };

    useEffect(() => {

        if (token) {
            localStorage.setItem('accessToken', token);
        }
    }, [token]);

    const handlePhoneNumberChange = (event: { target: { value: React.SetStateAction<string>; }; }) => {
        setPhoneNumber(event.target.value);
    };

    const handlePasswordChange = (event: { target: { value: React.SetStateAction<string>; }; }) => {
        setPassword(event.target.value);
    };

    return (
        <div>
            <h1>Логін</h1>
            <form onSubmit={(e) => {
                e.preventDefault();
                handleLogin();
            }}>
                <div>
                    <label htmlFor="phoneNumber">Номер телефону:</label>
                    <input
                        type="text"
                        id="phoneNumber"
                        value={phone_number}
                        onChange={handlePhoneNumberChange}
                    />
                </div>
                <div>
                    <label htmlFor="password">Пароль:</label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={handlePasswordChange}
                    />
                </div>
                <button type="submit">Увійти</button>
            </form>
        </div>
    );
};

export default LoginPage;