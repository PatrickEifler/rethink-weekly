(function() {
  let injectTapEventPlugin = require('react-tap-event-plugin')
  let App = React.createClass({
    render: function(){
      return (
        <Header/>
      )
    }
  });
  React.render(<App />, document.getElementById('app'));
})()
