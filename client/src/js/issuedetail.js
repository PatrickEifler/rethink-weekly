import React  from 'react'
import { Router, Route, Link } from 'react-router'

import {
  Row,
  Col
} from 'elemental'

import Storage from './service/storage'

const storage = new Storage()

export default React.createClass({
  getInitialState: function() {
    return {links: []}
  },

  componentWillMount: function() {
    const id = this.props.params.id
    storage.getIssue(id)
      .then((issue) => { this.setState(issue)})
  },

  render: function(){
    return (
      <Row>
        <Link to="/">Home</Link>

        {this.state.links.map(link => (
          <Row key={link.id}>
            <a href={link.uri}>{link.title}</a>
            <p>{link.desc}</p>
          </Row>
        ))}
      </Row>
    )
  }

})
