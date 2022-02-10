import React from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import HomeTemplate from '../templates/HomeTemplate';

function Home() {
  axios.get('http://localhost:8080/api/v1/courses')
    .then((response) => {
      console.log(response.data);
    }, (error) => {
      console.log(error);
    });
  return (
    <div>
      <HomeTemplate />
      <Link to="/course/detail">CourseDetail</Link>
    </div>
  );
}
export default Home;
