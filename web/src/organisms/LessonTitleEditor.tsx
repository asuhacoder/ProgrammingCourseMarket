import React, { useState } from 'react';
import { SetterOrUpdater, useRecoilState } from 'recoil';
import axios from 'axios';
import { Snackbar } from '@mui/material';
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import { User, Lesson } from '../config/Type';
import { userState, lessonState } from '../config/Recoil';
import CustomTextField from '../atoms/CustomTextField';

function LessonTitleEditor() {
  const user: User = useRecoilState(userState)[0];
  const [lesson, setLesson]: [Lesson, SetterOrUpdater<Lesson>] = useRecoilState<Lesson>(lessonState);
  const [titleHasError, setTitleHasError] = useState(false);
  const [titleHelperText, setTitleHelperText] = useState('');
  const [open, setOpen] = React.useState(false);

  const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(props, ref) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });

  const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen(false);
  };

  const validateTitle = (): boolean => {
    let isValid = true;
    if (lesson.title.length < 3) {
      isValid = false;
      setTitleHasError(true);
      setTitleHelperText('length of title must be more than or equal 3');
    } else if (lesson.title.length > 50) {
      setTitleHasError(true);
      setTitleHelperText('length of title must be less than or equal 50');
    } else {
      setTitleHasError(false);
      setTitleHelperText('');
    }
    return isValid;
  };
  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    let tmp: Lesson = Object.assign({}, lesson);
    tmp.title = e.target.value;
    setLesson(tmp);
  };

  const submitTitle = (): void => {
    if (validateTitle()) {
      const url = new URL(location.href);
      const instance = axios.create({baseURL: `${url.protocol}//${url.hostname}:8080`})
      instance
        .put(`/api/v1/lessons/${lesson.uuid}`, {
          token: window.localStorage.getItem('programming-course-market'),
          user_id: user.uuid,
          course_id: lesson.course_id,
          sequence_number: lesson.sequence_number,
          title: lesson.title,
          introduction: lesson.introduction,
          body: lesson.body,
          default_code: lesson.default_code,
          answer_code: lesson.answer_code,
          language: lesson.language,
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
    }
  };

  return (
    <div>
      <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success" sx={{ width: '100%' }}>
          Saved Changes
        </Alert>
      </Snackbar>
      <CustomTextField
        required
        id="outlined-required"
        label="Title"
        value={lesson.title}
        helperText={titleHelperText}
        error={titleHasError}
        onChange={handleTitleChange}
        onBlur={submitTitle}
      />
    </div>
  );
}

export default LessonTitleEditor;
