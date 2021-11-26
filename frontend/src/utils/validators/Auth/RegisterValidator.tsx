interface ValidatorProps {
  email: string;
  password: string;
  confirmPassword: string;
}

interface errorsProps {
  email?: string;
  password?: string;
  confirmPassword?: string;
}

export const RegisterValidator = (values: ValidatorProps) => {
  const errors: errorsProps = {};

  // EMAIL
  (!values.email && (errors.email = "Required")) ||
    (!/^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/g.test(
      values.email
    ) &&
      (errors.email = "Invalid email address"));
  // PASSWORD
  (!values.password && (errors.password = "Required")) ||
    (/(\s)/g.test(values.password) &&
      (errors.password = "Password must not contain spaces ")) ||
    (values?.password?.length < 8 &&
      (errors.password =
        "Minimum 8 characters, at least 1 letter and 1 number")) ||
    (/[^a-zA-Z0-9]/g.test(values.password) &&
      (errors.password = "Password must not contain cyrillic "));

  // CONFIRM PASSWORD
  (!values.confirmPassword && (errors.confirmPassword = "Required")) ||
    (values.password !== values.confirmPassword &&
      (errors.confirmPassword = "Passwords must match"));

  return errors;
};
