import React  from 'react'
import { Router, Route, Link } from 'react-router'

import {
  Row,
  Col
} from 'elemental'

import Storage from './service/storage'

let storage = new Storage()

export default class IssuelistCompoentn extends React.Component{
  getInitialState () {
    return {issues: []}
  }

  componentWillMount () {
    storage.getIssues()
      .then((issues) => { this.setState({issues: issues})})
  }

  render (){
    return (
      <div>
      lol
      <h3>CHECK OUT WHAT WE SENT OUT BEFORE</h3>
      {this.state.issues.map(issue => (
        // We need the key attr for react https://fb.me/react-warning-keys
        <Row key={issue.id}>
          <Col>
            <Link to={"/issues/" + issue.id}>{issue.name}</Link>
            &nbsp;-&nbsp;{new Date(issue.date).toDateString()}
          </Col>
        </Row>
      ))}
      </div>
    )
  }
}
