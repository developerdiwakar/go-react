// import { useState} from 'react'

import { Link } from "react-router-dom";
import Breadcrumb from "../main/Breadcrumb";

interface BreadCrumbs {
  breadcrumbs: Array<any>;
}
export default function Login(props: BreadCrumbs) {
  return (
    <>
      <main className="col-md-12 col-lg-12 px-md-4 py-4">
        <Breadcrumb breadcrumbs={props.breadcrumbs} />
        <h1 className="h2">Login</h1>
        <p>Login here to access the dashboad.</p>

        <div className="row my-4">
          <div className="col-12 col-md-6 col-lg-10 mb-4 mb-lg-0">
            <form>
              <div className="form-row">
                <div className="col-sm-12 col-md-10 col-lg-6  mb-3">
                  <label htmlFor="username">Username</label>
                  <input
                    type="text"
                    className="form-control"
                    id="username"
                    placeholder="Email or Mobile Number"
                    required
                  />
                  <div className="valid-feedback">Looks good!</div>
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
              </div>
              <button className="btn btn-primary" type="submit">
                Login
              </button>
            </form>
            <p className="mt-3">
              Don't have an account?{" "}
              <Link to="/register">Please register here</Link>
            </p>
          </div>
        </div>
      </main>
    </>
  );
}
