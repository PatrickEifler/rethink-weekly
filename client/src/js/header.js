var Stats = require("./stats")

var Header = React.createClass({
  render: function(){
    return (
      <header>
       <h2>RethinkDB Weekly Goodies</h2>
       <span>A hand-picked weekly selection of the best RethinkDB resources</span>
       <Stats issues="10" subscribers="10" />
      </header>
    )
  }
})

module.exports = Header
