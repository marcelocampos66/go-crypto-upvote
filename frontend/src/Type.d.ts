interface IAppContext {
  register: IRegister;
  setRegister: React.Dispatch<React.SetStateAction<IRegister>>;
  login: ILogin;
  setLogin: React.Dispatch<React.SetStateAction<ILogin>>;
  name: string;
  setName: React.Dispatch<React.SetStateAction<string>>;
  errorMessage: string;
  setErrorMessage: React.Dispatch<React.SetStateAction<string>>;
}

type IRegister = {
  name: string;
  email: string;
  password: string;
}

type ILogin = {
  email: string;
  password: string;
}

type ILocalStorage = {
  token: string;
}

type onChange = (e: React.ChangeEvent<HTMLInputElement>) => void

type onChangeDropDown = (e: React.ChangeEvent<HTMLSelectElement>) => void
