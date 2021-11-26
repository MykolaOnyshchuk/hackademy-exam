import { Field, Form } from "react-final-form";
import styled from "styled-components";

export const FormContainer = styled.div`
  form {
    width: 100%;
    margin: auto;
    margin-top: 5%;
  }
`;
export const FieldCustom = styled(Field)`
  width: 100%;
  margin: 10px 0 !important;

  .css-1q6at85-MuiInputBase-root-MuiOutlinedInput-root.Mui-focused
    .MuiOutlinedInput-notchedOutline {
    border-color: rgb(122, 122, 122);
  }
  .css-1kty9di-MuiFormLabel-root-MuiInputLabel-root.Mui-focused {
    color: black;
  }
`;

export const ButtonContainer = styled.div`
  width: 100%;
  display: flex;
  justify-content: flex-end;
  margin-top: 15%;
  .MuiButton-root{
    height: 32px;
    width: 108px;
    background-color: #FCD620;
  }
  a {
    text-decoration: none;
    color: black;
  }
`;
