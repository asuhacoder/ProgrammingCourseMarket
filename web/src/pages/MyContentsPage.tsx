import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useRecoilState } from 'recoil';
import { userState } from '../config/Recoil';
import MyContentsTemplate from '../templates/MyContentsTemplate';

function MyContentsPage() {
  const [courses, setCourses] = useState([]);
  const user = useRecoilState(userState)[0];
  console.log('user in my contents pages component', user);
  useEffect(() => {
    console.log('useEffect in mycontents is running', user);
    axios.get('http://localhost:8080/api/v1/courses', {
      data: {
        user_id: user.uuid,
        only_public: false,
        only_mine: true,
      },
    })
      .then((response) => {
        setCourses(response.data.courses);
        console.log('suceed', response.data);
      }, (error) => {
        console.log('failed', error);
        console.log(error);
      });
  }, []);
  console.log('useEffect in mycontents ran');
  return <MyContentsTemplate courses={courses} />;
}
export default MyContentsPage;
