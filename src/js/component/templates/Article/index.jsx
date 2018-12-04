import React from "react";
import PropTypes from "prop-types";
import styles from "./index.sass";
import { Header } from "../../organisms/";

class Article extends React.Component {
  render() {
    //TODO write component detail
    return (
      <div>
        <Header />
      </div>
    );
  }
}

Article.propTypes = {
  //TODO props: PropTypes.string.isRequired
};

export default Article;
