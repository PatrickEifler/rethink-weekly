import Footer from './footer'
import Header from './header'

(function() {
  //let injectTapEventPlugin = require('react-tap-event-plugin')

  class App extends React.Component {
    render() {
      return (
        <div>
          <Header />
          <Footer />
        </div>
      )
    }
  }

  console.log("Woohoo...loaded")
  React.render(<App />, document.getElementById('app'))
})()
