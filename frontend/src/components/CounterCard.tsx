import { useState } from "react";

function CounterCard() {
  const [count, setCount] = useState(0);
  return (
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
  );
}

export default CounterCard;
