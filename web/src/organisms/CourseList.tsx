import React from 'react';
import { Link } from 'react-router-dom';
import {
  Card,
  CardContent,
  Typography,
} from '@mui/material';
import { Course } from '../config/Type';
import {
  NoContentsStyle,
  ContentsStyle,
  CardContentStyle,
  TypographyStyle,
  TitleLink,
} from './CourseList.css';

function CourseList(props: any) {
  const { courses } = props;
  return (
    <div className={ContentsStyle}>
      {(courses === null || JSON.stringify(courses) === '{}') && (
        <div className={NoContentsStyle}>No Contents</div>
      )}
      {!(courses === null || JSON.stringify(courses) === '{}') && courses.map((myContent: Course) => (
        <Card key={myContent.uuid} className={CardContentStyle}>
          <CardContent>
            <Typography variant="h4" component="div">
              <Link className={TitleLink} to={`/course/detail/${myContent.uuid}`}>
                {myContent.title}
              </Link>
            </Typography>
            <Typography variant="body2" className={TypographyStyle}>
              {myContent.introduction}
            </Typography>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}

export default CourseList;
