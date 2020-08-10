import './App.css';

import { Book, Electronics, Home, NotFound } from './views';
import { Desktop, Laptop, Mobile } from './views/Electronics';
import { Link, Router } from '@reach/router';

import React from 'react';
import logo from './logo.svg';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <nav>
          <Link className="App-link" to="/">
            Home
          </Link>{' '}
          |{' '}
          <Link className="App-link" to="/book/1">
            Book
          </Link>{' '}
          |{' '}
          <Link className="App-link" to="/electronics">
            Electronics
          </Link>
        </nav>

        <Router>
          <Home path="/" />
          <Book path="/book/:bookId" />
          <Electronics path="/electronics">
            <Mobile path="mobile" />
            <Desktop path="desktop" />
            <Laptop path="laptop" />
          </Electronics>
          <NotFound default />
        </Router>
      </header>
    </div>
  );
}

export default App;
