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
  Table
} from 'elemental'

export default React.createClass({
  render: function(){
    return (
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

        <Button type="default">Subscribe now</Button>
      </Form>
    )
  }
})
