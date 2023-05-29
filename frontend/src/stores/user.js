//管理用户数据

import { defineStore } from "pinia";
import { ref } from "vue";
import { loginAPI } from '@/apis/login'
import { registerAPI } from "@/apis/register";


export const useUserStore = defineStore ( 'user', ()=> {
    const userInfo = ref({})
    //获取用户信息
    const getUserInfo = async ({ account, password }) => {
        const res = await loginAPI({ account, password })
        userInfo.value = res.result
    }

    //注册用户
    const registerUser = async ({ account, password, nickname, gender }) => {
        const res = await registerAPI({ account, password, nickname, gender })
        userInfo.value = res.result
    }

    //退出登录清除用户信息
    const clearUserInfo = () => {
        userInfo.value = {}
    }

    return{
        userInfo,
        getUserInfo,
        registerUser,
        clearUserInfo
    }
},{
    persist: true,
}
)