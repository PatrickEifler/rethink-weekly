import Footer from './footer'
import Header from './header'
import SubscribeForm from './form/subscribe'
import IssueList from './issuelist'
import IssueDetail from './issuedetail'
import About from './about'
import NavMenu from './nav'

import React  from 'react'
import createBrowserHistory from 'history/lib/createBrowserHistory'
import { Router, Route, Link, IndexRoute } from 'react-router'

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
  Table,
  Row
} from 'elemental'

(function() {

  const App = React.createClass({
    render: function() {
      return (
        <div className="page-wrapper">
          <NavMenu />
          <div className="page-body">
            <Header />
            <SubscribeForm />
            <Container maxWidth={768} className="demo-container">
              {this.props.children}
            </Container>
          </div>
          <Footer />
        </div>
      )
    }
  })

  console.log("Thanks for interesting. Help me by buying my book Simply RethinkDB")

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

// Using https://github.com/rackt/react-router/blob/master/docs/guides/basics/Histories.md#createbrowserhistory
// We want browser history for better looking URL instead of hash URL
// This is used with the package https://github.com/rackt/history
// We are writing our own Go lang server for serving static file.
// The server has to be able to handle those link, basiclly whatever URL, we always show index.html file,
// unless the /api route or some special route like /stats
  React.render((
    <Router history={createBrowserHistory()}>
      <Route name="home" path="/" component={App}>
        <IndexRoute component={IssueList}/>
        <Route name="issues" path="issues" component={IssueList} />
        <Route path="issues/:id" component={IssueDetail} />
        <Route name="about" path="about" component={About} />
      </Route>
    </Router>
  ), document.getElementById('app'))
})()
