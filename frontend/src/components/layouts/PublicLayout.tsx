import { Outlet } from "react-router-dom";
import Layout from "./Layout";

export default function PublicLayout() {
  return (
    <Layout>
      <Outlet />
    </Layout>
  );
}
