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

const PHASE_SUBSCRIBING = "subscribing"
const PHASE_INIT = "init"

export default React.createClass({
  getInitialState: function() {
    return {phase: PHASE_INIT}
  },

  handleSubscribe: function() {
    this.setState({phase: PHASE_SUBSCRIBING})
  },

  render: function(){
    if (this.state.phase == PHASE_SUBSCRIBING) {
      return (
        <Spinner type="warning" message="Loading..." />
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
