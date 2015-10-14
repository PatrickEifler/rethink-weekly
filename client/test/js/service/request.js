import {
  React,
  sinon,
  assert,
  expect,
  TestUtils
} from '../test_helper';

import { Post, Get } from '../../src/js/service/request'

describe('Request', () => {
  setUp: function () {
    this.server = sinon.fakeServer.create()
  },

  tearDown: function () {
    this.server.restore()
  },


  beforeEach(() => {
    sandbox = sinon.sandbox.create()

    footerComponent = TestUtils.renderIntoDocument(<Issuelist />)
    anchors = TestUtils.scryRenderedDOMComponentsWithTag(footerComponent, 'a');
    menuitems = TestUtils.findRenderedDOMComponentWithTag(footerComponent, 'ul');
  })

  afterEach(() => {
    sandbox.restore()
  })

  describe('Get', () => {
    this.server.respondWith("GET", "/foo",
            [200, { "Content-Type": "application/json" },
             '[{ "id": 12, "comment": "Hey there" }]']);

    it('returns response', (done) => {
      Get('foo', 'foo=bar').then((result) => {
        expect(result).to.equal({id: 12, comment: "Hey there"})
        done()
      })
    })
  })
})
