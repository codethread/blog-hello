fetch('/api/__healthcheck')
    .then(res => res.ok && res.json().then(handleResponse))
    .catch(console.error)


function handleResponse(res = {}) {
    console.log(res);
}
