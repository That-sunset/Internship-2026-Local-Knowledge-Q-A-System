import { api } from './api'

export function Getchat(recordId: string | number) {
  return api.get(`/api/chat/${recordId}`).then(function (data) {
    console.log(data)
    return data
  })
}

export function Getrecord() {
  return api.get('/api/record').then(function (data) {
    console.log(data)
    return data
  })
}

export function Post(payload: { text: string }) {
  return api.post('/api/post', payload).then(function (data) {
    console.log(data)
    return data
  })
}

export function Deletebyid(recordId: string | number) {
  return api.delete(`/api/record/${recordId}`).then(function (data) {
    console.log(data)
    return data
  })
}