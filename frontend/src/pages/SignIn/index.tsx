import React, { useContext, useEffect, useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import AppContext from '../../context/AppContext';
import PopUpMessage from '../../components/PopUpMessage';
import { validateLoginCredentials } from '../../helpers/Helpers';
import CryptoApi from '../../services/Http';
import logoImg from '../../images/crypto.png';
import './style.css';

const SignInPage: React.FC = () => {
  const {
    login,
    setLogin,
    errorMessage,
    setErrorMessage,
  } = useContext(AppContext);
  const [disabled, setDisabled] = useState(true);
  const history = useHistory();

  const handleClick = async () => {
    const { token }: { token: string | undefined } =
      await CryptoApi.login(login);
    if (!token) {
      setErrorMessage('Invalid email or password');
      return;
    }
    const localStorageData = { token };
    localStorage.setItem('user', JSON.stringify(localStorageData));
    history.push('/cryptos');
  }

  const handleChange: onChange = ({ target: { name, value } }) => {
    setLogin({ ...login, [name]: value });
  }

  const validateInfos = () => {
    const isValid: boolean = validateLoginCredentials(login);
    setDisabled(isValid);
  }

  useEffect(() => {
    validateInfos();
  }, [login]);

  return (
    <main className="wrapper fadeInDown">
      { errorMessage && <PopUpMessage /> }
      <div id="formContent">
        <h2 className="active"> Sign In </h2>
        <Link to="/signup">
          <h2 className="inactive underlineHover"> Sign Up </h2>
        </Link>
        <div className="fadeIn first">
          <img src={ logoImg } id="icon" alt="Logo Icon" />
        </div>

        <form>
          <input
            type="text"
            id="login"
            className="fadeIn second"
            name="email"
            placeholder="login"
            value={ login['email'] }
            onChange={ handleChange }
            maxLength={ undefined }
          />
          <input
            type="text"
            id="password"
            className="fadeIn third"
            name="password"
            placeholder="password"
            value={ login['password'] }
            onChange={ handleChange }
            maxLength={ undefined }
          />
          <input
            type="button"
            className="fadeIn fourth"
            onClick={ handleClick }
            disabled={ disabled }
            value="Login"
          />
        </form>

        <div id="formFooter">
          <Link className="underlineHover" to="/">Forgot Password?</Link>
        </div>
      </div>
    </main>
  );
}

export default SignInPage;
