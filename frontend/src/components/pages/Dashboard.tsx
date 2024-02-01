import Breadcrumb from "../main/Breadcrumb";

interface BreadCrumbs {
  breadcrumbs: Array<any>;
}

function Dashboard(props: BreadCrumbs) {
  return (
    <main className="col-md-9 ml-sm-auto col-lg-10 px-md-4 py-4">
      <Breadcrumb breadcrumbs={props.breadcrumbs} />
      <h1 className="h2">Dashboard</h1>
      <p>
        This is the homepage of a simple admin interface which is part of a
        tutorial written on Themesberg
      </p>

      <div className="row my-4">
        <div className="col-12 col-md-6 col-lg-3 mb-4 mb-lg-0">
          <div className="card">
            <h5 className="card-header">Customers</h5>
            <div className="card-body">
              <h5 className="card-title">345k</h5>
              <p className="card-text">Feb 1 - Apr 1, United States</p>
              <p className="card-text text-success">
                18.2% increase since last month
              </p>
            </div>
          </div>
        </div>
        <div className="col-12 col-md-6 mb-4 mb-lg-0 col-lg-3">
          <div className="card">
            <h5 className="card-header">Revenue</h5>
            <div className="card-body">
              <h5 className="card-title">$2.4k</h5>
              <p className="card-text">Feb 1 - Apr 1, United States</p>
              <p className="card-text text-success">
                4.6% increase since last month
              </p>
            </div>
          </div>
        </div>
        <div className="col-12 col-md-6 mb-4 mb-lg-0 col-lg-3">
          <div className="card">
            <h5 className="card-header">Purchases</h5>
            <div className="card-body">
              <h5 className="card-title">43</h5>
              <p className="card-text">Feb 1 - Apr 1, United States</p>
              <p className="card-text text-danger">
                2.6% decrease since last month
              </p>
            </div>
          </div>
        </div>
        <div className="col-12 col-md-6 mb-4 mb-lg-0 col-lg-3">
          <div className="card">
            <h5 className="card-header">Traffic</h5>
            <div className="card-body">
              <h5 className="card-title">64k</h5>
              <p className="card-text">Feb 1 - Apr 1, United States</p>
              <p className="card-text text-success">
                2.5% increase since last month
              </p>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
}

export default Dashboard;
