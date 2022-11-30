const MIN_PASSWORD_LENGTH: number = 6;
const MIN_NAME_LENGTH: number = 4;
const EMAIL_REGEX: RegExp = /^([\w.%+-]+)@([\w-]+\.)+([\w]{2,})$/i;

export const validateNewUserData = (data: IRegister): boolean => {
  const { name, email, password } = data;
  if (name.length < MIN_NAME_LENGTH
    || !EMAIL_REGEX.test(email)
    || password.length < MIN_PASSWORD_LENGTH) {
    return true;
  }
  return false;
};

export const validateLoginCredentials = (data: ILogin): boolean => {
  const { email, password } = data;
  if (!EMAIL_REGEX.test(email)
    || password.length < MIN_PASSWORD_LENGTH) {
    return true;
  }
  return false;
};
