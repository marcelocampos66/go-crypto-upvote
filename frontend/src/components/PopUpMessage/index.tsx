import React, { useContext } from 'react';
import AppContext from '../../context/AppContext';
import './style.css';

const PopUpMessage: React.FC = () => {
  const { errorMessage, setErrorMessage } = useContext(AppContext);

  return (
    <div className='popMessage'>
      <p>{ errorMessage }</p>
      <input
        type="button"
        onClick={ () => setErrorMessage('') }
        value="Ok"
      />
    </div>
  );
}

export default PopUpMessage;
