import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { Lesson } from '../config/Type';
import CourseDetailTemplate from '../templates/CourseDetailTemplate';
import { SortLessons } from '../utils/LessonsUtils';

function CourseDetail() {
  const { id } = useParams();
  const [course, setCourse] = useState({});
  const [lessons, setLessons] = useState<Lesson[]>([]);
  useEffect(() => {
    console.log('useEffect in CourseDetail is running');
    const instance = axios.create({baseURL: process.env.REACT_APP_API_URL})
    instance.get(`/api/v1/courses/${id}`).then(
      (response) => {
        setCourse(response.data);
        console.log(response.data);
      },
      (error) => {
        console.log(error);
      },
    );
    instance
      .get(`/api/v1/lessons`, {
        params: {
          course_id: id,
        },
      })
      .then(
        (response) => {
          console.log(response.data);
          setLessons(SortLessons(response.data.lessons));
        },
        (error) => {
          console.log(error);
        },
      );
  }, []);
  return <CourseDetailTemplate lessons={lessons} course={course} />;
}
export default CourseDetail;
