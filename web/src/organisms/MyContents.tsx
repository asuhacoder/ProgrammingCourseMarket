import React from 'react';
import {
  Button,
  Card,
  CardActions,
  CardContent,
  Typography,
} from '@mui/material';
import { Course } from '../config/Type';
import {
  MyContentsStyle,
  CardContentStyle,
  TypographyStyle,
  ButtonDivStyle,
} from './MyContents.css';

function MyContents(props: any) {
  const { courses } = props;
  console.log('data.courses', courses);
  console.log(courses === undefined);

  return (
    <div className={MyContentsStyle}>
      {courses === null && (
        <div>No Contents</div>
      )}
      {courses.map((course: Course) => (
        <Card className={CardContentStyle}>
          <CardContent>
            <Typography variant="h4" component="div">
              {course.title}
            </Typography>
            <Typography variant="body2" className={TypographyStyle}>
              {course.introduction}
            </Typography>
            <div className={ButtonDivStyle}>
              <CardActions>
                <Button variant="contained" size="small">Edit</Button>
              </CardActions>
            </div>

          </CardContent>
        </Card>
      ))}

    </div>
  );
}

export default MyContents;
