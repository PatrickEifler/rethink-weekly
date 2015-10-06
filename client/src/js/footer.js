import React  from 'react'

export default React.createClass({
  render: function(){
    return (
      <footer>
        <p>
          By <a href="http://noty.im">kureikain</a>.
          Support me by buying my book <a href="http://leanpub.com/simplyrethinkdb">Simply RethinkDB</a>
          <nav class="nav">
            <ul>
              <li class="menu-item"><a href="#">Home</a></li>
              <li class="menu-item"><a href="#">Issues</a></li>
              <li class="menu-item"><a href="#">About</a></li>
            </ul>
          </nav>
        </p>
      </footer>
    )
  }
})
