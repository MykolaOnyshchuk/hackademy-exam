import styled from "styled-components";
export const AuthContainer = styled.div`
  position: relative;
  width: 100%;
  height: 100vh;
  background: linear-gradient(to right, #fdfaa8, white);
  display: flex;
  align-items: center;
`;
export const BackgroundImg = styled.img`
  position: absolute;
  top: 20px;
  left: 20px;
  width: 883px;
  z-index: 0;
`;

export const AuthBlock = styled.div`
  z-index: 2;
  margin-left: auto;
  margin-right: auto;
  width: 440px;
  padding: 30px 15px;
  background-color: white;
  @media (max-width: 1160px) {
    width: 50%;
    margin-top: 15%;
  }
  @media (max-width: 665px) {
    width: 70%;
    margin-top: 25%;
  }
`;

export const AuthTitle = styled.div`
  color: rgb(42, 42, 42);
  font-size: 18px;
  font-weight: 400;
`;
export const AuthLogo = styled.div`
  color: rgb(42, 42, 42);
  font-size: 24px;
  font-weight: 400;
  padding-top: 25px;
  &img {
    width: 35px;
  }
`;
