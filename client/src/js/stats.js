import React  from 'react'

var Stats = React.createClass({
  render: function(){
    return (
      <div>
        <span>{this.props.subscribers}</span> subscribers. <span>{this.props.issues}</span> issues.
      </div>
    )
  }
})

module.exports = Stats
