import React from 'react';
import Header from '../organisms/Header';
import CourseList from '../organisms/CourseList';
import HomeDivStyle from './HomeTemplate.css';

function HomeTemplate() {
  return (
    <div>
      <Header />
      <div className={HomeDivStyle}>
        <CourseList />
      </div>
    </div>
  );
}
export default HomeTemplate;
