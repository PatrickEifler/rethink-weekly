import {
  React,
  sinon,
  assert,
  expect,
  TestUtils
} from '../test_helper';

import Footer from '../../src/js/footer'

describe('Footer component', () => {
 let data = {
    form: {
      formAttrs: {
        email: {
          label: "Email Address",
          type: "email",
          value: "foo@bar.com",
          validation: "value.match(/^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9]+$/i)",
          errorMessage: "A valid email address is required"
        },
        password: {
          label: "Password",
          type: "password",
          value: "foobar",
          validation: "value.length > 0 && value.length < 73",
          errorMessage: "Password must be between 1 and 72 characters long"
        }
      }
    }
  }

  let sandbox, footerComponent, anchors, menuitems

  beforeEach(() => {
    sandbox = sinon.sandbox.create()

    footerComponent = TestUtils.renderIntoDocument(<Footer />)
    anchors = TestUtils.scryRenderedDOMComponentsWithTag(footerComponent, 'a');
    menuitems = TestUtils.findRenderedDOMComponentWithTag(footerComponent, 'footer');
  })

  afterEach(() => {
    sandbox.restore()
  })

  it('should generate footer block', () => {
    expect(anchors.length).to.equal(4)
    expect(menuitems).to.not.equal(null)
  })
})
