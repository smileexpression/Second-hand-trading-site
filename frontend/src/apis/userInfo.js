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
export const updateAvatarAPI = (pictureID) => {
    return request({
        url: '/member/updateavatar',
        method: 'POST',
        params: {pictureID}
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
export const changeInfoAPI = ({ nickname, gender }) => {
    return request({
        url: 'member/changeinfo',
        method: 'POST',
        data: {
            nickname,
            gender
        }
    })
}

//新增地址
export const addAddressAPI = ({ receiver, contact, address }) => {
    return request({
        url: 'member/addaddress',
        method: 'POST',
        data:{
            receiver,
            contact,
            address
        }
    })
}

//删除地址
export const delAddressAPI = (id) => {
    return request({
        url: 'member/deladdress',
        method: 'POST',
        params: {id}
    })
}