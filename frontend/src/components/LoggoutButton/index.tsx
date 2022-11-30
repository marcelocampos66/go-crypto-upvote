import React from "react";
import { useHistory } from "react-router-dom";
import './style.css';

const LoggoutButton: React.FC = () => {
  const history = useHistory();

  const handleClick = () => {
    localStorage.removeItem('user');
    history.push('/');
  }

  return (
    <div className="loggout">
      <button
        className="loggout-button"
        onClick={ () => handleClick() }
      >
        LOGGOUT
      </button>
    </div>
  );
}

export default LoggoutButton;
