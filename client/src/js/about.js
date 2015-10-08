import React  from 'react'
import { Router, Route, Link } from 'react-router'

import {
  Row,
  Col,
  Container
} from 'elemental'

export default React.createClass({
  render: function(){
    return (
      <div className="demo-banner demo-banner--tertiary">
        <Container maxWidth={768} className="demo-container">
          <h2 className="demo-banner__heading demo-banner__heading-2">What is RethinkDB Weekly</h2>
          <p>I started to collect useful news, article, video about RethinkDB and I think I will
          create a news letter to share about it.</p>

          <p>Once you subscribe to it, I will send out email on Tuesday about RethinkDB stuff.
          <strong>Why tuesday?</strong>, you asked. On Monday I'm busy to clean up thing and finalize stuff.
          On Tuesday, it's out to warm up our week. Then wenesday cames, then thursday and weekend.
          </p>
 
          <p>Also, I want to learn more about React and Go lang. I write front-end of this site in React.
          Next, I can implement a back-end in any language to learn about it</p>

        </Container>
      </div>
    )
  }

})
