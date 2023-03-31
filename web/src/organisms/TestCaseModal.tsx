import React, { useState } from 'react';
import axios from 'axios';
import {
  Box,
  Button,
  TableContainer,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Paper,
  Modal,
  CircularProgress,
  Stack,
} from '@mui/material';
import CheckIcon from '@mui/icons-material/Check';
import CloseIcon from '@mui/icons-material/Close';
import { Case } from '../config/Type';
import data from '../config/language.json';
import CustomLink from '../atoms/CustomLink';
import ButtonDiv from '../molecules/ButtonDiv';
import { ButtonDivStyle, ButtonStyle } from './TestCaseModal.css';

const style = {
  position: 'absolute' as 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 800,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

function TestCaseModal(props: any) {
  const languageList = data as any;
  const { lesson, cases, code } = props;
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const [outputs, setOutputs] = useState<any>({});
  const [complete, setComplete] = useState(false);

  const handleClick = () => {
    handleOpen();
    runCode();
  };

  const runCode = () => {
    let tmp: any = {};
    let done = true;
    cases.forEach((testCase: Case) => {
      const instance = axios.create({baseURL: process.env.REACT_APP_API_URL})
      instance
        .post(`/api/v1/runner`, {
          code: code,
          input: testCase.input,
          language: languageList[lesson.language.split('@')[0]]['jdoodle'],
          version: lesson.language.split('@')[1],
        })
        .then(
          (response) => {
            tmp[testCase.uuid] = response.data.output;
            setOutputs(tmp);
            if (testCase.output !== response.data.output) {
              done = false;
            }
            setComplete(done);
            console.log('runner api response: ', response);
          },
          (error) => {
            console.log(error);
          },
        );
    });
  };

  return (
    <div>
      <div className={ButtonDivStyle}>
        <Button className={ButtonStyle} color="primary" variant="contained" onClick={handleClick}>
          Run Code
        </Button>
      </div>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <Stack justifyContent="flex-start" alignItems="flex-start" spacing={2}>
            <TableContainer component={Paper}>
              <Table sx={{ minWidth: 400 }} aria-label="simple table">
                <TableHead>
                  <TableRow>
                    <TableCell>Test Case Number</TableCell>
                    <TableCell align="left">Input</TableCell>
                    <TableCell align="left">Answer Output</TableCell>
                    <TableCell align="left">Your Output</TableCell>
                    <TableCell align="left">Judge</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {cases === null && (
                    <TableRow key="empty" sx={{ '&:last-child td, &:last-child th': { border: 0 } }}>
                      <TableCell component="th" scope="row">
                        No Test Case
                      </TableCell>
                      <TableCell align="left">No Data</TableCell>
                      <TableCell align="left">No Data</TableCell>
                      <TableCell align="left">Skipped</TableCell>
                      <TableCell align="left">Skipped</TableCell>
                    </TableRow>
                  )}
                  {cases !== null &&
                    cases.map((testCase: Case, index: number) => (
                      <TableRow key={testCase.uuid} sx={{ '&:last-child td, &:last-child th': { border: 0 } }}>
                        <TableCell component="th" scope="row">
                          Test Case {index + 1}
                        </TableCell>
                        <TableCell align="left">
                          {testCase.input.split('\n').map((line, idx) => (
                            <p key={idx}>{line}</p>
                          ))}
                        </TableCell>
                        <TableCell align="left">
                          {testCase.output.split('\n').map((line, idx) => (
                            <p key={idx}>{line}</p>
                          ))}
                        </TableCell>
                        <TableCell align="left">
                          {testCase.uuid in outputs &&
                            outputs[testCase.uuid]
                              .split('\n')
                              .map((line: string, idx: number) => <p key={idx}>{line}</p>)}
                        </TableCell>
                        <TableCell align="left">
                          {!(testCase.uuid in outputs) && <CircularProgress />}
                          {testCase.uuid in outputs && testCase.output === outputs[testCase.uuid] && <CheckIcon />}
                          {testCase.uuid in outputs && testCase.output !== outputs[testCase.uuid] && (
                            <CloseIcon color="error" />
                          )}
                        </TableCell>
                      </TableRow>
                    ))}
                </TableBody>
              </Table>
            </TableContainer>
            {complete && (
              <CustomLink className={ButtonDivStyle} to={`/course/detail/${lesson.course_id}`}>
                <ButtonDiv body="Back to Lesson List" />
              </CustomLink>
            )}
          </Stack>
        </Box>
      </Modal>
    </div>
  );
}

export default TestCaseModal;
