import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { defaultLesson, defaultCases } from '../config/Type';
import LessonDetailTemplate from '../templates/LessonDetailTemplate';

function LessonDetail() {
  const { id } = useParams();
  const [lesson, setLesson] = useState(defaultLesson);
  const [cases, setCases] = useState(defaultCases);
  console.log('LessonDetail in page');
  useEffect(() => {
    console.log('useEffect in LessonDetail is running');
    console.log('id: ', id);
    axios.get(`http://localhost:8080/api/v1/lessons/${id}`)
    .then((response) => {
      setLesson(response.data);
      console.log('response.data: ', response.data);
    }, (error) => {
      console.log(error);
    });
    axios.get('http://localhost:8080/api/v1/cases', {
      params: {
        lesson_id: id,
      },
    })
    .then((response) => {
      setCases(response.data.cases);
      console.log('response.data: ', response.data.cases);
    }, (error) => {
        console.log(error);
    });
  }, [id]);
  return <LessonDetailTemplate lesson={lesson} cases={cases} />;
}
export default LessonDetail;
