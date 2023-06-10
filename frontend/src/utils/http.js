// axios基础的封装
import axios from 'axios'
import { ElMessage } from 'element-plus'
import 'element-plus/es/components/message/style/css'
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'

const httpInstance = axios.create({
//   baseURL: 'https://mock.apifox.cn/m1/2726765-0-default',
  baseURL: 'http://localhost:8080',
  timeout: 5000
})

// axios请求拦截器
httpInstance.interceptors.request.use(config => {
  //获取token数据
  const userStore = useUserStore()
  const token = userStore.userInfo.token
  //将token拼接到数据报头部
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
}, e => Promise.reject(e))

// axios响应式拦截器
httpInstance.interceptors.response.use(res => res?.data, e => {
  const userStore = useUserStore()
  const router = useRouter()

  // console.log(e.response)
  //统一提示错误
  ElMessage({
    type: 'warning',
    message: e.response.data.msg
  })

  //token失效处理
  if (e.response.status === 401) {
    userStore.clearUserInfo()
    router.push('/login')
  }

  return Promise.reject(e)
})

export default httpInstance