(function() {
  let injectTapEventPlugin = require('react-tap-event-plugin')

  let Header = require('./header')
  let Footer = require('./footer')

  let App = React.createClass({
    render: function(){
      return (
        <Header/>
        <Footer/>
      )
    }
  })
  
  console.log("sa   sas ")
  React.render(<App />, document.getElementById('app'))
})()
