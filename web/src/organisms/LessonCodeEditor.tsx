import React, { useState, useEffect, useRef } from 'react';
import { SetterOrUpdater, useRecoilState } from 'recoil';
import axios from 'axios';
import Editor from "@monaco-editor/react";
import { Snackbar } from '@mui/material';
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import { User, Lesson } from '../config/Type';
import { userState, lessonState } from '../config/Recoil';
import data from '../config/language.json';
import { LessonCodeEditorStyle } from './LessonCodeEditor.css';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
  defaultOrAnswer: keyof Lesson;
}

function LessonCodeEditor(props: TabPanelProps) {
  const languageList = data as any;
  const user: User = useRecoilState(userState)[0];
  const [lesson, setLesson]:[Lesson, SetterOrUpdater<Lesson>] = useRecoilState<Lesson>(lessonState);
  const { children, value, index, defaultOrAnswer, ...other } = props;
	const [code , setCode] = useState('');
  const [codeHasError, setCodeHasError] = useState(false);
  const [codeHelperText, setCodeHelperText] = useState('');
  const [open, setOpen] = React.useState(false);
  const options = {
    readOnly: false,
    minimap: { enabled: false },
  };

  useEffect(() => {
    setCode(lesson[defaultOrAnswer] as string);
  }, [lesson]);

  const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref,
  ) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });

  const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen(false);
  };

  function handleEditorChange(value: string | undefined, event: any) {
    setCode(value as string);
  }

  const submitCode = (): void => {
    const newLesson:Lesson = {
      uuid: lesson.uuid,
      user_id: user.uuid,
      course_id: lesson.course_id,
      sequence_number: lesson.sequence_number,
      title: lesson.title,
      introduction: lesson.introduction,
      body: lesson.body,
      default_code: lesson.default_code,
      answer_code: lesson.answer_code,
      language: lesson.language,
    };
    newLesson[defaultOrAnswer] = code as never;
    setLesson(newLesson);
    axios.put(`http://localhost:8080/api/v1/lessons/${lesson.uuid}`, {
      token: window.localStorage.getItem('programming-course-market'),
      ...newLesson,
    })
      .then((response) => {
        console.log(response);
        setOpen(true);
      }, (error) => {
        console.log(error);
      });
  };

  return (
    <div
      className={LessonCodeEditorStyle}
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <div>
          <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success" sx={{ width: '100%' }}>
              Saved Changes
            </Alert>
          </Snackbar>
          <div onBlur={submitCode}>
          <Editor
            height="70vh"
            defaultLanguage={languageList[lesson.language.split('@')[0]].monaco}
            defaultValue={code}
            onChange={handleEditorChange}
            options={options}
          />
          </div>
        </div>
      )}
    </div>
  );
}

export default LessonCodeEditor;
