import Service from "/src/plugin/axios/axios"

export function Login(data){
    return Service({
        url : '/User/login',
        method: 'post',
        data
    })
}
