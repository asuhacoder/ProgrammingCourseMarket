import React, { useState, useEffect } from 'react';
import { SetterOrUpdater, useRecoilState } from 'recoil';
import axios from 'axios';
import MDEditor from '@uiw/react-md-editor';
import rehypeSanitize from 'rehype-sanitize';
import { Snackbar } from '@mui/material';
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import { User, Lesson } from '../config/Type';
import { userState, lessonState } from '../config/Recoil';
import { LessonBodyEditorStyle } from './LessonBodyEditor.css';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function LessonBodyEditor(props: TabPanelProps) {
  const user: User = useRecoilState(userState)[0];
  const [lesson, setLesson]: [Lesson, SetterOrUpdater<Lesson>] = useRecoilState<Lesson>(lessonState);
  const { children, value, index, ...other } = props;
  const [body, setBody] = useState<string | undefined>('');
  const [open, setOpen] = useState(false);
  useEffect(() => {
    setBody(lesson.body);
  }, [lesson]);

  const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(props, ref) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });

  const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen(false);
  };

  const submitBody = (): void => {
    const newLesson: Lesson = {
      uuid: lesson.uuid,
      user_id: user.uuid,
      course_id: lesson.course_id,
      sequence_number: lesson.sequence_number,
      title: lesson.title,
      introduction: lesson.introduction,
      body: body as string,
      default_code: lesson.default_code,
      answer_code: lesson.answer_code,
      language: lesson.language,
    };
    setLesson(newLesson);
    const url = new URL(location.href);
    const instance = axios.create({baseURL: `${url.protocol}//${url.hostname}:8080`})
    instance
      .put(`/api/v1/lessons/${lesson.uuid}`, {
        token: window.localStorage.getItem('programming-course-market'),
        ...newLesson,
      })
      .then(
        (response) => {
          console.log(response);
          setOpen(true);
        },
        (error) => {
          console.log(error);
        },
      );
  };

  return (
    <div
      className={LessonBodyEditorStyle}
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <div className={LessonBodyEditorStyle}>
          <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success" sx={{ width: '100%' }}>
              Saved Changes
            </Alert>
          </Snackbar>
          <MDEditor
            className={LessonBodyEditorStyle}
            height={600}
            value={body}
            onChange={setBody}
            onBlur={submitBody}
            previewOptions={{
              rehypePlugins: [[rehypeSanitize]],
            }}
          />
        </div>
      )}
    </div>
  );
}
export default LessonBodyEditor;
