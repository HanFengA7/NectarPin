import Service from "/src/plugin/axios/axios"

//删除分类 [DeleteCategory]
export async function DeleteCategory(id: number) {
    try {
        return await Service({
            url: '/Category/'+id,
            method: 'delete'
        });
    } catch (error) {
        console.error("Error in DeleteCategory:", error);
        throw error;
    }
}