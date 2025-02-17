import { ErrorObject } from "@/app/types/utils";

interface BaseValidation {
  email: string;
  password: string;
}

export interface RegisterValidation extends BaseValidation {
  username: string;
}

const validateEmail = (email: string): string | null => {
  if (!email.trim()) {
    return "Email is required";
  }
  const emailRegex = /\S+@\S+\.\S+/;
  if (!emailRegex.test(email)) {
    return "Email is invalid";
  }
  return null;
};

const validatePassword = (password: string): string | null => {
  if (!password) {
    return "Password is required";
  }
  if (password.length < 8) {
    return "Password must be at least 6 characters";
  }
  return null;
};

const validateUsername = (username: string): string | null => {
  const trimmedUsername = username.trim();
  if (!trimmedUsername) {
    return "Username is required";
  }
  if (trimmedUsername.length < 3) {
    return "Username must be at least 3 characters";
  }
  return null;
};

const accumulateErrors = (validators: { [key: string]: () => string | null }): ErrorObject | null => {
  const errors: ErrorObject = {};

  for (const [field, validate] of Object.entries(validators)) {
    const error = validate();
    if (error) {
      errors[field] = error;
    }
  }

  return Object.keys(errors).length > 0 ? errors : null;
};

export const validateLogin = (data: BaseValidation): ErrorObject | null => {
  return accumulateErrors({
    email: () => validateEmail(data.email),
    password: () => validatePassword(data.password),
  });
};

export const validateRegister = (data: RegisterValidation): ErrorObject | null => {
  return accumulateErrors({
    username: () => validateUsername(data.username),
    email: () => validateEmail(data.email),
    password: () => validatePassword(data.password),
  });
};
