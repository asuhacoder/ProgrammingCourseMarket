import React, { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import axios from 'axios';
import { Stack, FormGroup, FormControlLabel, Checkbox } from '@mui/material';
import { useRecoilState } from 'recoil';
import { userState } from '../config/Recoil';
import { User } from '../config/Type';
import CustomTextField from '../atoms/CustomTextField';
import ButtonDiv from '../molecules/ButtonDiv';
import TextStyle from './CourseEditor.css';

function CourseEditor() {
  const { id } = useParams();
  const user: User = useRecoilState(userState)[0];
  const navigate = useNavigate();
  const [title, setTitle] = useState('');
  const [introduction, setIntroduction] = useState('');
  const [titleHasError, setTitleHasError] = useState(false);
  const [titleHelperText, setTitleHelperText] = useState('');
  const [introductionHasError, setIntroductionHasError] = useState(false);
  const [introductionHelperText, setIntroductionHelperText] = useState('');
  const [checked, setChecked] = useState(false);
  useEffect(() => {
    if (id) {
      console.log('useEffect in home is running');
      axios.get(`http://localhost:8080/api/v1/courses/${id}`, {}).then(
        (response) => {
          console.log(response.data);
          setTitle(response.data.title);
          setIntroduction(response.data.introduction);
          setChecked(response.data.is_public);
        },
        (error) => {
          console.log(error);
        },
      );
    }
  }, []);

  const validateTitle = (): boolean => {
    let isValid = true;
    if (title.length < 3) {
      isValid = false;
      setTitleHasError(true);
      setTitleHelperText('length of title must be more than or equal 3');
    } else if (title.length > 50) {
      setTitleHasError(true);
      setTitleHelperText('length of title must be less than or equal 50');
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
  const handleIsPublic = (): void => {
    setChecked(!checked);
  };

  const submitCourseForm = (): void => {
    if (validateTitle() && validateIntroduction()) {
      if (id) {
        axios
          .put(`http://localhost:8080/api/v1/courses/${id}`, {
            token: window.localStorage.getItem('programming-course-market'),
            user_id: user.uuid,
            title,
            introduction,
            image: '',
            price: 0,
            unit_sales: 0,
            is_public: checked,
          })
          .then(
            (response) => {
              console.log(response);
              navigate('/');
            },
            (error) => {
              console.log(error);
            },
          );
      } else {
        axios
          .post('http://localhost:8080/api/v1/courses', {
            token: window.localStorage.getItem('programming-course-market'),
            user_id: user.uuid,
            title,
            introduction,
            image: '',
            price: 0,
            unit_sales: 0,
            is_public: checked,
          })
          .then(
            (response) => {
              console.log(response);
              navigate('/');
            },
            (error) => {
              console.log(error);
            },
          );
      }
    }
  };

  return (
    <Stack justifyContent="flex-start" alignItems="flex-start" spacing={2}>
      <h1 className={TextStyle}>Describe Your Course</h1>
      <CustomTextField
        required
        id="outlined-basic"
        label="Title(required)"
        value={title}
        defaultValue={title}
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
        defaultValue={introduction}
        helperText={introductionHelperText}
        error={introductionHasError}
        onChange={handleIntroductionChange}
        onBlur={validateIntroduction}
        multiline
        minRows={5}
        maxRows={20}
      />
      {checked && (
        <FormGroup>
          <FormControlLabel
            control={<Checkbox checked={checked} defaultChecked onChange={handleIsPublic} />}
            label="Make your course public"
          />
        </FormGroup>
      )}
      {!checked && (
        <FormGroup>
          <FormControlLabel
            control={<Checkbox checked={checked} onChange={handleIsPublic} />}
            label="Make your course public"
          />
        </FormGroup>
      )}
      <ButtonDiv body="Submit" onClick={submitCourseForm} />
    </Stack>
  );
}

export default CourseEditor;
