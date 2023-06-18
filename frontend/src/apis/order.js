//订单模块接口

import request from '@/utils/http'

export const getUserOrder = (params) => {
    return request({
        url: '/member/getorder',
        method: 'GET',
        params
    })
}

export const getSoldOrder = (params) => {
    return request({
        url: '/member/soldorder',
        method: 'GET',
        params
    })
}

export const getRemain = (params) => {
    return request({
        url: '/member/remain',
        method: 'GET',
        params
    })
}