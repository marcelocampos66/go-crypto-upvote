import React from 'react';
import { Switch, Route } from 'react-router-dom';
import SignInPage from './pages/SignIn';
import SignUpPage from './pages/SignUp';
import CryptosPage from './pages/Cryptos';

function App() {
  return (
    <Switch>
      <Route path="/cryptos" component={ CryptosPage } />
      <Route path="/signup" component={ SignUpPage } />
      <Route exact path="/" component={ SignInPage } />
    </Switch>
  );
}

export default App;
