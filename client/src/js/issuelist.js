import React  from 'react'

import {
  Row,
  Col
} from 'elemental'

export default React.createClass({
  getInitialState: function() {
    return {issues: [
      {name: "Issue #1", date: "Septembar 27"},
      {name: "Issue #2", date: "Septembar 29"}
    ]}
  },

  render: function(){
    return (
      <div>
      <h3>CHECK OUT WHAT THE REACT DIGEST SENT OUT BEFORE</h3>
      {this.state.issues.map(issue => (
        <Row>
          <Col sm="1/2">{issue.name}</Col>
          <Col sm="1/2">{issue.date}</Col>
        </Row>
      ))}
      </div>
    )
  }

})
