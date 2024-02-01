import { Link } from "react-router-dom";

interface BreadCrumbs {
  breadcrumbs: Array<any>;
}
export default function Breadcrumb(props: BreadCrumbs) {
  return (
    <nav aria-label="breadcrumb">
      <ol className="breadcrumb">
        <li className="breadcrumb-item">
          <Link to="/">Home</Link>
        </li>
        {props.breadcrumbs &&
          props.breadcrumbs.map((item, index) => {
            if (props.breadcrumbs.length - 1 === index) {
              return (
                <li
                  key={index}
                  className="breadcrumb-item active"
                  aria-current="page"
                >
                  {item.atext}
                </li>
              );
            } else {
              return (
                <li key={index} className="breadcrumb-item">
                  <Link to={item.ahref}>{item.atext}</Link>
                </li>
              );
            }
          })}
      </ol>
    </nav>
  );
}
