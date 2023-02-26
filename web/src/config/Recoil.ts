import { atom } from 'recoil';
import { defaultCourseArray, defaultUser, defaultLessons, Lesson } from './Type';

export const userState = atom({
  key: 'userState',
  default: defaultUser,
});

export const myContentsState = atom({
  key: 'myContentsState',
  default: defaultCourseArray,
});

export const lessonsState = atom<Lesson[]>({
  key: 'lessonsState',
  default: defaultLessons,
});