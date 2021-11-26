import { Button } from "@mui/material";
export const LoginFields = [
  {
    type: "text",
    name: "email",
    label: "Email",
  },
  {
    type: "password",
    name: "password",
    label: "Password",
  },
];

export const LoginButtons = [
  {
    button: (
      <Button
        size="large"
        color="inherit"
        variant="contained"
        type="submit"
      >
        Sign In
      </Button>
    ),
  }
];
