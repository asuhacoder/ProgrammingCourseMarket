import { atom } from 'recoil';

export interface User {
  token: string,
  uuid: string,
  name: string,
  email: string,
  introduction: string,
}

export const userState = atom({
  key: 'userState',
  default: {
    token: '',
    uuid: '',
    name: '',
    email: '',
    introduction: '',
  },
});
