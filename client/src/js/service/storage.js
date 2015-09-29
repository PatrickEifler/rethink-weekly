import Promise from 'bluebird'
import {Get, Post} from './request'

export default class Storage {
  constructor() {
    this.apiUri = "/"
  }

  generateUrl(path) {
    return this.apiUri + path
  }

  // Subscribe into our mailing list
  subscribe(reader) {
    // @TODO validation?
    return new Promise((resolve, reject) => {
      resolve({id: "subscribe-id", message: "Check your email to confirm "})
    })
  }

  // Get a list of issues
  getIssues() {
    return Get("api/issues")

    return new Promise((resolve, reject) => {
      resolve([
        {id: 1, name: "Issue #1", date: "Septembar 27"},
        {id: 2, name: "Issue #2", date: "Septembar 29"},
        {id: 3, name: "Issue #3", date: "Septembar 29"}
      ])
    })
  }

  // Get a particular issue
  getIssue(id) {
    console.log("WIll get issue", id)
    return new Promise((resolve, reject) => {
      resolve({
        id: 1,
        name: "Issues #2",
        links: [
          {id: 111, title: 'Simply RethinkDB', uri: 'http://leanpub.com/simplyrethinkdb', desc: "A simply rethinkdb book"},
          {id: 112, title: 'Server Monitoring', uri: 'https://github.com/kureikain/simplyrethink/tree/master/example/server_monitoring', desc: "A simple server monitoring solution with RethinkDB"},
        ]
      })
    })
  }

  // Get our stas
  getStats() {
  }
}
