import React from 'react';
import Header from '../organisms/Header';
import CourseDetail from '../organisms/CourseDetail';
import CourseDetailDiv from './CourseDetailTemplate.css';

function CourseDetailTemplate(props: any) {
  return (
    <div>
      <Header />
      <div className={CourseDetailDiv}>
        <CourseDetail {...props} />
      </div>
    </div>
  );
}
export default CourseDetailTemplate;
