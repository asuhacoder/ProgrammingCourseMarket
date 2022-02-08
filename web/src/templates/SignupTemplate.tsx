import React from 'react';
import Header from '../organisms/Header';
import SignupForm from '../organisms/SignupForm';
import SignupDivStyle from './SignupTemplate.css';

function SignupTemplate() {
  return (
    <div>
      <Header />
      <div className={SignupDivStyle}>
        <SignupForm />
      </div>
    </div>
  );
}
export default SignupTemplate;
