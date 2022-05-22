import React from 'react';
import Header from '../organisms/Header';
import CourseDetailForm from '../organisms/CourseDetail';
import CourseDetailDiv from './CourseDetailTemplate.css';

function CourseDetailTemplate(props: any) {
  return (
    <div>
      <Header />
      <div className={CourseDetailDiv}>
        <CourseDetailForm
          {...props}
        />
      </div>
    </div>
  );
}
export default CourseDetailTemplate;
