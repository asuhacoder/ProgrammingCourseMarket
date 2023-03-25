import React from 'react';
import Header from '../organisms/Header';
import CourseList from '../organisms/CourseList';
import HomeDivStyle from './HomeTemplate.css';

function HomeTemplate(props: any) {
  return (
    <div>
      <Header />
      <div className={HomeDivStyle}>
        <CourseList {...props} />
      </div>
    </div>
  );
}
export default HomeTemplate;
