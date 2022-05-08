import React from 'react';
import { Routes, Route } from 'react-router-dom';
import axios from 'axios';
import { useRecoilState } from 'recoil';
import { userState, defaultUser } from './config/Recoil';
import Home from './pages/Home';
import Signup from './pages/Signup';
import Login from './pages/Login';
import CourseDetail from './pages/CourseDetail';
import CourseEditor from './pages/CourseEditor';
import MyContentsPage from './pages/MyContentsPage';

function App() {
  console.log('token', window.localStorage.getItem('programming-course-market'));
  const [user, setUser] = useRecoilState(userState);
  if (JSON.stringify(user) === JSON.stringify(defaultUser)) {
    const token = window.localStorage.getItem('programming-course-market');
    axios.post('http://localhost:8080/api/v1/authz', {
      token,
    })
      .then((response) => {
        console.log(response);
        window.localStorage.setItem('programming-course-market', response.data.token);
        setUser({
          token: response.data.token,
          uuid: response.data.uuid,
          name: response.data.name,
          email: response.data.email,
          introduction: response.data.introduction,
        });
      }, (error) => {
        console.log(error);
      });
  }
  console.log('user', user);

  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
        <Route path="/course/detail" element={<CourseDetail />} />
        <Route path="/course/editor" element={<CourseEditor />} />
        <Route path="/mycontents" element={<MyContentsPage />} />
      </Routes>
    </div>
  );
}

export default App;
