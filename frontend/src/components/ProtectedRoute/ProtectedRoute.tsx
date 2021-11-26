import { FC } from "react";
import { Route } from "react-router";
import { Header } from "@components/Header/Header";
import { SideMenu } from "@components/SideMenu/SideMenu";



import { SideContainer, GlobalContainer, SideChildren } from "./styled/styled";

interface RouteProps {
  path: string;
}

export const ProtectedRoute: FC<RouteProps> = ({
  children,
  ...rest
}) => {
  return (
    <Route
      exact
      {...rest}
      render={() => (
        <GlobalContainer>
          <SideMenu />
          <SideContainer>
            <Header />
            <SideChildren>{children}</SideChildren>
          </SideContainer>
        </GlobalContainer>
      )}
    />
  );
};
