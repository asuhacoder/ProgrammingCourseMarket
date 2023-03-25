import React, { useState, useEffect } from 'react';
import axios from 'axios';
import HomeTemplate from '../templates/HomeTemplate';

function Home() {
  const [data, setData] = useState({});
  useEffect(() => {
    console.log('useEffect in home is running');
    axios
      .get('http://localhost:8080/api/v1/courses', {
        params: {
          only_public: true,
          only_mine: false,
        },
      })
      .then(
        (response) => {
          setData(response.data.courses);
          console.log(response.data.courses);
        },
        (error) => {
          console.log(error);
        },
      );
  }, []);

  return (
    <div>
      <HomeTemplate courses={data} />
    </div>
  );
}
export default Home;
