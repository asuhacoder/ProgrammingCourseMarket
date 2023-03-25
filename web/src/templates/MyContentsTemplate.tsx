import React from 'react';
import Header from '../organisms/Header';
import MyContents from '../organisms/MyContents';
import MyContentsDivStyle from './MyContentsTemplate.css';

function MyContentsTemplate(props: any) {
  return (
    <div>
      <Header />
      <div className={MyContentsDivStyle}>
        <MyContents {...props} />
      </div>
    </div>
  );
}
export default MyContentsTemplate;
