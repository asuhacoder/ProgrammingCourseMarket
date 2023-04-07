import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import { SortableContainer, SortableElement, SortableHandle } from 'react-sortable-hoc';
import { arrayMoveImmutable } from 'array-move';
import { Button, Card, CardContent, CardActions, Stack, Snackbar, Typography } from '@mui/material';
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import { useRecoilState, SetterOrUpdater } from 'recoil';
import { userState, lessonsState } from '../config/Recoil';
import { User, Lesson } from '../config/Type';
import CustomTextField from '../atoms/CustomTextField';
import CustomLink from '../atoms/CustomLink';
import ButtonDiv from '../molecules/ButtonDiv';
import AlertModal from '../molecules/AlertModal';
import { NumberingLessons, UpdateLessons } from '../utils/LessonsUtils';
import {
  LessonListEditorStackStyle,
  NewLessonTitleTextFieldStyle,
  SortableListStyle,
  ButtonDivStyle,
  ButtonStyle,
} from './LessonListEditor.css';

function LessonListEditor() {
  const { id } = useParams();
  const user: User = useRecoilState(userState)[0];
  const [lessons, setLessons]: [Lesson[], SetterOrUpdater<Lesson[]>] = useRecoilState<Lesson[]>(lessonsState);
  const [title, setTitle] = useState('');
  const [titleHasError, setTitleHasError] = useState(false);
  const [titleHelperText, setTitleHelperText] = useState('');
  const [open, setOpen] = useState(false);
  const instance = axios.create({baseURL: process.env.REACT_APP_API_URL})

  const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(props, ref) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });

  const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen(false);
  };

  const onSortEnd = ({ oldIndex, newIndex }: { oldIndex: number; newIndex: number }) => {
    console.log('oldIndex: ', oldIndex);
    console.log('newIndex: ', newIndex);
    setLessons(arrayMoveImmutable(lessons, oldIndex, newIndex));
    setLessons(NumberingLessons(lessons));
    UpdateLessons(id as string, user, lessons[oldIndex]);
    UpdateLessons(id as string, user, lessons[newIndex]);
    setOpen(true);
  };

  const validateTitle = (): boolean => {
    let isValid = true;
    if (title.length < 3) {
      isValid = false;
      setTitleHasError(true);
      setTitleHelperText('length of title must be more than or equal 3');
    } else if (title.length > 50) {
      setTitleHasError(true);
      setTitleHelperText('length of title must be less than or equal 50');
    } else {
      setTitleHasError(false);
      setTitleHelperText('');
    }
    return isValid;
  };

  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setTitle(e.target.value);
  };

  const submitNewLesson = (): void => {
    if (validateTitle()) {
      instance
        .post(`/api/v1/lessons`, {
          token: window.localStorage.getItem('programming-course-market'),
          user_id: user.uuid,
          course_id: id,
          sequence_number: lessons.length,
          title,
          introduction: '',
          body: '**Hello world!!!**',
          default_code: '',
          answer_code: '',
          language: 'Python 3@4',
        })
        .then(
          (response) => {
            console.log(response);
            setLessons(lessons.concat(response.data));
          },
          (error) => {
            console.log(error);
          },
        );
      setTitle('');
    }
  };

  const deleteLesson = (uuid: string): void => {
    instance
      .delete(`/api/v1/lessons/${uuid}`, {
        data: {
          user_id: user.uuid,
          token: window.localStorage.getItem('programming-course-market'),
        },
      })
      .then(
        (response) => {
          console.log(response);
          setLessons(lessons.filter((lesson: Lesson) => lesson.uuid !== uuid));
          setLessons(NumberingLessons(lessons));
          for (let i = 0; i < lessons.length; i++) {
            UpdateLessons(id as string, user, lessons[i]);
          }
        },
        (error) => {
          console.log(error);
        },
      );
  };

  const DragHandle = SortableHandle(() => <span>::</span>);

  const SortableItem = SortableElement(({ lesson }: { lesson: Lesson }) => (
    <Card className={SortableListStyle}>
      <DragHandle />
      <CardContent>
        <Typography variant="h6" component="div">
          {lesson.title}
        </Typography>
        <div className={ButtonDivStyle}>
          <CardActions>
            <CustomLink className={ButtonStyle} to={`/lesson/editor/${lesson.uuid}`}>
              <Button variant="contained" size="small">
                Edit
              </Button>
            </CustomLink>
            <AlertModal
              actionButtonBody="Delete"
              actionButtonColor="error"
              modalBody="Are you sure to delete your lesson?"
              modalButtonColor="error"
              onClickActionButton={() => deleteLesson(lesson.uuid)}
            />
          </CardActions>
        </div>
      </CardContent>
    </Card>
  ));

  const SortableList = SortableContainer(({ children }: { children: any }) => {
    return <div className={SortableListStyle}>{children}</div>;
  });

  return (
    <Stack className={LessonListEditorStackStyle} justifyContent="flex-start" alignItems="flex-start" spacing={2}>
      <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success" sx={{ width: '100%' }}>
          Saved Changes
        </Alert>
      </Snackbar>
      <SortableList onSortEnd={onSortEnd} useDragHandle>
        {lessons.map((value: Lesson, index: number) => (
          <SortableItem key={`item-${value.uuid}`} index={index} lesson={value} />
        ))}
      </SortableList>
      <CustomTextField
        className={NewLessonTitleTextFieldStyle}
        required
        id="outlined-basic"
        label="Lesson Title"
        value={title}
        defaultValue={title}
        helperText={titleHelperText}
        error={titleHasError}
        onChange={handleTitleChange}
        onBlur={validateTitle}
      />
      <ButtonDiv body="Add" variant="contained" size="small" onClick={submitNewLesson}></ButtonDiv>
    </Stack>
  );
}
export default LessonListEditor;
