import { createStore, applyMiddleware } from "redux";
import initialState from "./state";
import logger from "redux-logger";
import reducers from "./reducers";

const middlewares = [];

// redux-logger
if (process.env.NODE_ENV === "development") {
  middlewares.push(logger);
}

const store = createStore(
  reducers,
  initialState,
  applyMiddleware(...middlewares)
);

export default store;
