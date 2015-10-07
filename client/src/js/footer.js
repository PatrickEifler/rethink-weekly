import React  from 'react'

export default React.createClass({
  render: function(){
    return (
      <footer className="page-footer">
        <p>
          By <a href="https://axcoto.com">kureikain</a>.
          Support me by buying my book <a href="http://leanpub.com/simplyrethinkdb">Simply RethinkDB</a>.
        </p>
        <p>
          <a href="github.com/axcoto/rethink-weekly">Build</a> with Golang, React and <a href="elemental-ui.com">ElementalUI</a>.
        </p>
      </footer>
    )
  }
})
