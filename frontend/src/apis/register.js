//封装和注册相关的接口
import request from '@/utils/http';

export const registerAPI = ({account, 
                             password,
                             nickname,
                             gender}) => {
    return request({
        url: '/register',
        method: 'POST',
        data:{
            account,
            password,
            nickname,
            gender
        }
    })
}