//封装购物车接口

import request from '@/utils/http'

//加入购物车
export const insertCartAPI = (id) => {
    return request({
        url: '/member/cart',
        method: 'POST',
        data: {id}
    })
}
//获取最新购物车列表
export const findNewCartLisrAPI = () =>{
    return request({
        url: '/member/cart'
    })
}