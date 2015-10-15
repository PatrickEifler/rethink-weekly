import Promise from 'bluebird'

function serialize(obj) {
  var a = []
  for (var p in obj) {
    a.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));
  }
  return a.join("&")
}

export function Post(url, data) {
  const request = new XMLHttpRequest()
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
    request.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    request.send(serialize(data))
  })
}

export function Get(url, query) {
  let request = new XMLHttpRequest()

  console.log("=", url)
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
