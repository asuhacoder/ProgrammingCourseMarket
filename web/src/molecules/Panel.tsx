import React from 'react';
import { PanelDivStyle } from './Panel.css';

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function Panel(props: TabPanelProps) {
  const { value, index, children } = props;

  return (
    <div
      className={PanelDivStyle}
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
    >
      {value === index && (
        <div >
          {children}
        </div>
      )}
    </div>
  );
}

export default Panel;
