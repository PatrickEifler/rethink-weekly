import Footer from './footer'
import Header from './header'
import React  from 'react'
import { Router, Route, Link } from 'react-router'

import {
  Button,
  Checkbox,
  Container,
  EmailInputGroup,
  FileDragAndDrop,
  FileUpload,
  Form,
  FormField,
  FormIconField,
  FormInput,
  FormNote,
  FormRow,
  FormSelect,
  InputGroup,
  PasswordInputGroup,
  Radio,
  RadioGroup,
  Table
} from 'elemental'

(function() {

  const App = React.createClass({
    render: function() {
      return (
        <div>
          <Header />
          <Button size="lg">Large Button</Button>

          <Footer />
        </div>
      )
    }
  })

  console.log("Woohoo...loaded")

  // Declarative route configuration (could also load this config lazily
  // instead, all you really need is a single root route, you don't need to
  // colocate the entire config).
//  React.render((
//    <Router>
//      <Route path="/" component={App}>
//        <Route path="about" component={About}/>
//        <Route path="users" component={Users}>
//          <Route path="/user/:userId" component={User}/>
//        </Route>
//        <Route path="*" component={NoMatch}/>
//      </Route>
//    </Router>
//  ), document.getElementById('app'))
  React.render(<App />, document.getElementById('app'))

})()
