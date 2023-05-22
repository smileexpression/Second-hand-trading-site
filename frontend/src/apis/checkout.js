import request from '@/utils/http'
/**
 * 获取结算信息
 */
export const getCheckInfoAPI = () => {
  return request({
    url: '/member/order/pre'
  })
}