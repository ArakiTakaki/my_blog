import "../sass/base.sass";
import ReactDOM from "react-dom";
import React from "react";
import { Provider } from "react-redux";

import Root from "./container/Root";
import store from "./store/configreStore";

ReactDOM.render(
  <Provider store={store}>
    <Root />
  </Provider>,
  document.getElementById("app")
);
