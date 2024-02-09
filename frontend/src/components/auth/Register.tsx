import { Link, useNavigate } from "react-router-dom";
import Breadcrumb from "../main/Breadcrumb";
import { useEffect } from "react";

type BreadCrumbs = {
  breadcrumbs: Array<any>;
  isLoggedIn: Boolean;
};

export default function Register(props: BreadCrumbs) {
  const history = useNavigate();
  useEffect(() => {
    if (props.isLoggedIn) {
      history("/user/dashboard");
    }
  });
  return (
    <main className="col-md-12 col-lg-12 px-md-4 py-4">
      <Breadcrumb breadcrumbs={props.breadcrumbs} />
      <h1 className="h2">Register</h1>
      <p>Register here to create a new account.</p>

      <div className="row my-4">
        <div className="col-12 col-md-6 col-lg-10 mb-4 mb-lg-0">
          <form>
            <div className="form-row">
              <div className="col-sm-12 col-md-10 col-lg-6  mb-3">
                <label htmlFor="name">Full Name</label>
                <input
                  type="text"
                  className="form-control"
                  id="name"
                  placeholder="Foo Bar"
                  required
                />
                <div className="valid-feedback">Looks good!</div>
              </div>
              <div className="col-sm-12 col-md-10 col-lg-6  mb-3">
                <label htmlFor="mobileNumber">Mobile Number</label>
                <div className="input-group">
                  <div className="input-group-prepend">
                    <span className="input-group-text" id="inputGroupPrepend3">
                      +91
                    </span>
                  </div>
                  <input
                    type="text"
                    className="form-control"
                    id="mobileNumber"
                    placeholder="923xxxxx01"
                    aria-describedby="inputGroupPrepend3"
                    required
                  />
                  <div className="invalid-feedback">
                    Please enter a mobile number.
                  </div>
                </div>
              </div>
              <div className="col-sm-12 col-md-10 col-lg-6  mb-3">
                <label htmlFor="email">Email Address</label>
                <div className="input-group">
                  <div className="input-group-prepend">
                    <span className="input-group-text" id="emailGroup">
                      @
                    </span>
                  </div>
                  <input
                    type="text"
                    className="form-control"
                    id="email"
                    placeholder="foobar@example.com"
                    aria-describedby="emailGroup"
                    required
                  />
                  <div className="valid-feedback">Looks good!</div>
                </div>
              </div>
              <div className="col-sm-12 col-md-10 col-lg-6  mb-3">
                <label htmlFor="password">Password</label>
                <input
                  type="text"
                  className="form-control"
                  id="password"
                  placeholder="Password"
                  required
                />
                <div className="valid-feedback">Looks good!</div>
              </div>
              <div className="col-sm-12 col-md-10 col-lg-6  mb-3">
                <label htmlFor="confirmPassword">Confirm Password</label>
                <input
                  type="text"
                  className="form-control"
                  id="confirmPassword"
                  placeholder="Re-enter password"
                  required
                />
                <div className="valid-feedback">Looks good!</div>
              </div>
            </div>
            <button className="btn btn-primary" type="submit">
              Create new account
            </button>
          </form>
          <p className="mt-3">
            Already have an account? <Link to="/login">Please login here</Link>
          </p>
        </div>
      </div>
    </main>
  );
}
