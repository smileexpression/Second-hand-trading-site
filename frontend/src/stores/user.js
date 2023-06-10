//管理用户数据

import { defineStore } from "pinia";
import { ref } from "vue";
import { loginAPI } from '@/apis/login'
import { registerAPI } from "@/apis/register";
import { useCartStore } from "./cartStore";
import { getUserInfoAPI, updateAvatarAPI,changePasswordAPI, changeInfoAPI } from "@/apis/userInfo"

export const useUserStore = defineStore ( 'user', ()=> {
    const userInfo = ref({})
    const cartStore = useCartStore()
    //获取用户信息
    const getUserInfo = async ({ account, password }) => {
        const res = await loginAPI({ account, password })
        userInfo.value = res.result
        //cartStore.updateCart()
    }

    //注册用户
    const registerUser = async ({ account, password, nickname, gender }) => {
        const res = await registerAPI({ account, password, nickname, gender })
        userInfo.value = res.result
    }

    //退出登录清除用户信息
    const clearUserInfo = () => {
        userInfo.value = {}
        cartStore.clearCart()
    }

    //更改头像
    const updateAvatar = async ({ pictrueID }) => {
        const res = await updateAvatarAPI(pictrueID)
        userInfo.avatar = res.avatarID
    }

    //更改密码
    const changePassword = async ({ oldpassword, newpassword }) => {
        const res = await changePasswordAPI({ oldpassword, newpassword })
    }

    //修改个人信息
    const changeInfo = async ({ name, gender }) => {
        const res = await changeInfoAPI({ name, gender })
        userInfo.nickname = res.result.name
        userInfo,gender = res.result.gender
    }

    return{
        userInfo,
        getUserInfo,
        registerUser,
        clearUserInfo,
        updateAvatar,
        changePassword,
        changeInfo
    }
},{
    persist: true,
}
)