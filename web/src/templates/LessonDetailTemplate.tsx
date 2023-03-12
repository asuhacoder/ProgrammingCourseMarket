import React from 'react';
import Header from '../organisms/Header';
import LessonDetail from '../organisms/LessonDetail';
import LessonDetailDivStyle from './LessonDetailTemplate.css';

function LessonDetailTemplate(props: any) {
  return (
    <div>
      <Header />
      <div className={LessonDetailDivStyle}>
        <LessonDetail
          {...props}
        />
      </div>
    </div>
  );
}
export default LessonDetailTemplate;
