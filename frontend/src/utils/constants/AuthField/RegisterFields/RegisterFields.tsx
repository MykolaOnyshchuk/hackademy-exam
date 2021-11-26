import { Button } from "@mui/material";
export const RegisterFields = [
  {
    type: "text",
    name: "email",
    label: "Email",
    placeholder: "Email",
  },
  {
    type: "password",
    name: "password",
    label: "Password",
    placeholder: "Password",
  },
  {
    type: "password",
    name: "confirmPassword",
    label: "Confirm Password",
    placeholder: "Confirm Password",
  },
];

export const RegisterButtons = [
  {
    button: (
      <Button
        size="large"
        color="inherit"
        variant="contained"
        type="submit"
      >
        Sign Up
      </Button>
    ),
  },
];
