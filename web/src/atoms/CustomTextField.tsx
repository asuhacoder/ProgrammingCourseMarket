import React from 'react';
import { TextField } from '@mui/material';

import CustomTextFieldStyle from './CustomTextField.css';

function CustomTextField(props: any) {
  return (
    <TextField
      className={CustomTextFieldStyle}
      {...props}
    />
  );
}

export default CustomTextField;
