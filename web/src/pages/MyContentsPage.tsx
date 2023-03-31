import React, { useEffect } from 'react';
import axios from 'axios';
import { useRecoilState } from 'recoil';
import { userState, myContentsState } from '../config/Recoil';
import MyContentsTemplate from '../templates/MyContentsTemplate';

function MyContentsPage() {
  const user = useRecoilState(userState)[0];
  const setMyContents = useRecoilState(myContentsState)[1];
  console.log('user.uuid: ', user.uuid);
  console.log('user in my contents pages component', user);
  useEffect(() => {
    console.log('useEffect in mycontents is running', user);
    const instance = axios.create({baseURL: process.env.REACT_APP_API_URL})
    instance
      .get(`/api/v1/courses`, {
        params: {
          user_id: user.uuid,
          only_public: false,
          only_mine: true,
        },
      })
      .then(
        (response) => {
          setMyContents(response.data.courses);
          console.log('suceed', response.data);
        },
        (error) => {
          console.log('failed', error);
          console.log(error);
        },
      );
  }, []);
  console.log('useEffect in mycontents ran');
  return <MyContentsTemplate />;
}
export default MyContentsPage;
