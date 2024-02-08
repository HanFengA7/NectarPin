import Service from "/src/plugin/axios/axios"

//Token信息解密 [TokenGetUserInfo]
export async function TokenGetUserInfo(data: any) {
    try {
        return await Service({
            url: '/User/tokenInfo',
            method: 'post',
            data: {
                "token": data,
            }
        });
    } catch (error) {
        console.error("Error in TokenGetUserInfo:", error);
        throw error;
    }
}

//获取用户信息 [GetUserInfo]
export async function GetUserInfo(id: any) {
    try {
        return await Service({
            url: '/User/getInfo/1/' + id,
            method: 'get'
        });
    } catch (error) {
        console.error("Error in GetUserInfo:", error);
        throw error;
    }
}

//编辑用户信息 [EditUserInfo]
export async function EditUserInfo(id: int, data: any) {
    try {
        return await Service({
            url: '/User/editInfo/' + id,
            method: 'put',
            data
        })
    } catch (error) {
        console.error("Error in EditUserInfo:", error);
        throw error;
    }
}

//修改用户密码 [EditUserPwd]
export async function EditUserPwd(username: string, password: string) {
    try {
        return await Service({
            url: '/User/editPwd',
            method: 'put',
            data:{
                "username":username,
                "password":password
            }
        })
    } catch (error){
        console.error("Error in EditUserPwd:", error);
        throw error;
    }
}