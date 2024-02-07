import Service from "/src/plugin/axios/axios"

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

export async function EditUserInfo(id: int,data: any) {
    try {
        return await Service({
            url: '/User/editInfo/'+ id,
            method: 'put',
            data
        })
    } catch (error) {
        console.error("Error in EditUserInfo:", error);
        throw error;
    }
}