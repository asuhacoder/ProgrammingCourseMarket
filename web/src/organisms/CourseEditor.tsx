import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { Stack } from '@mui/material';
import CustomTextField from '../atoms/CustomTextField';
import ButtonDiv from '../molecules/ButtonDiv';
import TextStyle from './CourseEditor.css';

function CourseEditor() {
  const navigate = useNavigate();
  const [title, setTitle] = useState('');
  const [introduction, setIntroduction] = useState('');
  const [titleHasError, setTitleHasError] = useState(false);
  const [titleHelperText, setTitleHelperText] = useState('');
  const [introductionHasError, setIntroductionHasError] = useState(false);
  const [introductionHelperText, setIntroductionHelperText] = useState('');

  const validateTitle = (): boolean => {
    let isValid = true;
    if (title.length < 3) {
      isValid = false;
      setTitleHasError(true);
      setTitleHelperText('length of title must be more than or equal 3');
    } else if (title.length > 20) {
      setTitleHasError(true);
      setTitleHelperText('length of title must be less than or equal 20');
    } else {
      setTitleHasError(false);
      setTitleHelperText('');
    }
    return isValid;
  };
  const validateIntroduction = (): boolean => {
    let isValid = true;
    if (introduction.length > 500) {
      isValid = false;
      setIntroductionHasError(true);
      setIntroductionHelperText('length of introduction must be less than or equal 500');
    } else {
      setIntroductionHasError(false);
      setIntroductionHelperText('');
    }
    return isValid;
  };

  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setTitle(e.target.value);
  };
  const handleIntroductionChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setIntroduction(e.target.value);
  };

  const submitCourseForm = (): void => {
    if (validateTitle() && validateIntroduction()) {
      axios.post('http://localhost:8080/api/v1/courses', {
        token: window.localStorage.getItem('programming-course-market'),
        title,
        introduction,
        image: '',
        price: 0,
        unit_sales: 0,
        is_public: false,
      })
        .then((response) => {
          console.log(response);
          navigate('/');
        }, (error) => {
          console.log(error);
        });
    }
  };

  return (
    <Stack
      justifyContent="flex-start"
      alignItems="flex-start"
      spacing={2}
    >
      <h1 className={TextStyle}>Describe Your Course</h1>
      <CustomTextField
        required
        id="outlined-basic"
        label="Title(required)"
        value={title}
        helperText={titleHelperText}
        error={titleHasError}
        onChange={handleTitleChange}
        onBlur={validateTitle}
      />
      <CustomTextField
        required
        id="outlined-basic"
        label="Introduction(required)"
        value={introduction}
        helperText={introductionHelperText}
        error={introductionHasError}
        onChange={handleIntroductionChange}
        onBlur={validateIntroduction}
        multiline
        minRows={5}
        maxRows={20}
      />
      <ButtonDiv
        body="Submit"
        onClick={submitCourseForm}
      />
    </Stack>
  );
}

export default CourseEditor;
