import Service from "/src/plugin/axios/axios"

export function TokenGetUserInfo(data:any){
    return Service({
        url : '/User/tokenInfo',
        method: 'post',
        data: {
            "token": data,
        }
    })
}

export function GetUserInfo(id:any){
    return Service({
        url : '/User/getInfo/1/'+id,
        method: 'get'
    })
}