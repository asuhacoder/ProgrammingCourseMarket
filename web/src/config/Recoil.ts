import { atom } from 'recoil';
import { defaultUser } from './Type';

const userState = atom({
  key: 'userState',
  default: defaultUser,
});

export default userState;
