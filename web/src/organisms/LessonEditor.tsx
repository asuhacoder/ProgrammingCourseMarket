import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
import { Tabs, Tab, Typography, Box } from '@mui/material';
import LessonTitleEditor from './LessonTitleEditor'
import LessonBodyEditor from './LessonBodyEditor'
import { PanelStyle } from './LessonEditor.css';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          <Typography>{children}</Typography>
        </Box>
      )}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

function LessonEditor() {
  const { id } = useParams();
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
            <Tab label="Default Code Editor" {...a11yProps(1)} />
            <Tab label="Answer Code Editor" {...a11yProps(2)} />
            <Tab label="Test Cases Editor" {...a11yProps(3)} />
          </Tabs>
        </Box>
        <div className={PanelStyle}>
          <LessonBodyEditor value={value} index={0} />
          <TabPanel value={value} index={1}>
            Item Two
          </TabPanel>
          <TabPanel value={value} index={2}>
            Item Three
          </TabPanel>
          <TabPanel value={value} index={3}>
            Item Four
          </TabPanel>
        </div>
      </Box>
    </div>

  );
}
export default LessonEditor;