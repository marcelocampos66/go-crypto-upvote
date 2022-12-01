import React, { useContext, useState } from "react";
import AppContext from '../../context/AppContext';
import CryptoApi from '../../services/Http';
import './style.css';
import logoLike from '../../images/like.png';
import logoDeslike from '../../images/deslike.png';

type Props = { crypto: TCrypto };

const CryptoCard: React.FC<Props> = ({ crypto }) => {
  const { cryptos, setCryptos } = useContext(AppContext);
  const [update, setUpdate] = useState(false);

  const user: ILocalStorage = JSON.parse(localStorage.getItem('user')!);

  const voteUp = async (cryptoId: number) => {
    const newState = cryptos;
    const index = newState.findIndex((crypto) => crypto.id === cryptoId);
    await CryptoApi.cryptoVoteUp(cryptoId, user.token);
    let updatedVotes = crypto.votes;
    updatedVotes += 1;
    newState[index].votes = updatedVotes;
    setCryptos(newState);
    setUpdate(!update);
  };

  const voteDown = async (cryptoId: number) => {
    const newState = cryptos;
    const index = newState.findIndex((crypto) => crypto.id === cryptoId);
    await CryptoApi.cryptoVoteDown(cryptoId, user.token);
    let updatedVotes = crypto.votes;
    if (updatedVotes >= 1) {
      updatedVotes -= 1;
      newState[index].votes = updatedVotes;
      setCryptos(newState);
      setUpdate(!update);
    }
  };

  return (
    <div className="card">
      <div className="card-header">
        <img src={ crypto.imageUrl } alt="rover" />
      </div>
      <div className="card-body">
        <span className="tag tag-teal">{ `Ranking #${crypto.id}` }</span>
        <h4>{ crypto.cryptoName }</h4>
        <p>{ `Votes quantity: ${crypto.votes}` }</p>
        <p>{ `Current quotation: ${parseFloat(crypto.quotation).toLocaleString('pt-br',{style: 'currency', currency: 'BRL'})}` }</p>
        <div className="vote">
          <img src={ crypto.imageUrl } alt="vote" />
          <div className="vote-info">
            <button
              type="button"
              className="vote-button"
              onClick={ () => voteDown(crypto.id) }
            >
              <img src={ logoDeslike } alt="logo like" style={{ width: '15px', height: '15px' }} />
            </button>
            <button
              type="button"
              className="vote-button"
              onClick={ () => voteUp(crypto.id) }
            >
              <img src={ logoLike } alt="logo deslike" style={{ width: '15px', height: '15px' }} />
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default CryptoCard;
