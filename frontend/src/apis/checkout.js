import request from '@/utils/http'
/**
 * 获取结算信息
 */
export const getCheckInfoAPI = (id) => {
  return request({
    url: '/member/order/pre',
    id
  })
}

// 创建订单
export const createOrderAPI = (data) => {
  return request({
    url: '/member/order',
    method: 'POST',
    data
  })
}