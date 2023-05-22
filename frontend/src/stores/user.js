//管理用户数据

import { defineStore } from "pinia";
import { ref } from "vue";
import { loginAPI } from '@/apis/login'

export const useUserStore = defineStore ( 'user', ()=> {
    const userInfo = ref({})
    //获取用户信息
    const getUserInfo = async ({ account, password }) => {
        const res = await loginAPI({ account, password })
        userInfo.value = res.result
    }

    //退出登录清除用户信息
    const clearUserInfo = () => {
        userInfo.value = {}
    }

    return{
        userInfo,
        getUserInfo,
        clearUserInfo
    }
},{
    persist: true,
}
)