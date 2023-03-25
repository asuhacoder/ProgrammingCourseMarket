import React, { useState, useEffect } from 'react';
import { SetterOrUpdater, useRecoilState } from 'recoil';
import axios from 'axios';
import { Stack, Snackbar, MenuItem, InputLabel, FormControl } from '@mui/material';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import { User, Lesson } from '../config/Type';
import { userState, lessonState } from '../config/Recoil';
import data from '../config/language.json';
import { LessonLanguageEditorStyle, SelectFormStyle } from './LessonLanguageEditor.css';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function LessonLanguageEditor(props: TabPanelProps) {
  const languageList = data as any;
  const user: User = useRecoilState(userState)[0];
  const [lesson, setLesson]: [Lesson, SetterOrUpdater<Lesson>] = useRecoilState<Lesson>(lessonState);
  const { children, value, index, ...other } = props;
  const [language, setLanguage] = useState<string>('');
  const [version, setVersion] = useState<string>('');
  const [open, setOpen] = useState(false);
  useEffect(() => {
    setLanguage(lesson.language.split('@')[0]);
    setVersion(lesson.language.split('@')[1]);
  }, [lesson]);

  const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(props, ref) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });

  const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen(false);
  };

  const handleLanguageChange = (event: SelectChangeEvent) => {
    setLanguage(event.target.value as string);
    setVersion('0');
    submitLanguage(event.target.value as string, '0');
  };
  const handleVersionChange = (event: SelectChangeEvent) => {
    setVersion(event.target.value as string);
    submitLanguage(language, event.target.value as string);
  };

  const submitLanguage = (language: string, version: string): void => {
    const newLesson: Lesson = {
      uuid: lesson.uuid,
      user_id: user.uuid,
      course_id: lesson.course_id,
      sequence_number: lesson.sequence_number,
      title: lesson.title,
      introduction: lesson.introduction,
      body: lesson.body,
      default_code: lesson.default_code,
      answer_code: lesson.answer_code,
      language: language + '@' + version,
    };
    setLesson(newLesson);
    axios
      .put(`http://localhost:8080/api/v1/lessons/${lesson.uuid}`, {
        token: window.localStorage.getItem('programming-course-market'),
        ...newLesson,
      })
      .then(
        (response) => {
          console.log(response);
          setOpen(true);
        },
        (error) => {
          console.log(error);
        },
      );
  };

  return (
    <div
      className={LessonLanguageEditorStyle}
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <div className={LessonLanguageEditorStyle}>
          <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success" sx={{ width: '100%' }}>
              Saved Changes
            </Alert>
          </Snackbar>
          <Stack
            id="organisms-lesson-language-stack"
            className={SelectFormStyle}
            justifyContent="flex-start"
            alignItems="flex-start"
            spacing={2}
          >
            <FormControl>
              <InputLabel variant="standard" htmlFor="uncontrolled-native">
                Language
              </InputLabel>
              <Select
                labelId="version select"
                value={language}
                defaultValue={language}
                label="Language"
                onChange={handleLanguageChange}
              >
                {Object.keys(languageList).map((language: string) => (
                  <MenuItem key={language} value={language}>
                    {language}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
            <FormControl>
              <InputLabel variant="standard" htmlFor="uncontrolled-native">
                Language
              </InputLabel>
              <Select
                labelId="version select"
                value={version}
                defaultValue={version}
                label="Version"
                onChange={handleVersionChange}
              >
                {languageList[language]['versions'].map((version: string, index: number) => (
                  <MenuItem key={version} value={index}>
                    {version}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Stack>
        </div>
      )}
    </div>
  );
}
export default LessonLanguageEditor;
