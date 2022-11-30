import React, { useContext, useEffect, useState } from 'react';
import { useHistory, Link } from 'react-router-dom';
import AppContext from '../../context/AppContext';
import PopUpMessage from '../../components/PopUpMessage';
import { validateNewUserData } from '../../helpers/Helpers';
import CryptoApi from '../../services/Http';
import logoImg from '../../images/crypto.png';
import './style.css';

const SignUpPage: React.FC = () => {
  const {
    register,
    setRegister,
    errorMessage,
    setErrorMessage,
  } = useContext(AppContext);
  const [disabled, setDisabled] = useState(true);
  const history = useHistory();

  const handleClick = async () => {
    const newUser = {
      ...register,
    };
    const { message } = await CryptoApi.register(newUser);
    if (message) {
      setErrorMessage(message);
      return;
    }
    history.push('/');
  }

  const handleChange: onChange = ({ target: { name, value } }) => {
    setRegister({ ...register, [name]: value });
  }

  const validateInfos = () => {
    const isValid: boolean = validateNewUserData(register);
    setDisabled(isValid);
  }

  useEffect(() => {
    validateInfos();
  }, [register]);

  return (
    <main className="wrapper fadeInDown">
      { errorMessage && <PopUpMessage /> }
      <div id="formContent">
        <Link to="/">
          <h2 className="inactive underlineHover"> Sign In </h2>
        </Link>
        <h2 className="active"> Sign Up </h2>
        
        <div className="fadeIn first">
          <img src={ logoImg } id="icon" alt="Logo Icon" />
        </div>

        <form>
          <input
            type="text"
            id="name"
            className="fadeIn second"
            name="name"
            placeholder="name"
            value={ register['name'] }
            onChange={ handleChange }
            maxLength={ undefined }
          />
          <input
            type="text"
            id="email"
            className="fadeIn second"
            name="email"
            placeholder="login"
            value={ register['email'] }
            onChange={ handleChange }
            maxLength={ undefined }
          />
          <input
            type="text"
            id="password"
            className="fadeIn third"
            name="password"
            placeholder="password"
            value={ register['password'] }
            onChange={ handleChange }
            maxLength={ undefined }
          />
          <input
            type="button"
            className="fadeIn fourth"
            onClick={ handleClick }
            disabled={ disabled }
            value="Register"
          />
        </form>

        <div id="formFooter">
          <Link className="underlineHover" to="/signup">Registration terms. Click here to learn more!</Link>
        </div>
      </div>
    </main>
  );
}

export default SignUpPage;
