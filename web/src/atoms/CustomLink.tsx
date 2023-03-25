import React from 'react';
import { Link } from 'react-router-dom';

import CustomLinkStyle from './CustomLink.css';

function CustomLink(props: any) {
  return <Link className={CustomLinkStyle} {...props} />;
}

export default CustomLink;
