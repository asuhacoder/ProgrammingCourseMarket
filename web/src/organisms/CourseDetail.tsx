import React from 'react';
import {
  Card,
  CardContent,
  Typography,
} from '@mui/material';
import {
  CardContentStyle,
  TypographyStyle,
} from './CourseDetail.css';

function CourseDetail(props: any) {
  const { course } = props;
  console.log(course);
  return (
    <Card className={CardContentStyle}>
      <CardContent>
        <Typography variant="h4" component="div">
          {course.title}
        </Typography>
        <Typography variant="body2" className={TypographyStyle}>
          {course.introduction}
        </Typography>
      </CardContent>
    </Card>
  );
}

export default CourseDetail;
