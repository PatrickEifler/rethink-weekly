import React  from 'react'

import { Router, Route, Link, IndexRoute } from 'react-router'

const NavItems = [
  { value: '/',         label: 'Home' },
  { value: '/issues',        label: 'Issues' },
  { value: '/about',     label: 'About' },
]

export default React.createClass({
  getInitialState: function (){
    return  {
      mobileMenuIsVisible: false,
      windowHeight: window.innerHeight,
      windowWidth: window.innerWidth
    }
  },

  toggleMenu: function () {
    this.setState({
      mobileMenuIsVisible: !this.state.mobileMenuIsVisible
    })
  },

  render: function (){
    var self = this;
    var height = (this.state.windowWidth < 768) ? this.state.windowHeight : 'auto';
    var menuClass = this.state.mobileMenuIsVisible ? 'primary-nav-menu is-visible' : 'primary-nav-menu is-hidden';
    var menuItems = NavItems.map(function(item) {
      return (
        <Link key={item.value} className="primary-nav__item" onClick={self.toggleMenu} to={item.value}>
          <span className="primary-nav__item-inner">{item.label}</span>
        </Link>
      )
    })

    return (
      <nav className="primary-nav">
        <Link to="home" className="primary-nav__brand special" title="Home">
          <img src="./images/elemental-logo-paths.svg" className="primary-nav__brand-src" />
        </Link>
        <button onClick={this.toggleMenu} className="primary-nav__item primary-nav-menu-trigger">
          <span className="primary-nav-menu-trigger-icon octicon octicon-navicon" />
          <span className="primary-nav-menu-trigger-label">{this.state.mobileMenuIsVisible ? 'Close' : 'Menu'}</span>
        </button>
        <div className={menuClass} style={{ height }}>
          <div className="primary-nav-menu-inner">
            {menuItems}
          </div>
        </div>
        <a href="https://github.com/elementalui/elemental" target="_blank" title="View on GitHub" className="primary-nav__brand right">
          <img src="./images/github-logo.svg" className="primary-nav__brand-src" />
        </a>
      </nav>
    )
  }
})
