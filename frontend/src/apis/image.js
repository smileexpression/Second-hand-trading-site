import request from '@/utils/http'

export const uploadImageAPI = (formData) => {
  return request({
    url: '/image/upload',
    method: 'POST',
    body: formData,
  })
}

export const getImageUrl = (id) => {
  return `http://localhost:8080/image/get?id=${id}`
}

export const deleteImageAPI = (id) => {
  return request({
    url: '/image/delete',
    method: 'POST',
    params: { id },
  })
}

// export const url = request.defaults.upload + '/image/upload'