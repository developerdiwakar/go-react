import { ReactNode } from "react";
import { checkIsLoggedIn } from "../../utils/auth";
import Navbar from "../main/Navbar";

interface Props {
  children?: ReactNode;
}
export default function Layout({ children }: Props) {
  return (
    <div className="contanier">
      <Navbar isLoggedIn={checkIsLoggedIn()} />
      <div className="container-fluid">
        <div className="row">{children}</div>
      </div>
    </div>
  );
}
