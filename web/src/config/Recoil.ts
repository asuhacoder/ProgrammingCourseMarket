import { atom } from 'recoil';
import { defaultCourseArray, defaultUser } from './Type';

export const userState = atom({
  key: 'userState',
  default: defaultUser,
});

export const myContentsState = atom({
  key: 'myContentsState',
  default: defaultCourseArray,
});
