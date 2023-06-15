import request from '@/utils/http'
/**
 * 获取结算信息
 */
export const getCheckInfoAPI = (goodID) => {
  return request({
    url: '/member/order/pre',
    params: {
      goodID
    }
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