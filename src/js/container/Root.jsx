import React from "react";
import { hot } from "react-hot-loader";
import { Index } from "../pages";

@hot(module)
class Root extends React.Component {
  render() {
    return <Index />;
  }
}

export default Root;
