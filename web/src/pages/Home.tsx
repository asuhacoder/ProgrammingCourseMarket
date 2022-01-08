import React from 'react';
import { Link, useLocation } from 'react-router-dom';

function Home() {
  const location = useLocation();
  return <Link to="/signup" state={{ from: location }}>Signup</Link>;
}

export default Home;
