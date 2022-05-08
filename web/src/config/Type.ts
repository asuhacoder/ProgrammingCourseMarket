export interface User {
  token: string,
  uuid: string,
  name: string,
  email: string,
  introduction: string,
}

export interface Course {
  uuid: string,
  user_id: string,
  title: string,
  introduction: string,
  is_public: boolean,
}
