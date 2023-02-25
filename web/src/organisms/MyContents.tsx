import React from 'react';
import axios from 'axios';
import {
  Button,
  Card,
  CardActions,
  CardContent,
  Typography,
} from '@mui/material';
import { useRecoilState } from 'recoil';
import { userState, myContentsState } from '../config/Recoil';
import { User, Course } from '../config/Type';
import {
  NoContentsStyle,
  MyContentsStyle,
  CardContentStyle,
  TypographyStyle,
  ButtonDivStyle,
  ButtonStyle,
} from './MyContents.css';
import CustomLink from '../atoms/CustomLink';
import AlertModal from '../molecules/AlertModal';

function MyContents() {
  const user: User = useRecoilState(userState)[0];
  const [myContents, setMyContents] = useRecoilState(myContentsState);
  const deleteCourse = (uuid: string):void => {
    axios.delete(`http://localhost:8080/api/v1/courses/${uuid}`, {
      data: {
        user_id: user.uuid,
      },
    })
      .then((response) => {
        setMyContents(myContents.filter((myContent: Course) => myContent.uuid !== uuid));
        console.log(response);
      }, (error) => {
        console.log(error);
      });
  };

  return (
    <div className={MyContentsStyle}>
      {myContents === null && (
        <div className={NoContentsStyle}>No Contents</div>
      )}
      {myContents !== null && myContents.map((myContent: Course) => (
        <Card key={myContent.uuid} className={CardContentStyle}>
          <CardContent>
            <Typography variant="h4" component="div">
              {myContent.title}
            </Typography>
            <Typography variant="body2" className={TypographyStyle}>
              {myContent.introduction}
            </Typography>
            <div className={ButtonDivStyle}>
              <CardActions>
                <CustomLink className={ButtonStyle} to={`/course/editor/${myContent.uuid}`}>
                  <Button variant="contained" size="small">Edit</Button>
                </CustomLink>
                <AlertModal
                  actionButtonBody="Delete"
                  actionButtonColor="error"
                  modalBody="Are you sure to delete your contents?"
                  modalButtonColor="error"
                  onClickActionButton={() => deleteCourse(myContent.uuid)}
                />
              </CardActions>
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}

export default MyContents;
