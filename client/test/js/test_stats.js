import {
  React,
  sinon,
  assert,
  expect,
  TestUtils
} from '../test_helper';

import Stat from '../../src/js/stats'

describe('Stat component', () => {
  let sandbox, statComponent, spans

  before(() => {
    sandbox = sinon.sandbox.create()

    statComponent = TestUtils.renderIntoDocument(<Stat subscribers="10" issues="10" />)
    spans = TestUtils.scryRenderedDOMComponentsWithTag(statComponent, 'span')
  })

  after(() => {
    sandbox.restore()
  })

  it('should generate stat block', () => {
    expect(spans.length).to.equal(2)
    //expect(menuitems).to.not.equal(null)
    spans.map((item) => {
      expect(item.getDOMNode().innerHTML).to.equal("10")
    })
    expect(statComponent.getDOMNode().innerHTML).to.contain('subscribers')
    expect(statComponent.getDOMNode().innerHTML).to.contain('issues')
  })
})
