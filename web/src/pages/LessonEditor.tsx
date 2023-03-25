import React, { useEffect } from 'react';
import { SetterOrUpdater, useRecoilState } from 'recoil';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { Lesson, Case } from '../config/Type';
import { lessonState, casesState } from '../config/Recoil';
import LessonEditorTemplate from '../templates/LessonEditorTemplate';

function LessonEditor() {
  const { id } = useParams();
  const [, setLesson]: [Lesson, SetterOrUpdater<Lesson>] = useRecoilState<Lesson>(lessonState);
  const [, setCases]: [Case[], SetterOrUpdater<Case[]>] = useRecoilState<Case[]>(casesState);
  useEffect(() => {
    console.log('useEffect in LessonEditor is running');
    console.log('id: ', id);
    axios.get(`http://localhost:8080/api/v1/lessons/${id}`).then(
      (response) => {
        setLesson(response.data);
        console.log('response.data: ', response.data);
      },
      (error) => {
        console.log(error);
      },
    );
    axios
      .get('http://localhost:8080/api/v1/cases', {
        params: {
          lesson_id: id,
        },
      })
      .then(
        (response) => {
          setCases(response.data.cases);
          console.log('response.data: ', response.data.cases);
        },
        (error) => {
          console.log(error);
        },
      );
  }, [id]);
  return <LessonEditorTemplate />;
}
export default LessonEditor;
