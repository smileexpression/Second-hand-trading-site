// 封装更改用户信息的接口
import request from '@/utils/http';

//获取用户信息
export const getUserInfoAPI = () => {
    return request({
        url: '/member',
        method: 'GET',
    })
}

//更换头像
export const updateAvatarAPI = (pictrueID) => {
    return request({
        url: '/member/updateavatar',
        method: 'POST',
        params: pictrueID
    })
}

//更换密码


//更改个人信息


//地址管理