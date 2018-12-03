import * as Actions from "./actions";

const reducer = (state, action) => {
  const { payload } = action;
  switch (action.type) {
    case Actions.INCREMENT:
      return { ...state, num: state.num++ };
    case Actions.DECREMENT:
      return { ...state, num: state.num-- };
    default:
      return state;
  }
};

export default reducer;
