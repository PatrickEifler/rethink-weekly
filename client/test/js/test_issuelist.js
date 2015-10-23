import {
  React,
  sinon,
  assert,
  expect,
  TestUtils
} from '../test_helper';

import Issuelist from '../../src/js/issuelist'

describe('Issuelist component', () => {
  let sandbox, issueListComponent, headline, items;
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

  before((done) => {
    sandbox = sinon.sandbox.create()

    Issuelist.__Rewire__('storage', new storageStub())
    Issuelist.__Rewire__('Link', React.createClass({
      render: function() { return (
        <a {...this.props}>link name</a>
        )
      }
    }))

    issueListComponent = TestUtils.renderIntoDocument(<Issuelist />)

    React.render(<Issuelist />, React.findDOMNode(issueListComponent).parentNode);
    //expect(view.state.someVal).to.be(someNewValue);
    headline = TestUtils.findRenderedDOMComponentWithTag(issueListComponent, 'h3')
    issueListComponent.componentWillMount()
    setTimeout(() => {
      items = TestUtils.scryRenderedDOMComponentsWithTag(issueListComponent, 'a')
      done()
    }, 100)
  })

  after(() => {
    sandbox.restore()
    Issuelist.__ResetDependency__('storage')
    Issuelist.__ResetDependency__('Link')
  })

  it('should generate headline', () => {
    expect(headline.props.children).to.equal("CHECK OUT WHAT WE SENT OUT BEFORE")
  })

  it('should generate list of links', () => {
    expect(issueListComponent.getDOMNode().innerHTML).to.contains("link name")
    expect(items.length).to.equal(2)
  })

})
