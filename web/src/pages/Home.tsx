import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import HomeTemplate from '../templates/HomeTemplate';

function Home() {
  const [data, setData] = useState({});
  useEffect(() => {
    console.log('useEffect in home is running');
    axios.get('http://localhost:8080/api/v1/courses', {
      params: {
        only_public: true,
        only_mine: false,
      },
    })
      .then((response) => {
        setData(response.data);
        console.log(response.data);
      }, (error) => {
        console.log(error);
      });
  }, []);

  return (
    <div>
      <HomeTemplate data={data} />
      <Link to="/course/detail">CourseDetail</Link>
    </div>
  );
}
export default Home;
