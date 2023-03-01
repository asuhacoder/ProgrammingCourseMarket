import React from 'react';
import Header from '../organisms/Header';
import LessonEditor from '../organisms/LessonEditor';
import LessonEditorDivStyle from './LessonEditorTemplate.css';

function LessonEditorTemplate() {
  return (
    <div>
      <Header />
      <div className={LessonEditorDivStyle}>
        <LessonEditor />
      </div>
    </div>
  );
}
export default LessonEditorTemplate;
