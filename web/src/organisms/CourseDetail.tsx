import React from 'react';
import { Link } from 'react-router-dom';
import {
  Card,
  CardContent,
  Stack,
  Typography,
} from '@mui/material';
import { Lesson } from '../config/Type';
import {
  CardContentStyle,
  ContentsStyle,
  CourseDetailStackStyle,
  TextStyle,
  TypographyStyle,
  TitleLink,
  NoContentsStyle,
} from './CourseDetail.css';

function CourseDetail(props: any) {
  const { course, lessons } = props;
  console.log(course);
  console.log('lessons: ', lessons)
  return (
    <Stack
      className={CourseDetailStackStyle}
      justifyContent="flex-start"
      alignItems="flex-start"
      spacing={2}
    >
      <h1 className={TextStyle}>Course Description</h1>
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
      <h1 className={TextStyle}>Lesson List</h1>
      <div className={ContentsStyle}>
        {(lessons === null || JSON.stringify(lessons) === '[]') && (
          <div className={NoContentsStyle}>No Contents</div>
        )}
        {!(lessons === null || JSON.stringify(lessons) === '[]') && lessons.map((lesson: Lesson) => (
          <Card key={lesson.uuid} className={CardContentStyle}>
            <CardContent>
              <Typography variant="h6" component="div">
                <Link className={TitleLink} to={`/lesson/${lesson.uuid}`}>
                  {lesson.title}
                </Link>
              </Typography>
            </CardContent>
          </Card>
        ))}
      </div>
    </Stack>
  );
}

export default CourseDetail;
