import React from 'react';
import Header from '../organisms/Header';
import CourseEditor from '../organisms/CourseEditor';
import CourseUpdaterDiv from './CourseUpdaterTemplate.css';

function CourseUpdaterTemplate() {
  return (
    <div>
      <Header />
      <div className={CourseUpdaterDiv}>
        <CourseEditor />
      </div>
    </div>
  );
}
export default CourseUpdaterTemplate;
