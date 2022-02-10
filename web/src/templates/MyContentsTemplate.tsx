import React from 'react';
import Header from '../organisms/Header';
import CourseList from '../organisms/MyContents';
import MyContentsDivStyle from './MyContentsTemplate.css';

function MyContentsTemplate() {
  return (
    <div>
      <Header />
      <div className={MyContentsDivStyle}>
        <CourseList />
      </div>
    </div>
  );
}
export default MyContentsTemplate;
