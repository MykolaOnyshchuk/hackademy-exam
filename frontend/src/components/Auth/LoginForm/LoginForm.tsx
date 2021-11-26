import { FC } from "react";
import { Form } from "react-final-form";

import FormField from "@components/FormField/FormField";
import { AuthButtonContainer } from "@components/Auth/AuthButtons/AuthButtonsContainer/AuthButtonsContainer";

import {
  LoginButtons,
  LoginFields,
} from "@utils/constants/AuthField/LoginFields/LoginFields";
import { LoginValidator } from "@utils/validators/Auth/LoginValidator";

import {
  ButtonContainer,
  FieldCustom,
  FormContainer,
} from "@modules/Auth/styled/styled";
import { NavLink } from "react-router-dom";
import { RegisterLine, RegisterLink, TextBlock } from "./stylel";

interface LoginProps {
  onSubmit: any;
}

export const LoginForm: FC<LoginProps> = ({ onSubmit }) => {
  return (
    <Form
      onSubmit={onSubmit}
      validate={LoginValidator}
      render={({ handleSubmit }) => (
        <FormContainer>
          <form onSubmit={handleSubmit}>
            {LoginFields.map(({ type, name, label }) => (
              <FieldCustom
                type={type}
                name={name}
                variant="standard"
                placeholder={label}
                component={FormField}
              ></FieldCustom>
            ))}
            <TextBlock>
              <RegisterLine>
                No account? <NavLink to="/auth/register"><RegisterLink>Create one</RegisterLink></NavLink>
              </RegisterLine>
              <NavLink to="#">Forgot password?</NavLink>
            </TextBlock>
            <ButtonContainer>
              {LoginButtons.map(({ button }) => (
                <AuthButtonContainer button={button} />
              ))}
            </ButtonContainer>
          </form>
        </FormContainer>
      )}
    />
  );
};
