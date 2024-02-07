import { Link } from "react-router-dom";

function PageNotFound() {
  return (
    <div className="container">
      <div
        className="d-flex flex-row align-items-center"
        style={{ minHeight: "100v" }}
      >
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-md-12 text-center">
              <span className="display-1 d-block">404</span>
              <div className="mb-4 lead">
                The page you are looking for was not found.
              </div>
              <Link to="/" className="btn btn-link">
                Back to Website Homepage
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default PageNotFound;
