import React, { useContext, useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import CryptoApi from '../../services/Http';

const CryptosPage: React.FC = () => {
  const history = useHistory();

  const user: ILocalStorage = JSON.parse(localStorage.getItem('user')!);

  const getCryptos = async () => {
    const response: any = await CryptoApi.getCryptos(user.token);
    // if (response.message) {
    //   localStorage.removeItem('user');
    //   history.push('/');
      // return;
    // }
    console.log(response);
  }

  useEffect(() => {
    getCryptos();
  }, []);

  return (
    <main>
      <h1>Cryptos</h1>
    </main>
  );
}

export default CryptosPage;
