const Api = class Api {
  private baseUrl = 'http://127.0.0.1:8080'

  get = function (this: Api, endpoint: string) {
    return fetch(`${this.baseUrl}${endpoint}`).then(async function (res) {
      const json = await res.json()
      if (!res.ok) throw new Error(json.error || 'Request failed')
      return json
    })
  }

  post = function (this: Api, endpoint: string, data: unknown) {
    return fetch(`${this.baseUrl}${endpoint}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    }).then(async function (res) {
      const json = await res.json()
      if (!res.ok) throw new Error(json.error || 'Request failed')
      return json
    })
  }

  upload = function (this: Api, endpoint: string, formData: FormData) {
    return fetch(`${this.baseUrl}${endpoint}`, {
      method: 'POST',
      body: formData,
    }).then(async function (res) {
      const json = await res.json()
      if (!res.ok) throw new Error(json.error || 'Upload failed')
      return json
    })
  }

  delete = function (this: Api, endpoint: string) {
    return fetch(`${this.baseUrl}${endpoint}`, {
      method: 'DELETE',
    }).then(async function (res) {
      const json = await res.json()
      if (!res.ok) throw new Error(json.error || 'Request failed')
      return json
    })
  }
}

export const api = new Api()