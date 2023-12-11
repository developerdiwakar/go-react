function Navbar() {
  return (
    <nav className="navbar navbar-light bg-light p-3">
      <div className="d-flex col-12 col-md-3 col-lg-2 mb-2 mb-lg-0 flex-wrap flex-md-nowrap justify-content-between">
        <a className="navbar-brand" href="#">
          GoReact Dashboard
        </a>
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
        {/* <div className="mr-3 mt-1">
        <a
          target="_blank"
          className="github-button"
          href="https://github.com/developerdiwakar/go-react"
          data-color-scheme="no-preference: dark; light: light; dark: light;"
          data-icon="octicon-star"
          data-size="large"
          data-show-count="true"
          aria-label="Star /developerdiwakar/go-react"
        >
          Star
        </a>
      </div> */}
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
              <a className="dropdown-item" href="#">
                Settings
              </a>
            </li>
            <li>
              <a className="dropdown-item" href="#">
                Messages
              </a>
            </li>
            <li>
              <a className="dropdown-item" href="#">
                Sign out
              </a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
