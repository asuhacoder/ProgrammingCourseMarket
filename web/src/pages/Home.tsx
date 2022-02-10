import React from 'react';
import { Link } from 'react-router-dom';
import HomeTemplate from '../templates/HomeTemplate';

function Home() {
  return (
    <div>
      <HomeTemplate />
      <Link to="/course/detail">CourseDetail</Link>
    </div>
  );
}
export default Home;
