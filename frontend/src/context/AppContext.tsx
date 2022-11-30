import { createContext } from 'react';

export const registerInitialState = {
  name: '',
  email: '',
  password: '',
}

export const loginInitialState = {
  email: '',
  password: '',
}

export const DEFAULT_STATE = {
  register: registerInitialState,
  setRegister: () => {},
  login: loginInitialState,
  setLogin: () => {},
  name: '',
  setName: () => {},
  errorMessage: '',
  setErrorMessage: () => {},
};

const AppContext = createContext<IAppContext>(DEFAULT_STATE);

AppContext.displayName='App Context';

export default AppContext;
