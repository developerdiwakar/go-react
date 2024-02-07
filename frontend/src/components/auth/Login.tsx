// import { useState} from 'react'

import { Link, useNavigate } from "react-router-dom";
import Breadcrumb from "../main/Breadcrumb";
import { useEffect, useState } from "react";
import axios from "axios";
import { setAuthToken } from "../../utils/auth";

type LoginProps = {
  breadcrumbs: Array<any>;
  isLoggedIn: Boolean;
};

const useFormInput = (initialValue: string) => {
  const [value, setValue] = useState(initialValue);

  const handleChange = (e: any) => {
    setValue(e.target.value);
  };

  return { value, onChange: handleChange };
};

export default function Login(props: LoginProps) {
  const history = useNavigate();
  const username = useFormInput("");
  const password = useFormInput("");
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    // Redirect to Dashboard if logged in
    if (props.isLoggedIn) {
      history("/user/dashboard");
    }
  });

  // handle button click of login form
  const handleLogin = () => {
    setError(null);
    setLoading(true);
    // Calling Login API
    axios
      .post("http://127.0.0.1:8000/login", {
        username: username.value,
        password: password.value,
      })
      .then((response) => {
        console.log(response);
        // Login Sucess
        history("/user/dashboard");
        setLoading(false);
        setAuthToken(response.data.data.token);
      })
      .catch((error) => {
        // Login failed
        setLoading(false);
        if (error.data.success === false) {
          setError(error.message);
        }
      });
  };

  //return the content
  return (
    <>
      <main className="col-12 col-md-12 col-lg-12 px-md-4 py-4">
        <Breadcrumb breadcrumbs={props.breadcrumbs} />
        <h1 className="h2">Login</h1>
        <p>Login here to access the dashboad.</p>

        <div className="row my-4">
          <div className="col-12 col-md-8 col-lg-6 mb-4 mb-lg-0">
            <form>
              <div className="form-row">
                <div className="col-sm-12 col-md-10 col-lg-6  mb-3">
                  <label htmlFor="username">Username</label>
                  <input
                    type="text"
                    className={"form-control " + (error ? "is-invalid" : "")}
                    id="username"
                    placeholder="Email or Mobile Number"
                    autoComplete="new-password"
                    {...username}
                    required
                  />
                  {/* <div className="valid-feedback">Looks OK!</div> */}
                </div>
                <div className="col-sm-12 col-md-10 col-lg-6  mb-3">
                  <label htmlFor="password">Password</label>
                  <input
                    type="text"
                    className={"form-control " + (error ? "is-invalid" : "")}
                    id="password"
                    placeholder="Password"
                    autoComplete="new-password"
                    {...password}
                    required
                  />
                  {/* <div className="valid-feedback"></div>
                  <div className="invalid-feedback"></div> */}
                </div>
              </div>
              <button
                className="btn btn-primary"
                type="submit"
                disabled={loading}
                onClick={handleLogin}
              >
                {loading ? "logging in..." : "Login"}
              </button>

              {error ? (
                <div className="invalid-feedback">{error}</div>
              ) : (
                <div className="valid-feedback">Login Successful!</div>
              )}
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
