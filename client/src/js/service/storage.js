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
    return Post(this.generateUrl("api/subscriptions"), reader)
  }

  // Get a list of issues
  getIssues() {
    return Get(this.generateUrl("api/issues"))
  }

  // Get a particular issue
  getIssue(id) {
    console.log("WIll get issue", id)
    return Get(`api/issues/${id}`)
  }

  // Get our stas
  getStats() {
  }
}
