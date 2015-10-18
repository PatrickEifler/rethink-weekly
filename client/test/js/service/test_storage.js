import {
  React,
  sinon,
  assert,
  expect,
  TestUtils
} from '../../test_helper';

import Storage from '../../../src/js/service/storage'

describe('Storage', () => {
  let storage, sandbox
      , post, get

  before(()=> {
    GLOBAL.window = GLOBAL
    window.location = {port:80, hostname: 'foo'}

    post = sinon.stub()
    get = sinon.stub()

    // Stub out those dependencies on Storage
    Storage.__Rewire__('Post', post)
    Storage.__Rewire__('Get', get)
    storage = new Storage()

  })

  after(() => {
  })

  beforeEach(() => {
    sandbox = sinon.sandbox.create()
  })

  afterEach(() => {
    sandbox.restore()
  })

  describe('generateUrl', () => {
    it('returns url with host', () => {
      const url = storage.generateUrl("bar")
      expect(url).to.be.equal("//foo/bar")
    })
  })

  describe('subscribe', () => {
    it('send subscriber info to api endpoint', () => {
      let user = {email: 'khicon@axcoto.com'}
      storage.subscribe(user)
      sinon.assert.calledWith(post, storage.generateUrl("api/subscriptions"), user)
    })
  })

  describe('getIssues', () => {
    it('send correct api request', () => {
      let user = {email: 'khicon@axcoto.com'}
      storage.getIssues()
      sinon.assert.calledWith(get, "//foo/api/issues")
    })
  })

  describe('getIssue', () => {
    it('send correct api request', () => {
      storage.getIssue(111)
      sinon.assert.calledWith(get, "//foo/api/issues/111")
    })
  })


})
