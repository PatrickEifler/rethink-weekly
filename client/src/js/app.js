import Footer from './footer'
import Header from './header'

(function() {
  //let injectTapEventPlugin = require('react-tap-event-plugin')

  let App = React.createClass({
    render: function(){
      return (
        <div>
          <Header />
          <Footer />
        </div>
      )
    }
  })

  console.log("Woohoo...loaded")
  React.render(<App />, document.getElementById('app'))
})()
