import React from 'react';
import { Button } from '@mui/material';

import ButtonDivStyle from './ButtonDiv.css';

function ButtonDiv(props: any) {
  const { body } = props;
  return (
    <div className={ButtonDivStyle}>
      <Button
        color="primary"
        variant="contained"
        {...props}
      >
        {body}
      </Button>
    </div>

  );
}

export default ButtonDiv;
