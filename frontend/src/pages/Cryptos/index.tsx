import React, { useEffect, useContext } from 'react';
import { useHistory } from 'react-router-dom';
import AppContext from '../../context/AppContext';
import CryptoApi from '../../services/Http';
import CryptoCard from '../../components/CryptoCard';
import LoggoutButton from '../../components/LoggoutButton';
import './style.css'

const CryptosPage: React.FC = () => {
  const { cryptos, setCryptos } = useContext(AppContext);
  const history = useHistory();

  const user: ILocalStorage = JSON.parse(localStorage.getItem('user')!);

  const getCryptos = async () => {
    const response = await CryptoApi.getCryptos(user.token);
    if (response.message) {
      localStorage.removeItem('user');
      history.push('/');
      return;
    }

    setCryptos(response);
  }

  useEffect(() => {
    getCryptos();
  }, []);

  return (
    <main>
      <section className='cardsContainer'>
        <div className='cardsWrapper'>
          <LoggoutButton />
          {
            cryptos.map((crypto) => (
              <CryptoCard key={crypto.id} crypto={ crypto } />
            ))
          }
        </div>
      </section>
    </main>
  );
}

export default CryptosPage;
