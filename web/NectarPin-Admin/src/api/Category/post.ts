import Service from "/src/plugin/axios/axios"

//添加分类[CreateCategory]
export async function CreateCategory(data:any) {
    try {
        return await Service({
            url: '/Category',
            method: 'post',
            data
        });
    } catch (error) {
        console.error("Error in CreateCategory:", error);
        throw error;
    }
}