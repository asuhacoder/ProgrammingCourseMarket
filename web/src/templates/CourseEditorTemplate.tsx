import React from 'react';
import Header from '../organisms/Header';
import CourseEditorForm from '../organisms/CourseEditor';
import CourseEditorDiv from './CourseEditorTemplate.css';

function CourseEditorTemplate() {
  return (
    <div>
      <Header />
      <div className={CourseEditorDiv}>
        <CourseEditorForm />
      </div>
    </div>
  );
}
export default CourseEditorTemplate;
