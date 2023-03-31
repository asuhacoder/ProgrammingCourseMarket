import React, { useState, useEffect } from 'react';
import { SetterOrUpdater, useRecoilState } from 'recoil';
import axios from 'axios';
import { Stack, Snackbar } from '@mui/material';
import MuiAlert, { AlertProps } from '@mui/material/Alert';
import { Lesson, Case } from '../config/Type';
import { lessonState, casesState } from '../config/Recoil';
import data from '../config/language.json';
import CustomTextField from '../atoms/CustomTextField';
import ButtonDiv from '../molecules/ButtonDiv';
import CaseTable from '../organisms/CaseTable';
import { LessonCodeEditorStyle } from './LessonCaseEditor.css';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function LessonCaseEditor(props: TabPanelProps) {
  const languageList = data as any;
  const [lesson]: [Lesson, SetterOrUpdater<Lesson>] = useRecoilState<Lesson>(lessonState);
  const [cases, setCases]: [Case[], SetterOrUpdater<Case[]>] = useRecoilState<Case[]>(casesState);
  const { children, value, index, ...other } = props;
  const [open, setOpen] = useState(false);
  const [input, setInput] = useState('');

  const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(props, ref) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });

  const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen(false);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setInput(e.target.value);
  };

  const submitCase = (language: string, version: string): void => {
    const instance = axios.create({baseURL: process.env.REACT_APP_API_URL})
    instance
      .post(`/api/v1/runner`, {
        code: lesson.answer_code,
        input: input,
        language: languageList[lesson.language.split('@')[0]]['jdoodle'],
        version: lesson.language.split('@')[1],
      })
      .then(
        (response) => {
          instance
            .post(`/api/v1/cases`, {
              token: window.localStorage.getItem('programming-course-market'),
              lesson_id: lesson.uuid,
              input: input,
              output: response.data.output,
            })
            .then((response) => {
              setCases(cases.concat(response.data));
              console.log('cases api response: ', response);
            });
          console.log('runner api response: ', response);
        },
        (error) => {
          console.log(error);
        },
      );
    setInput('');
  };

  return (
    <div
      className={LessonCodeEditorStyle}
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <div>
          <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success" sx={{ width: '100%' }}>
              Saved Changes
            </Alert>
          </Snackbar>
          <Stack id="organisms-lesson-case-stack" justifyContent="flex-start" alignItems="flex-start" spacing={2}>
            <CaseTable />
            <CustomTextField
              required
              id="outlined-basic"
              label="Test Case Input"
              value={input}
              defaultValue={input}
              onChange={handleInputChange}
              multiline
              minRows={5}
              maxRows={20}
            />
            <ButtonDiv body="Submit" onClick={submitCase} />
          </Stack>
        </div>
      )}
    </div>
  );
}
export default LessonCaseEditor;
