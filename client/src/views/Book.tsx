import React from 'react';
import { RouteComponentProps } from '@reach/router';

interface BookProps extends RouteComponentProps {
  bookId?: string;
}

const Book: React.FC<BookProps> = (props) => {
  return <h1>Book: {props.bookId}</h1>;
};

export default Book;
