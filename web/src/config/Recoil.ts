import { atom } from 'recoil';
import { defaultCourseArray, defaultUser, defaultLessons, defaultLesson, Lesson } from './Type';

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

export const lessonState = atom<Lesson>({
  key: 'lessonState',
  default: defaultLesson,
})

export const languageState = atom<string>({
  key: 'languageState',
  default: '',
})

export const versionState = atom<string>({
  key: 'versionState',
  default: '',
})