// axios基础的封装
import axios from 'axios'
import { ElMessage } from 'element-plus'
import 'element-plus/es/components/message/style/css'

const httpInstance = axios.create({
  baseURL: 'https://mock.apifox.cn/m1/2726765-0-default',
  timeout: 5000
})

// axios请求拦截器
httpInstance.interceptors.request.use(config => {
  return config
}, e => Promise.reject(e))

// axios响应式拦截器
httpInstance.interceptors.response.use(res => res.data, e => {
  //统一提示错误
  ElMessage({
    type: 'warning',
    message: e.response.data.msg
  })
  return Promise.reject(e)
})

export default httpInstance