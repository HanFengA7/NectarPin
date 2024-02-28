import Service from "/src/plugin/axios/axios"

//编辑分类[editCategory]
export async function editCategory(id: number, data: any) {
    try {
        return await Service({
            url: '/Category/'+id,
            method: 'put',
            data
        });
    } catch (error) {
        console.error("Error in editCategory:", error);
        throw error;
    }
}