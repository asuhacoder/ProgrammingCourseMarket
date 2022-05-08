import { atom } from 'recoil';

export const defaultUser = {
  token: '',
  uuid: '',
  name: '',
  email: '',
  introduction: '',
};

export const userState = atom({
  key: 'userState',
  default: defaultUser,
});
