import React, { useState } from 'react';
import axios from 'axios';
import { Stack } from '@mui/material';
import { useNavigate, useLocation } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { userState } from '../config/Recoil';
import StackStyle from './LoginForm.css';
import CustomTextField from '../atoms/CustomTextField';
import ButtonDiv from '../molecules/ButtonDiv';

interface State {
  from: Location;
}

function LoginForm() {
  const setUser = useRecoilState(userState)[1];

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [emailHasError, setEmailHasError] = useState(false);
  const [emailHelperText, setEmailHelperText] = useState('');
  const [passwordHasError, setPasswordHasError] = useState(false);
  const [passwordHelperText, setPasswordHelperText] = useState('');

  const navigate = useNavigate();
  const routerLocation = useLocation();

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

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setEmail(e.target.value);
  };
  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setPassword(e.target.value);
  };

  const submitLoginForm = (): void => {
    if (validateEmail() && validatePassword()) {
      const instance = axios.create({baseURL: process.env.REACT_APP_API_URL})
      instance
        .post(`/api/v1/authn`, {
          email,
          password,
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
            setEmailHelperText('email address or password is incorrect');
          },
        );
    }
  };

  return (
    <Stack
      id="organisms-login-form-stack"
      className={StackStyle}
      justifyContent="flex-start"
      alignItems="flex-start"
      spacing={2}
    >
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
      <ButtonDiv body="Submit" onClick={submitLoginForm} />
    </Stack>
  );
}
export default LoginForm;
