import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import { ThemeProvider } from '@mui/material/styles';
import { RecoilRoot } from 'recoil';
import App from './App';
import Theme from './config/Theme';
import reportWebVitals from './reportWebVitals';

ReactDOM.render(
  <BrowserRouter>
    <RecoilRoot>
      <ThemeProvider theme={Theme}>
        <App />
      </ThemeProvider>
    </RecoilRoot>
  </BrowserRouter>,
  document.getElementById('root'),
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
