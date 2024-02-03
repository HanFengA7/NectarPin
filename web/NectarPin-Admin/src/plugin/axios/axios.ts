import axios from "axios";
import { Message, Notification } from "@arco-design/web-vue";

const Service = axios.create({
    //baseURL: "/",
    baseURL: "http://localhost:3001/api",
    timeout: 300000
})

// 请求拦截器
Service.interceptors.request.use((config: any) => {
    // 在发送请求之前做些什么
    config.headers.Authorization = `Bearer ${window.localStorage.getItem(
        "token"
    )}`;
    return config;
});

//响应拦截器
Service.interceptors.response.use((response) => {
    //获取接口返回结果
    const res = response.data
    if (response.status === 200) {
        return response
    } else {
        let ResMessage: string
        if (res.message != null) {
            ResMessage = res.message
        } else {
            ResMessage = '网络异常'
        }
        Message.loading('数据处理中...')
        Notification.error(ResMessage)
        return response
    }
})

export default Service