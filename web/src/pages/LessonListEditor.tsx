import React, { useEffect } from 'react';
import { SetterOrUpdater, useRecoilState } from 'recoil';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { Lesson } from '../config/Type';
import { lessonsState } from '../config/Recoil';
import LessonListEditorTemplate from '../templates/LessonListEditorTemplate';
import { SortLessons } from '../utils/LessonsUtils';

function LessonListEditor() {
  const { id } = useParams();
  const [, setLessons]: [Lesson[], SetterOrUpdater<Lesson[]>] = useRecoilState<Lesson[]>(lessonsState);
  useEffect(() => {
    console.log('useEffect in LessonListEditor is running');
    axios
      .get(`http://localhost:8080/api/v1/lessons`, {
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
  return <LessonListEditorTemplate />;
}
export default LessonListEditor;
