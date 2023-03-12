import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { Stack, Box, Tabs, Tab, Paper } from '@mui/material';
import Editor, { Monaco } from "@monaco-editor/react";
import MDEditor from '@uiw/react-md-editor';
import SplitPane from 'react-split-pane';
import data from '../config/language.json';
import Panel from '../molecules/Panel';
import TestCaseModal from './TestCaseModal';
import { TextStyle, HeightStyle, LessonDetailStackStyle } from './LessonDetail.css';
import TestCaseTable from '../molecules/TestCaseTable';

function a11yProps(index: number) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

function LessonDetail(props: any) {
  const { lesson, cases } = props;
  const languageList = data as any;
  const { id } = useParams();
  const [value, setValue] = useState(0);
  const [code , setCode] = useState(lesson.default_code);
  const [language, setLanguage] = useState('plain');

  const options = {
    readOnly: false,
    minimap: { enabled: false },
  };

  useEffect(() => {
    if (lesson.language !== '') {
      setLanguage(languageList[lesson.language.split('@')[0]].monaco);
    }
  });
  useEffect(() => {
    setCode(lesson.default_code);
  }, [lesson.default_code]);

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };
  function handleEditorChange(value: string | undefined, event: any) {
    setCode(value as string);
  };
  function handleBeforeMount(monaco: Monaco) {
    console.log(monaco);
  }

  return (
    <Stack
      className={LessonDetailStackStyle}
      id="organisms-lesson-detail-stack"
      justifyContent="flex-start"
      alignItems="flex-start"
      spacing={2}
    >
      <h1 className={TextStyle}>{lesson.title}</h1>
      <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
        <Tabs value={value} onChange={handleChange} aria-label="basic tabs example">
          <Tab label="Body & Editor" {...a11yProps(0)} />
          <Tab label="Test Cases" {...a11yProps(1)} />
        </Tabs>
      </Box>
      <Panel value={value} index={0} >
        <div className={HeightStyle}>
          <SplitPane split="vertical" minSize={50} defaultSize="50%">
            <Paper className={HeightStyle}>
              <MDEditor.Markdown source={lesson.body} style={{ whiteSpace: 'pre-wrap', height: "650px", overflow:"scroll" }} />
            </Paper>
            <Paper className={HeightStyle}>
              <Editor
                height="65vh"
                language={language}
                defaultValue={lesson.default_code}
                value={code}
                onChange={handleEditorChange}
                beforeMount={handleBeforeMount}
                options={options}
              />
              <TestCaseModal lesson={lesson} cases={cases} code={code} />
            </Paper>
          </SplitPane>
        </div>
      </Panel>
      <Panel value={value} index={1} >
        <TestCaseTable cases={cases} />
      </Panel>
    </Stack>
  );
}
export default LessonDetail;