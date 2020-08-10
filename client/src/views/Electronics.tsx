import { Link, RouteComponentProps } from '@reach/router';

import React from 'react';

export const Mobile: React.FC<RouteComponentProps> = () => {
  return <h3>Mobile Phone</h3>;
};

export const Desktop: React.FC<RouteComponentProps> = () => {
  return <h3>Desktop PC</h3>;
};

export const Laptop: React.FC<RouteComponentProps> = () => {
  return <h3>Laptop</h3>;
};

const Electronics: React.FC<RouteComponentProps> = (props) => {
  return (
    <>
      <h1>Electronics</h1>
      <ul>
        <li className="App-list-item">
          <Link className="App-link" to="/electronics/mobile">
            Mobile
          </Link>
        </li>
        <li className="App-list-item">
          <Link className="App-link" to="/electronics/desktop">
            Desktop
          </Link>
        </li>
        <li className="App-list-item">
          <Link className="App-link" to="/electronics/laptop">
            Laptop
          </Link>
        </li>
      </ul>

      {/* nested component will show here */}
      {props.children}
    </>
  );
};

export default Electronics;
