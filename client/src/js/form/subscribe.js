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
  getInitialState: function() {
    return {phase: PHASE_INIT}
  },

  handleSubscribe: function() {
    this.setState({phase: PHASE_SUBSCRIBING})
    storage.subscribe({email: "v@v.com"})
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
        <Alert type="success">Cool, check your email to verify the subscription.</Alert>
      )
    }

    if (this.state.phase == PHASE_ERROR) {
      return (
        <Alert type="success">Err, <a href="#">re-try</a> again.</Alert>
      )
    }

    return (
      <div>
      <Form>
        <InputGroup>
          <InputGroup.Section grow>
            <FormInput type="text" placeholder="Email address ..." />
          </InputGroup.Section>
        </InputGroup>
        <InputGroup>
          <InputGroup.Section grow>
            <FormInput type="text" placeholder="First name...(Optional)" />
            <FormInput type="text" placeholder="Last name...(Optinal)" />
          </InputGroup.Section>
        </InputGroup>

        <Button type="default" onClick={this.handleSubscribe}>Subscribe now</Button>
      </Form>
      </div>
    )
  }
})
