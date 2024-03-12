import config from "./config";

function RefreshToken():any{
    let oldToken = localStorage.getItem('token') || '{}'
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ token:oldToken })
    }
    return fetch(config.apiUrl+'refresh-token', requestOptions)
        .then(response => response.json())
        .then(data => {
            const {token}:any = data
            localStorage.setItem('token', token);
        });
}

export default RefreshToken