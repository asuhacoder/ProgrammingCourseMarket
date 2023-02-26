import React from 'react';
import Header from '../organisms/Header';
import LessonListEditor from '../organisms/LessonListEditor';
import LessonListEditorDivStyle from './LessonListEditorTemplate.css';

function LessonListEditorTemplate() {
  return (
    <div>
      <Header />
      <div className={LessonListEditorDivStyle}>
        <LessonListEditor />
      </div>
    </div>
  );
}
export default LessonListEditorTemplate;
