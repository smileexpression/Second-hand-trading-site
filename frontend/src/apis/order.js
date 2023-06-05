//订单模块接口

import request from '@/utils/http'

export const getUserOrder = (params) => {
    return request({
        url: '/member/order',
        method: 'GET',
        params
    })
}