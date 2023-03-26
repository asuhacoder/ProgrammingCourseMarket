import React, { useState } from 'react';
import axios from 'axios';
import { Stack } from '@mui/material';
import { useNavigate, useLocation } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { userState } from '../config/Recoil';
import StackStyle from './SignupForm.css';
import CustomTextField from '../atoms/CustomTextField';
import ButtonDiv from '../molecules/ButtonDiv';

interface State {
  from: Location;
}

function SignupForm() {
  const setUser = useRecoilState(userState)[1];

  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [introduction, setIntroduction] = useState('');
  const [password, setPassword] = useState('');
  const [nameHasError, setNameHasError] = useState(false);
  const [nameHelperText, setNameHelperText] = useState('');
  const [emailHasError, setEmailHasError] = useState(false);
  const [emailHelperText, setEmailHelperText] = useState('');
  const [introductionHasError, setIntroductionHasError] = useState(false);
  const [introductionHelperText, setIntroductionHelperText] = useState('');
  const [passwordHasError, setPasswordHasError] = useState(false);
  const [passwordHelperText, setPasswordHelperText] = useState('');

  const navigate = useNavigate();
  const routerLocation = useLocation();

  const validateName = (): boolean => {
    let isValid = true;
    if (name.length < 3) {
      isValid = false;
      setNameHasError(true);
      setNameHelperText('length of name must be more than or equal 3');
    } else if (name.length > 20) {
      setNameHasError(true);
      setNameHelperText('length of name must be less than or equal 20');
    } else {
      setNameHasError(false);
      setNameHelperText('');
    }
    return isValid;
  };
  const validateEmail = (): boolean => {
    let isValid = true;
    const re = /^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[A-Za-z]+$/;
    if (!re.test(email)) {
      isValid = false;
      setEmailHasError(true);
      setEmailHelperText('this is not email address');
    } else {
      setEmailHasError(false);
      setEmailHelperText('');
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
  const validatePassword = (): boolean => {
    let isValid = true;
    if (password.length < 8) {
      isValid = false;
      setPasswordHasError(true);
      setPasswordHelperText('length of password must be less than or equal 8');
    } else {
      setPasswordHasError(false);
      setPasswordHelperText('');
    }
    return isValid;
  };

  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setName(e.target.value);
  };
  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setEmail(e.target.value);
  };
  const handleIntroductionChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setIntroduction(e.target.value);
  };
  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setPassword(e.target.value);
  };

  const submitSignupForm = (): void => {
    if (validateName() && validateEmail() && validateIntroduction() && validatePassword()) {
      const url = new URL(location.href);
      const instance = axios.create({baseURL: `${url.protocol}//${url.hostname}:8080`})
      instance
        .post(`/api/v1/users`, {
          name,
          email,
          password,
          introduction,
        })
        .then(
          (response) => {
            console.log(response);
            window.localStorage.setItem('programming-course-market', response.data.token);
            setUser({
              token: response.data.token,
              uuid: response.data.uuid,
              name: response.data.name,
              email: response.data.email,
              introduction: response.data.introduction,
            });
            const state = routerLocation.state as State;
            const from = state.from.pathname || '/';
            console.log('from: ', from);
            navigate(from, { replace: true });
          },
          (error) => {
            console.log(error);
            setEmailHasError(true);
            setEmailHelperText('this email address is already used');
          },
        );
    }
  };

  return (
    <Stack
      id="organisms-signup-form-stack"
      className={StackStyle}
      justifyContent="flex-start"
      alignItems="flex-start"
      spacing={2}
    >
      <CustomTextField
        required
        id="outlined-required"
        label="Name(required)"
        value={name}
        helperText={nameHelperText}
        error={nameHasError}
        onChange={handleNameChange}
        onBlur={validateName}
      />
      <CustomTextField
        required
        id="outlined-required"
        label="Email(required)"
        value={email}
        helperText={emailHelperText}
        error={emailHasError}
        onChange={handleEmailChange}
        onBlur={validateEmail}
      />
      <CustomTextField
        id="outlined-basic"
        label="Introduction"
        value={introduction}
        helperText={introductionHelperText}
        error={introductionHasError}
        onChange={handleIntroductionChange}
        onBlur={validateIntroduction}
        multiline
        maxRows={20}
      />
      <CustomTextField
        required
        type="password"
        id="outlined-password-input"
        label="Password(required)"
        value={password}
        helperText={passwordHelperText}
        error={passwordHasError}
        onChange={handlePasswordChange}
        onBlur={validatePassword}
      />
      <ButtonDiv body="Submit" onClick={submitSignupForm} />
    </Stack>
  );
}
export default SignupForm;
