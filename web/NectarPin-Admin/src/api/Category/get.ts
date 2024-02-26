import Service from "/src/plugin/axios/axios"

//获取分类信息 [GetCategory]
export async function GetCategory(id: number) {
    try {
        return await Service({
            url: '/Category/'+id,
            method: 'get'
        });
    } catch (error) {
        console.error("Error in GetCategory:", error);
        throw error;
    }
}

//获取分类列表 [GetCategoryList]
export async function GetCategoryList(pageSize: number,pageNum: number) {
    try {
        return await Service({
            url: '/Category/list/'+pageSize+'/'+pageNum,
            method: 'get'
        });
    } catch (error) {
        console.error("Error in GetCategoryList:", error);
        throw error;
    }
}