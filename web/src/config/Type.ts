export interface User {
  token: string;
  uuid: string;
  name: string;
  email: string;
  introduction: string;
}

export const defaultUser = {
  token: '',
  uuid: '',
  name: '',
  email: '',
  introduction: '',
};

export interface Course {
  uuid: string;
  user_id: string;
  title: string;
  introduction: string;
  is_public: boolean;
}

export const defaultCourse = {
  uuid: '',
  user_id: '',
  title: '',
  introduction: '',
  is_public: false,
};

export const defaultCourseArray = [];

export interface Lesson {
  uuid: string;
  user_id: string;
  course_id: string;
  sequence_number: number;
  title: string;
  introduction: string;
  body: string;
  default_code: string;
  answer_code: string;
  language: string;
}

export const defaultLessons = [];

export const defaultLesson = {
  uuid: '',
  user_id: '',
  course_id: '',
  sequence_number: 0,
  title: '',
  introduction: '',
  body: '',
  default_code: '',
  answer_code: '',
  language: '',
};

export interface Case {
  uuid: string;
  user_id: string;
  lesson_id: string;
  input: string;
  output: string;
}

export const defaultCases = [];
