import {
  React,
  sinon,
  assert,
  expect,
  TestUtils
} from '../../test_helper';

import { Post, Get } from '../../../src/js/service/request'

describe('Request', () => {
  var FakeXMLHttpRequests = require('fakexmlhttprequest')

  var requests   = []
  XMLHttpRequest = function() {
    console.log("here call me")
    var r =  new FakeXMLHttpRequests()
    requests.push(r)
    return r
  }


  let server = sinon.fakeServer.create()
  let sandbox

  before(()=> {
    console.log("Invoke")
    server = sinon.fakeServer.create()
    console.log(server)
  })

  after(() => {
    server.restore()
  })

  beforeEach(() => {
    sandbox = sinon.sandbox.create()
    server = sinon.fakeServer.create()
  })

  afterEach(() => {
    sandbox.restore()
  })

  describe('Get', () => {

    console.log(server)
    server.respondWith("GET", "/foo",
            [200, { "Content-Type": "application/json" },
             '[{ "id": 12, "comment": "Hey there" }]']);

    it('returns response', (done) => {
      Get('http://foo', 'foo=bar').then((result) => {
        expect(result).to.equal({id: 12, comment: "Hey there"})
        done()
      })
      .error((err) => {
        console.log(err)
      })
    })

    it('returns error in invalid request', (done) => {
      Get('http://gmail.com', '').then((result) => {
        expect(result).to.equal({id: 12, comment: "Hey there"})
        done()
      })
      .error((err) => {
        done()
        expect(err.response.status).to.be.equal(503)
      })
    })
  })
})
