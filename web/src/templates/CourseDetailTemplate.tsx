import React from 'react';
import Header from '../organisms/Header';
import CourseDetailForm from '../organisms/CourseDetail';
import CourseDetailDiv from './CourseDetailTemplate.css';

function CourseDetailTemplate() {
  return (
    <div>
      <Header />
      <div className={CourseDetailDiv}>
        <CourseDetailForm />
      </div>
    </div>
  );
}
export default CourseDetailTemplate;
