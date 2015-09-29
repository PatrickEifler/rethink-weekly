import Promise from 'bluebird'

export function Post(url, data) {
  let request = new XMLHttpRequest()
  return new Promise(function(resolve, reject) {
    request.onload = function () {
      if (request.status >= 200 && request.status <400) {
        resolve(JSON.parse(request.responseText))
      } else {
        reject({error: "Server return error", response: {body: request.responseText, status: request.status}})
      }
    }

    request.onerror = function() {
      reject({error: "Connection error"})
    }

    request.open('POST', url, true)
    request.send()
  })
}

export function Get(url, query) {
  let request = new XMLHttpRequest()
  return new Promise(function(resolve, reject) {
    request.onload = function () {
      if (request.status >= 200 && request.status <400) {
        resolve(JSON.parse(request.responseText))
      } else {
        reject({error: "Server return error", response: {body: request.responseText, status: request.status}})
      }
    }

    request.onerror = function() {
      reject({error: "Connection error"})
    }

    request.open('GET', url, true)
    request.send()
  })
}
