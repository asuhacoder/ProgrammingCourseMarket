export interface User {
  token: string,
  uuid: string,
  name: string,
  email: string,
  introduction: string,
}

export const defaultUser = {
  token: '',
  uuid: '',
  name: '',
  email: '',
  introduction: '',
};

export interface Course {
  uuid: string,
  user_id: string,
  title: string,
  introduction: string,
  is_public: boolean,
}

export const defaultCourse = {
  uuid: '',
  user_id: '',
  title: '',
  introduction: '',
  is_public: false,
};
