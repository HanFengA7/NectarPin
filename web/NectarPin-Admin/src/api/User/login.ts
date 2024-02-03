import Service from "/src/plugin/axios/axios"

export function Login(data){
    return Service({
        url : '/User/login',
        method: 'post',
        data
    })
}

export function Login_CheckToken(token){
    return Service({
        url : '/User/checkToken',
        method: 'post',
        data:{
            "token": token
        }
    })
}

