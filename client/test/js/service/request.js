import {
  React,
  sinon,
  assert,
  expect,
  TestUtils
} from '../../test_helper';

import { Post, Get } from '../../../src/js/service/request'

describe('Request', () => {
  let server = sinon.fakeServer.create()
  let sandbox

  before(()=> {
    console.log("Invoke")
    server = sinon.fakeServer.create()
    console.log(server)
    var xhr = sinon.useFakeXMLHttpRequest();
  })

  after(() => {
    server.restore()
  })

  var xhr
  beforeEach(() => {
    sandbox = sinon.sandbox.create()
    server = sinon.fakeServer.create()

    // Overwrite the global XMLHttpRequest with Sinon fakes
    xhr = sinon.useFakeXMLHttpRequest();
    // Create an array to store requests
    var requests = [];
    // Keep references to created requests
    xhr.onCreate = function (xhr) {
      requests.push(xhr);
    };

  })

  afterEach(() => {
    sandbox.restore()
    xhr.restore()
  })

  describe('Get', () => {

    console.log(server)

    it('returns response', (done) => {
      Get('/foo', 'foo=bar').then((result) => {
        expect(result).to.equal({id: 12, comment: "Hey there"})
        done()
      })
      .error((err) => {
        console.log(err)
      })

       server.respondWith("GET", "/foo",
            [200, { "Content-Type": "application/json" },
             '[{ "id": 12, "comment": "Hey there" }]']);


    })

    it('returns error in invalid request', (done) => {
      Get('/foo', '').then((result) => {
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
