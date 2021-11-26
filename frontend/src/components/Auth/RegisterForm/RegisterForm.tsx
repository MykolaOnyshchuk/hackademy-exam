import { FC } from "react";
import { Form } from "react-final-form";

import FormField from "@components/FormField/FormField";
import { AuthButtonContainer } from "@components/Auth/AuthButtons/AuthButtonsContainer/AuthButtonsContainer";

import {
  RegisterButtons,
  RegisterFields,
} from "@utils/constants/AuthField/RegisterFields/RegisterFields";
import { RegisterValidator } from "@utils/validators/Auth/RegisterValidator";

import {
  ButtonContainer,
  FieldCustom,
  FormContainer,
} from "@modules/Auth/styled/styled";

interface RegisterProps {
  onSubmit: any;
}

export const RegisterForm: FC<RegisterProps> = ({ onSubmit }) => {
  return (
    <Form
      onSubmit={onSubmit}
      validate={RegisterValidator}
      render={({ handleSubmit }) => (
        <FormContainer>
          <form onSubmit={handleSubmit}>
            {RegisterFields.map(({ type, name, label }) => (
              <FieldCustom
                type={type}
                name={name}
                placeholder={label}
                variant="standard"
                component={FormField}
              >
              </FieldCustom>
            ))}
            <ButtonContainer>
              {RegisterButtons.map(({ button }) => (
                <AuthButtonContainer button={button} />
              ))}
            </ButtonContainer>
          </form>
        </FormContainer>
      )}
    />
  );
};
