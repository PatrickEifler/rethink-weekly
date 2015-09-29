import React  from 'react'

import {
  Alert,
  Spinner
} from 'elemental'

export default React.createClass({
  getDefaultProps: function() {
    return {
      message: "Loading...",
      type: "info",
      size: "lg"
    }
  },

  render: function(){
    return (
      <div>
        <Alert type="{this.props.type}"><strong>Warning:</strong> {this.props.message}</Alert>
        <Spinner size="{this.props.size}" />
      </div>
    )
  }
})
