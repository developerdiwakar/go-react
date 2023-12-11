import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/css/bootstrap-utilities.min.css";
import "./App.css";
import Navbar from "./components/main/Navbar";
import Sidebar from "./components/main/Sidebar";
import PageBody from "./components/main/PageBody";

function App() {
  return (
    <div>
      <Navbar />
      <div className="container-fluid">
        <div className="row">
          <Sidebar />
          <PageBody />
        </div>
      </div>
    </div>
  );
}

export default App;
