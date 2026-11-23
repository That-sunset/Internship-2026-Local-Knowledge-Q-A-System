import { api } from './api'

export interface ChatRecord {
  id: number
  session_id: number
  content: string
  source: 'user' | 'model'
  createtime: string
}

export function Getrecord() {
  return api.get('/api/records').then(function (data) {
    return data
  })
}

export function Postrecord() {
  return api.post('/api/records', {}).then(function (data) {
    return data
  })
}

export function Deletebyid(recordId: string | number) {
  return api.delete(`/api/records/${recordId}`).then(function (data) {
    return data
  })
}

export function Getchat(sessionId: string | number): Promise<ChatRecord[]> {
  return api.get(`/api/chat/${sessionId}`).then(function (data) {
    return data
  })
}

export function Postchat(
  sessionId: string | number,
  payload: { prompt: string }
) {
  return api.post(`/api/chat/${sessionId}`, payload).then(function (data) {
    return data
  })
}