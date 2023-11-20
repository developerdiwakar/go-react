import { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <h1>Goreact - CRUD App</h1>
      <div className="card">
        <div className="card-body">
          <div className="card-title">
            <p className="fw-bold pt-3">Counter Button</p>
          </div>
          <div className="d-grid">
            <button
              className="btn btn-primary btn-block"
              onClick={() => setCount((count) => count + 1)}
            >
              count is <span className="fw-bold">{count}</span>
            </button>
          </div>
        </div>
      </div>
    </>
  );
}

export default App;
