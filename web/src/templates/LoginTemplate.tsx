import React from 'react';
import Header from '../organisms/Header';
import LoginForm from '../organisms/LoginForm';
import LoginDivStyle from './LoginTemplate.css';

function LoginTemplate() {
  return (
    <div>
      <Header />
      <div className={LoginDivStyle}>
        <LoginForm />
      </div>
    </div>
  );
}
export default LoginTemplate;
