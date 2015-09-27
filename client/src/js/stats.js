import React  from 'react'

var Stats = React.createClass({
  render: function(){
    return (
      <div>
      {this.props.subscribers} subscribers . {this.props.issues} issues.
      </div>
    )
  }
})

module.exports = Stats
