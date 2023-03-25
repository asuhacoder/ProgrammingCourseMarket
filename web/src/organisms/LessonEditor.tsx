import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
import { Tabs, Tab, Box } from '@mui/material';
import LessonTitleEditor from './LessonTitleEditor';
import LessonBodyEditor from './LessonBodyEditor';
import LessonLanguageEditor from './LessonLanguageEditor';
import LessonCodeEditor from './LessonCodeEditor';
import LessonCaseEditor from './LessonCaseEditor';
import { PanelStyle } from './LessonEditor.css';

function a11yProps(index: number) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

function LessonEditor() {
  const [value, setValue] = useState(0);
  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

  return (
    <div>
      <LessonTitleEditor />
      <Box sx={{ width: '100%' }}>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs value={value} onChange={handleChange} aria-label="basic tabs example">
            <Tab label="Body Editor" {...a11yProps(0)} />
            <Tab label="Language Editor" {...a11yProps(1)} />
            <Tab label="Default Code Editor" {...a11yProps(2)} />
            <Tab label="Answer Code Editor" {...a11yProps(3)} />
            <Tab label="Test Cases Editor" {...a11yProps(4)} />
          </Tabs>
        </Box>
        <div className={PanelStyle}>
          <LessonBodyEditor value={value} index={0} />
          <LessonLanguageEditor value={value} index={1} />
          <LessonCodeEditor value={value} index={2} defaultOrAnswer={'default_code'} />
          <LessonCodeEditor value={value} index={3} defaultOrAnswer={'answer_code'} />
          <LessonCaseEditor value={value} index={4} />
        </div>
      </Box>
    </div>
  );
}
export default LessonEditor;
