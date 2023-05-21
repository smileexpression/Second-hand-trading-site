import httpInstance from '@/utils/http';

// 获取banner

export function getBannerAPI() {
  return httpInstance({
    url: '/home/banner',
  })
}

/**
 * @description: 获取最近几条发布闲置
 * @param {*}
 * @return {*}
 */
export const findNewAPI = () => {
  return httpInstance({
    url: '/home/new'
  })
}

/**
 * @description: 获取所有商品模块
 * @param {*}
 * @return {*}
 */
export const getGoodsAPI = () => {
  return httpInstance({
    url: '/home/goods'
  })
}