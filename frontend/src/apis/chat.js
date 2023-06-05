import httpInstance from '@/utils/http'

export const getChatListAPI = () => {
  return httpInstance({
    url: '/chat/getmsg',
  })
}

export const addChatListAPI = (data) => {
  return httpInstance({
    url: '/chat/sendmsg',
    method: 'POST',
    data
  })
}