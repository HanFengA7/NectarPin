import Service from "/src/plugin/axios/axios"

export function Login(data){
    return Service({
        url : '/User/login',
        method: 'post',
        data
    })
}

export function Login_CheckToken(data){
    return Service({
        url : '/User/login',
        method: 'post',
        data
    })
}