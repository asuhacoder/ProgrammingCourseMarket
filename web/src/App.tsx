import React from 'react';
import { Routes, Route } from 'react-router-dom';
import { RecoilRoot } from 'recoil';
import { ThemeProvider } from '@mui/material/styles';
import Theme from './config/Theme';
import Home from './pages/Home';
import Signup from './pages/Signup';
import Login from './pages/Login';

function App() {
  return (
    <div className="App">
      <RecoilRoot>
        <ThemeProvider theme={Theme}>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/signup" element={<Signup />} />
            <Route path="/login" element={<Login />} />
          </Routes>
        </ThemeProvider>
      </RecoilRoot>
    </div>
  );
}

export default App;
