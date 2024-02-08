import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/css/bootstrap-utilities.min.css";
import "./App.css";
import Navbar from "./components/main/Navbar";
import Sidebar from "./components/main/Sidebar";
import { Outlet, useNavigate } from "react-router-dom";
import AppRoutes from "./components/main/AppRoutes";
import { ReactNode, useEffect } from "react";
import { checkIsLoggedIn } from "./utils/auth";

interface Props {
  children?: ReactNode;
}
// AppRoutes contains all the routes for the app
function App() {
  return <AppRoutes />;
}

function PageLayout({ children }: Props) {
  return (
    <div className="contanier">
      <Navbar isLoggedIn={checkIsLoggedIn()} />
      <div className="container-fluid">
        <div className="row">{children}</div>
      </div>
    </div>
  );
}
export function Layout() {
  return (
    <PageLayout>
      <Outlet />
    </PageLayout>
  );
}

export function ProtectedLayout() {
  const history = useNavigate();
  useEffect(() => {
    // Redirect to Login if not logged in
    if (!checkIsLoggedIn()) {
      return history("/login");
    }
  }, [null]);
  return (
    <PageLayout>
      <Sidebar />
      <Outlet />
    </PageLayout>
  );
}

export default App;
