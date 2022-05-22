import React from 'react';
import Header from '../organisms/Header';
import CourseEditor from '../organisms/CourseEditor';
import CourseCreatorDiv from './CourseCreatorTemplate.css';

function CourseCreatorTemplate() {
  return (
    <div>
      <Header />
      <div className={CourseCreatorDiv}>
        <CourseEditor />
      </div>
    </div>
  );
}
export default CourseCreatorTemplate;
