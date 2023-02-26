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
    axios.get(`http://localhost:8080/api/v1/courses/${id}`)
      .then((response) => {
        setCourse(response.data);
        console.log(response.data);
      }, (error) => {
        console.log(error);
      });
    axios.get(`http://localhost:8080/api/v1/lessons`)
      .then((response) => {
        console.log(response.data);
        setLessons(SortLessons([...response.data.lessons.filter((lesson: Lesson) => lesson.course_id === id)]));
      }, (error) => {
        console.log(error);
      });
    
  }, []);
  return <CourseDetailTemplate lessons={lessons} course={course} />;
}
export default CourseDetail;
