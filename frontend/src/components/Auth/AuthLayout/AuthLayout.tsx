import { FC } from "react";
import { AuthBlock, AuthContainer, AuthLogo, AuthTitle, BackgroundImg } from "./styled/styled";
import logo from "../../../img/logo.png";
import background from "../../../img/Group.png";
interface AuthProps {
  title: string;
}

export const AuthLayout: FC<AuthProps> = ({ children, title }) => {
  return (
    <AuthContainer>
      <BackgroundImg src={background} alt="" />
      <AuthBlock>
      
        <AuthLogo>
          <img src={logo} alt="" />
          Todo
        </AuthLogo>
        <AuthTitle>{title}</AuthTitle>
        {children}
      </AuthBlock>
    </AuthContainer>
  );
};
