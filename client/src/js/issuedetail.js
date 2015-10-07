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
      .then((issue) => { this.setState({links: issue})})
  },

  render: function(){
    return (
      <div>
        {this.state.links.map(link => (
          <ul key={link.id}>
            <li>
              <h3><a href={link.uri}>{link.title}</a> by {link.author}</h3>
              <p>{link.desc}</p>
            </li>
          </ul>
        ))}
      </div>
    )
  }

})
