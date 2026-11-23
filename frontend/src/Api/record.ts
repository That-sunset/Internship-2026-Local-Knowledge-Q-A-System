import { api } from './api'

export function Getrecord() {
  return api.get('/api/records').then(function (data) {
    console.log(data)
    return data
  })
}

export function Post(payload: { text: string }) {
  return api.post('/api/records', { content: payload.text }).then(function (data) {
    console.log(data)
    return data
  })
}

export function Deletebyid(recordId: string | number) {
  return api.delete(`/api/records/${recordId}`).then(function (data) {
    console.log(data)
    return data
  })
}

// Only keep this if you add r.GET("/api/chat/:id", ...) to main.go
export function Getchat(recordId: string | number) {
  return api.get(`/api/chat/${recordId}`).then(function (data) {
    console.log(data)
    return data
  })
}