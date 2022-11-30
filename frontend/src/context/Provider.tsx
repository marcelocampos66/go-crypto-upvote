import React, { useState } from 'react';
import AppContext from './AppContext';
import { DEFAULT_STATE } from './AppContext';

const Provider: React.FC = ({ children }) => {
  const [register, setRegister] =
    useState<IRegister>(DEFAULT_STATE.register);
  const [login, setLogin] = useState<ILogin>(DEFAULT_STATE.login);
  const [name, setName] = useState<string>(DEFAULT_STATE.name);
  const [errorMessage, setErrorMessage] =
    useState<string>(DEFAULT_STATE.errorMessage);
  const [cryptos, setCryptos] = useState<[] | TCrypto[]>([]);

  const contextValue = {
    register,
    setRegister,
    login,
    setLogin,
    name,
    setName,
    errorMessage,
    setErrorMessage,
    cryptos,
    setCryptos,
  };

  return (
    <AppContext.Provider value={ contextValue }>
      { children }
    </AppContext.Provider>
  )
}

export default Provider;
