import Promise from 'bluebird'
import {Get, Post} from './request'

export default class Storage {
  constructor() {
    const port = window.location.port == "80"? '':`:${window.location.port}`
    this.apiUri = `//${window.location.hostname}${port}`
  }

  generateUrl(path) {
    return `${this.apiUri}/${path}`
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
    return Get(this.generateUrl(`api/issues/${id}`))
  }

  // Get our stas
  getStats() {
    return Get(this.generateUrl(`api/stats`))
  }
}
