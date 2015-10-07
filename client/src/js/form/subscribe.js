import React  from 'react'

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
  Alert
} from 'elemental'

import Spinner from '../spinner'
import Storage from '../service/storage'

const PHASE_SUBSCRIBING = "subscribing"
const PHASE_INIT = "init"
const PHASE_DONE = "done"
const PHASE_ERROR = "error"

const storage = new Storage()

export default React.createClass({
  mixins: [React.addons.LinkedStateMixin],

  getInitialState: function() {
    return {
      phase: PHASE_INIT,
      email: "",
      firstname: "",
      lastname: ""
    }
  },

  handleSubscribe: function() {
    this.setState({phase: PHASE_SUBSCRIBING})
    console.log(this.state)

    storage.subscribe({
      email: this.state.email,
      firstname: this.state.firstname,
      lastname: this.state.lastname,
    })
      .then(result => {
        console.log(result)
        this.setState({phase: PHASE_DONE})
      })
      .error(err => {
        this.setState({phase: PHASE_ERROR})
      })
  },

  render: function(){
    if (this.state.phase == PHASE_SUBSCRIBING) {
      return (
        <Spinner type="warning" message="Loading..." />
      )
    }

    if (this.state.phase == PHASE_DONE) {
      return (
        <Alert type="success">Thank you for subscribing. Just one more step, check your email and confirm it. We are doing this to prevent other people subscribe on your behalf :)</Alert>
      )
    }

    let alertBox
    if (this.state.phase == PHASE_ERROR) {
      alertBox = <Alert type="success">Err, <a href="#" onClick={this.handleSubscribe}>re-try</a> again.</Alert>
    } else {
    }

    return (
      <div className="demo-banner demo-banner--secondary">
        <Container maxWidth={768} className="demo-container">
          {alertBox}
          <Form>
            <InputGroup>
              <InputGroup.Section grow>
                <FormInput type="text" placeholder="Email address ..." valueLink={this.linkState('email')} />
              </InputGroup.Section>
            </InputGroup>
            <InputGroup>
              <InputGroup.Section grow>
                <FormInput type="text" placeholder="First name...(Optional)" valueLink={this.linkState('firstname')} />
              </InputGroup.Section>
              <InputGroup.Section grow>
                <FormInput type="text" placeholder="Last name...(Optinal)" valueLink={this.linkState('lastname')} />
              </InputGroup.Section>
            </InputGroup>

            <Button type="primary" onClick={this.handleSubscribe}>Subscribe now</Button>
            <p>We hate spam, just like you</p>
          </Form>
        </Container>
      </div>
    )
  }
})
