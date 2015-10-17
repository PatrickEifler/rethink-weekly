import {
  React,
  sinon,
  assert,
  expect,
  TestUtils
} from '../../test_helper'

import Promise from 'bluebird'

import Spinner from '../../../src/js/spinner'
import FormSubscribe from '../../../src/js/form/subscribe'

describe('Subscriber form component', () => {
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

  before(() => {
    console.log("before")
    // sinon.stub(storageStub.prototype)
    let storageStub = class {
      subscribe(obj) {
        return new Promise((resolve, reject) => {
          resolve(obj)
        })
      }
    }
    FormSubscribe.__Rewire__('storage', new storageStub())
  })

  after(() => {
    FormSubscribe.__ResetDependency__('storage')
  })

  let sandbox, subscribeForm, inputs, button

  beforeEach(() => {
    sandbox = sinon.sandbox.create()

    subscribeForm = TestUtils.renderIntoDocument(<FormSubscribe />)
    inputs = TestUtils.scryRenderedDOMComponentsWithTag(subscribeForm, 'input');
    button = TestUtils.findRenderedDOMComponentWithTag(subscribeForm, 'button')
  })

  afterEach(() => {
    sandbox.restore()
  })

  it('should generate form', () => {
    const p = TestUtils.findRenderedDOMComponentWithTag(subscribeForm, 'p')
    // Two ways, they are same. Either way works
    expect(p.props.children).to.equal("We hate spam, just like you")
    expect(p.getDOMNode().innerHTML).to.equal("We hate spam, just like you")
    expect(inputs.length).to.equal(3)

    expect(button.props.className).to.equal("Button Button--primary")
    expect(button.props.children).to.equal("Subscribe now")
  })

  it('submit form when clicking Subscribe', () => {
    let emailInput = inputs.find((el) => { return el.props.name == 'email' }),
    firstnameInput = inputs.find((el) => { return el.props.name == 'firstname' }),
    lastnameInput = inputs.find((el) => { return el.props.name == 'lastname' })

    TestUtils.Simulate.change(emailInput, { target: { value: 'khicon@axcoto.com' } })
    TestUtils.Simulate.change(firstnameInput, { target: { value: 'Khi' } })
    TestUtils.Simulate.change(lastnameInput, { target: { value: 'Con' } })

    TestUtils.Simulate.click(button)
    //storageStub.subscribe.called.should.be.equal(true)
    //storageStub.subscribe.calledWith({f: 1}).to.be(true)
    expect(subscribeForm).to.equal(<Spinner type="warning" message="Loading..." />)
  })
})

