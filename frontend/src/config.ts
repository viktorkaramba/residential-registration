const config = {
    apiUrl: 'http://20.52.189.179:80/',
    localApiUrl: 'http://localhost:8080/',
    headers:{ 'Content-Type': 'application/json',
        'Authorization': 'Bearer '.concat(localStorage.getItem('token') || '{}') },
};

export default config;