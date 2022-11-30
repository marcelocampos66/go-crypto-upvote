class UsersApi {
  private url: string;
  private content: string;

  constructor() {
    this.url = process.env.REACT_APP_CRYPTO_VOTE_API ||'http://0.0.0.0:3002';
    this.content = 'application/json';
  }

  public async register(data: IRegister) {
    const endpoint = `${this.url}/v1/crypto-upvote/users`;
    return fetch(endpoint, {
      method: 'POST',
      headers: {
        'Content-Type': this.content,
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((data) => data)
      .catch((err) => err);
  }

  public async login(data: ILogin) {
    const endpoint = `${this.url}/v1/crypto-upvote/login`;
    return fetch(endpoint, {
      method: 'POST',
      headers: {
        'Content-Type': this.content,
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((data) => data)
      .catch((err) => err);
  }

  public async getCryptos(token: string) {
    const endpoint = `${this.url}/v1/crypto-upvote/cryptos`;
    return fetch(endpoint, {
      method: 'GET',
      headers: {
        'Content-Type': this.content,
        'Authorization': `Bearer ${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => data)
      .catch((err) => err);
  }

  public async cryptoVoteUp(id: number, token: string) {
    const endpoint = `${this.url}/v1/crypto-upvote/cryptos/${id}/up`
    return fetch(endpoint, {
      method: 'PUT',
      headers: {
        'Content-Type': this.content,
        'Authorization': `Bearer ${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => data)
      .catch((err) => err);
  }

  public async cryptoVoteDown(id: number, token: string) {
    const endpoint = `${this.url}/v1/crypto-upvote/cryptos/${id}/down`
    return fetch(endpoint, {
      method: 'PUT',
      headers: {
        'Content-Type': this.content,
        'Authorization': `Bearer ${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => data)
      .catch((err) => err);
  }
}

export default new UsersApi();
