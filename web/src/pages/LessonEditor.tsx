import React, { useState, useEffect } from 'react';
import { SetterOrUpdater, useRecoilState } from 'recoil';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { Lesson } from '../config/Type';
import { lessonState } from '../config/Recoil';
import LessonEditorTemplate from '../templates/LessonEditorTemplate';

function LessonEditor() {
  const { id } = useParams();
  const [lesson, setLesson]:[Lesson, SetterOrUpdater<Lesson>] = useRecoilState<Lesson>(lessonState);
  useEffect(() => {
      console.log('useEffect in LessonEditor is running');
      axios.get(`http://localhost:8080/api/v1/lessons/${id}`)
      .then((response) => {
          setLesson(response.data)
          console.log(response.data);
      }, (error) => {
          console.log(error);
      });
  }, []);
  return <LessonEditorTemplate />;
}
export default LessonEditor;
