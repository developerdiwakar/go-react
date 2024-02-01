import { NavLink } from "react-router-dom";

function Navbar() {
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
          <NavLink
            // target="_blank"
            className=""
            to="/login"
          >
            Login
          </NavLink>
        </div>
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
              <NavLink className="dropdown-item" to="#">
                Settings
              </NavLink>
            </li>
            <li>
              <NavLink className="dropdown-item" to="#">
                Messages
              </NavLink>
            </li>
            <li>
              <NavLink className="dropdown-item" to="#">
                Sign out
              </NavLink>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
