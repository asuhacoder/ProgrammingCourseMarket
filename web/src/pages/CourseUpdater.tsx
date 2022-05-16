import axios from 'axios';
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { defaultCourse } from '../config/Type';
import CourseUpdaterTemplate from '../templates/CourseUpdaterTemplate';

function CourseUpdater() {
  const { id } = useParams();
  const [data, setData] = useState(defaultCourse);
  useEffect(() => {
    console.log('useEffect in home is running');
    axios.get(`http://localhost:8080/api/v1/courses/${id}`, {})
      .then((response) => {
        setData(response.data);
        console.log('course data: ', data);
        console.log(response.data);
      }, (error) => {
        console.log(error);
      });
  }, []);
  return (
    <CourseUpdaterTemplate />
  );
}
export default CourseUpdater;
