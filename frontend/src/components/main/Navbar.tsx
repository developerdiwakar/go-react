import { Link, NavLink, useNavigate } from "react-router-dom";
import { removeAuthToken } from "../../utils/auth";

type NavProps = {
  isLoggedIn: Boolean;
};

function Navbar(props: NavProps) {
  const history = useNavigate();
  const handleLogout = () => {
    removeAuthToken();
    return history("/login");
  };

  return (
    <nav className="navbar navbar-light bg-light p-3">
      <div className="d-flex col-12 col-md-3 col-lg-2 mb-2 mb-lg-0 flex-wrap flex-md-nowrap justify-content-between">
        <NavLink className="navbar-brand" to="/user/dashboard">
          Dashboard
        </NavLink>
        <button
          className="button-toggler d-md-none collapsed mb-3"
          type="button"
          data-toggle="collapse"
          data-target="#sidebar"
          aria-controls="sidebar"
          aria-expanded="false"
          aria-label="Toggle Navigation"
        >
          <span className="navbar-toggler-icon"></span>
        </button>
      </div>
      <div className="col-12 col-md-4 col-lg-2">
        <input
          className="form-control form-control-dark"
          type="text"
          placeholder="Search..."
        />
      </div>
      <div className="col-12 col-md-5 col-lg-8 d-flex align-items-center justify-content-md-end">
        <div className="mr-3 mt-1 m-2">
          {!props.isLoggedIn && (
            <NavLink className="m-1" to="/login">
              Login
            </NavLink>
          )}
        </div>
        {props.isLoggedIn && (
          <div className="dropdown">
            <button
              className="btn btn-secondary dropdown-toggle"
              type="button"
              id="dropdownMenuButton"
              data-toggle="dropdown"
              aria-expanded="false"
            >
              Hello, John Doe
            </button>
            <ul className="dropdown-menu" aria-labelledby="dropdownMenuButton">
              <li>
                <Link className="dropdown-item" to="#">
                  Settings
                </Link>
              </li>
              <li>
                <Link className="dropdown-item" to="#">
                  Messages
                </Link>
              </li>
              <li>
                <a
                  className="dropdown-item"
                  href="#"
                  onClick={() => handleLogout()}
                >
                  Sign out
                </a>
              </li>
            </ul>
          </div>
        )}
      </div>
    </nav>
  );
}

export default Navbar;
