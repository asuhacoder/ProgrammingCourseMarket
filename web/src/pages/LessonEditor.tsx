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
      console.log('id: ', id);
      axios.get(`http://localhost:8080/api/v1/lessons/${id}`)
      .then((response) => {
          setLesson(response.data)
        //   setLanguage(lesson.language.split('@')[0]);
        //   setVersion(lesson.language.split('@')[1]);
          console.log('response.data: ', response.data);
      }, (error) => {
          console.log(error);
      });
  }, [id]);
  return <LessonEditorTemplate />;
}
export default LessonEditor;
