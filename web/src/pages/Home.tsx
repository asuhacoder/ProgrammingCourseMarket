import React from 'react';
import { Link, useLocation } from 'react-router-dom';

function Home() {
  const location = useLocation();
  return (
    <div>
      <Link to="/signup" state={{ from: location }}>Signup</Link>
      <br />
      <Link to="/login" state={{ from: location }}>Login</Link>
    </div>
  );
}

export default Home;
