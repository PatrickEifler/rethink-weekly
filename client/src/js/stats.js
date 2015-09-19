var Stats = React.createClass({
  render: function(){
    return (
      <div>
      {this.props.subscribes} subscribers . {this.props.issues} issues.
      </div>
    )
  }
});
