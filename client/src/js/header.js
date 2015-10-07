import React  from 'react'
import Stats from './stats'
import { Container } from 'elemental'

var Header = React.createClass({
  render: function(){
    return (
      <header className="demo-banner demo-banner--primary">
        <Container maxWidth={768} className="demo-container">
         <h1>RethinkDB Weekly Stuff</h1>
         <h2>A hand-picked weekly selection of the best RethinkDB resources</h2>
         <Stats issues="10" subscribers="10" />
         <p>Ok, the stat is faked. We haven't done yet</p>
       </Container>
      </header>
    )
  }
})

module.exports = Header
