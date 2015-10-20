import {
  React,
  sinon,
  assert,
  expect,
  TestUtils
} from '../test_helper';

import Issuelist from '../../src/js/issuelist'

describe('Issuelist component', () => {
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

  let sandbox, issueListComponent, headline, items
  let storageStub = class {
    getIssues() {
      return new Promise((resolve, reject) => {
        resolve([
          {id: 1, links: "link1", name: "foo1", date: "d1"},
          {id: 2, links: "link2", name: "foo2", data: "d2"},
        ])
      })
    }
  }
  Issuelist.__Rewire__('storage', new storageStub())
  Issuelist.__Rewire__('Link', React.createClass({
    render: () => { <a {...this.props}> </a> }
  }))


  beforeEach(() => {
    sandbox = sinon.sandbox.create()

    issueListComponent = TestUtils.renderIntoDocument(<Issuelist />)
    console.log(issueListComponent)
    headline = TestUtils.scryRenderedDOMComponentsWithTag(issueListComponent, 'h3')
    items = TestUtils.findRenderedDOMComponentWithTag(issueListComponent, 'a')
  })

  afterEach(() => {
    sandbox.restore()
    Issuelist.__ResetDependency__('storage')
  })

  it('should generate headline', () => {
    expect(headline.childrens).to.equal("CHECK OUT WHAT WE SEND BEFORE")
  })

  it('should generate list of lins', () => {
    expect(items.length).to.equal(2)
  })

})
