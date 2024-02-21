import Service from "/src/plugin/axios/axios"

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