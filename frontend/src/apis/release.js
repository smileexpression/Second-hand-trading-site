//封装和发布闲置相关的接口
import request from '@/utils/http';

export const releaseAPI = ({ name,
  cate_id,
  description,
  picture,
  price, }) => {
  return request({
    url: '/member/release',
    method: 'POST',
    data: {
      name,
      cate_id,
      description,
      picture,
      price,
    }
  })
}