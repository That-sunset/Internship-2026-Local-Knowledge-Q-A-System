class Api {
  private baseUrl = 'http://127.0.0.1:8080'

  get = function (this: Api, endpoint: string) {
    return fetch(`${this.baseUrl}${endpoint}`).then(function (res) {
      return res.json()
    })
  }

  post = function (this: Api, endpoint: string, data: unknown) {
    return fetch(`${this.baseUrl}${endpoint}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    }).then(function (res) {
      return res.json()
    })
  }

  delete = function (this: Api, endpoint: string) {
    return fetch(`${this.baseUrl}${endpoint}`, {
      method: 'DELETE',
    }).then(function (res) {
      return res.json()
    })
  }
}

export const api = new Api()