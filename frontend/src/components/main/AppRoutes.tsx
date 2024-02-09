import { Route, Routes } from "react-router-dom";
import Home from "../pages/Home";
import Register from "../auth/Register";
import Login from "../auth/Login";
import Dashboard from "../pages/Dashboard";
import PageNotFound from "../pages/PageNotFound";
import { checkIsLoggedIn } from "../../utils/auth";
import { UserLayout } from "../layouts/UserLayout";
import PublicLayout from "../layouts/PublicLayout";

export default function AppRoutes() {
  return (
    <Routes>
      <Route path="*" element={<PageNotFound />} />
      <Route path="/" element={<PublicLayout />}>
        <Route path="/" element={<Home />} />
        <Route
          path="register"
          element={
            <Register
              isLoggedIn={checkIsLoggedIn()}
              breadcrumbs={[{ ahref: "#", atext: "Register" }]}
            />
          }
        />
        <Route
          path="login"
          element={
            <Login
              isLoggedIn={checkIsLoggedIn()}
              breadcrumbs={[{ ahref: "#", atext: "Login" }]}
            />
          }
        />
      </Route>
      <Route path="/user" element={<UserLayout />}>
        <Route
          path="dashboard"
          element={
            <Dashboard breadcrumbs={[{ ahref: "#", atext: "Dashboard" }]} />
          }
        />
      </Route>
    </Routes>
  );
}
