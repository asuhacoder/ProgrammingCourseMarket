import React from 'react';
import axios from 'axios';
import { SetterOrUpdater, useRecoilState } from 'recoil';
import { TableContainer, Table, TableHead, TableRow, TableCell, TableBody, Paper } from '@mui/material';
import { User, Case } from '../config/Type';
import { userState, casesState } from '../config/Recoil';
import AlertModal from '../molecules/AlertModal';

function CaseTable() {
  const [cases, setCases]: [Case[], SetterOrUpdater<Case[]>] = useRecoilState<Case[]>(casesState);
  const user: User = useRecoilState(userState)[0];
  const deleteCase = (uuid: string): void => {
    axios
      .delete(`http://localhost:8080/api/v1/cases/${uuid}`, {
        data: {
          user_id: user.uuid,
          token: window.localStorage.getItem('programming-course-market'),
        },
      })
      .then(
        (response) => {
          console.log(response);
          setCases(cases.filter((testCase: Case) => testCase.uuid !== uuid));
        },
        (error) => {
          console.log(error);
        },
      );
  };

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Test Case Number</TableCell>
            <TableCell align="left">Input</TableCell>
            <TableCell align="left">Output generated by Answer Code</TableCell>
            <TableCell align="left"></TableCell>
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
              <TableCell align="left" />
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
                <TableCell align="right">
                  <AlertModal
                    actionButtonBody="Delete"
                    actionButtonColor="error"
                    modalBody="Are you sure to delete your Test Case?"
                    modalButtonColor="error"
                    onClickActionButton={() => deleteCase(testCase.uuid)}
                  />
                </TableCell>
              </TableRow>
            ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}

export default CaseTable;
