import Service from "/src/plugin/axios/axios"

//创建文章 [CreateArticle]
export async function CreateArticle(data: any) {
    try {
        return await Service({
            url: '/Article',
            method: 'post',
            data
        });
    } catch (error) {
        console.error("Error in TokenGetUserInfo:", error);
        throw error;
    }
}