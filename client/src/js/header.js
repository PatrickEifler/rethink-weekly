import React  from 'react'
import Stats from './stats'
import { Container } from 'elemental'

import Storage from './service/storage'
const storage = new Storage()

var Header = React.createClass({
  getInitialState () {
    return {stat: {}}
  },

  componentWillMount () {
    storage.getStats()
      .then((stat) => { this.setState({stat}) })
  },

  render (){
    return (
      <header className="demo-banner demo-banner--primary">
        <Container maxWidth={768} className="demo-container">
         <h1>RethinkDB Weekly Stuff</h1>
         <h2>A hand-picked weekly selection of the best RethinkDB resources</h2>
         <Stats issues={this.state.stat.issues} subscribers={this.state.stat.subscribers} />
       </Container>
      </header>
    )
  }
})

module.exports = Header
