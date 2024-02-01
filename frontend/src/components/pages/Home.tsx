function Home() {
  return (
    <main className="col-md-9 ml-sm-auto col-lg-10 px-md-4 py-4">
      {/* <Breadcrumb breadcrumbs={props.breadcrumbs} /> */}
      <h1 className="h2">Home Page</h1>
      <p>This is the homepage that will contains the list of the website.</p>

      <div className="row my-4">
        <div className="col-12 col-md-6 col-lg-3 mb-4 mb-lg-0">
          <div className="card">
            <h5 className="card-header">Card One</h5>
            <div className="card-body">
              <h5 className="card-title">345k</h5>
              <p className="card-text">Feb 1 - Apr 1, United States</p>
            </div>
          </div>
        </div>
        <div className="col-12 col-md-6 mb-4 mb-lg-0 col-lg-3">
          <div className="card">
            <h5 className="card-header">Card Two</h5>
            <div className="card-body">
              <h5 className="card-title">$2.4k</h5>
              <p className="card-text">Feb 1 - Apr 1, United States</p>
            </div>
          </div>
        </div>
        <div className="col-12 col-md-6 mb-4 mb-lg-0 col-lg-3">
          <div className="card">
            <h5 className="card-header">Card Three</h5>
            <div className="card-body">
              <h5 className="card-title">43</h5>
              <p className="card-text">Feb 1 - Apr 1, United States</p>
            </div>
          </div>
        </div>
        <div className="col-12 col-md-6 mb-4 mb-lg-0 col-lg-3">
          <div className="card">
            <h5 className="card-header">Card Four</h5>
            <div className="card-body">
              <h5 className="card-title">64k</h5>
              <p className="card-text">Feb 1 - Apr 1, United States</p>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
}

export default Home;
