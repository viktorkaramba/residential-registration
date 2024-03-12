import config from "./config";

function RefreshToken():boolean{
    let oldToken = localStorage.getItem('token') || '{}'
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ token:oldToken })
    }
    let isRefreshed = false;
    fetch(config.apiUrl+'refresh-token', requestOptions)
        .then(response => response.json())
        .then(data => {
            console.log(data)
            // isRefreshed = data.toString().includes("ok");
            // if(isRefreshed){
            //     const {token}:any = data
            //     localStorage.setItem('token', token);
            // }
        });
    return isRefreshed
}

export default RefreshToken