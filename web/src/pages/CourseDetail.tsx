import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import CourseDetailTemplate from '../templates/CourseDetailTemplate';

function CourseDetail() {
  const { id } = useParams();
  const [data, setData] = useState({});
  useEffect(() => {
    console.log('useEffect in home is running');
    axios.get(`http://localhost:8080/api/v1/courses/${id}`)
      .then((response) => {
        setData(response.data);
        console.log(response.data);
      }, (error) => {
        console.log(error);
      });
  }, []);
  return <CourseDetailTemplate course={data} />;
}
export default CourseDetail;
