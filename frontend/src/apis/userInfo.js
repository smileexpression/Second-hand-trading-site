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
export const changePasswordAPI = ({ oldpassword, newpassword }) => {
    return request({
        url: 'member/changepassword',
        method: 'POST',
        data: {
            oldpassword,
            newpassword
        }
    })
}

//更改个人信息
export const changeInfoAPI = ({ name, gender }) => {
    return request({
        url: 'member/changeinfo',
        method: 'POST',
        data: {
            name,
            gender
        }
    })
}

//地址管理