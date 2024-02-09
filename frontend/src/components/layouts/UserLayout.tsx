import { Outlet, Navigate } from "react-router-dom";
import { checkIsLoggedIn } from "../../utils/auth";
import Layout from "./Layout";
import Sidebar from "../main/Sidebar";

export function UserLayout() {
  return checkIsLoggedIn() ? (
    <Layout>
      <Sidebar />
      <Outlet />
    </Layout>
  ) : (
    <Navigate to={"/login"} />
  );
}
