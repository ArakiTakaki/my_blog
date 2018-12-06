import React from "react";
import PropTypes from "prop-types";
import styles from "./index.sass";

class ArticleExcerpt extends React.Component {
  render() {
    //TODO write component detail
    return (
      <div className={styles.root}>
        <div className={styles.header}>
          <h1>sample</h1>
        </div>
        <div className={styles.image} />
        <div className={styles.content}>
          <p>ExcerptContent</p>
        </div>
        <div className={styles.foot_action}>
          <button>MORE</button>
        </div>
      </div>
    );
  }
}

ArticleExcerpt.propTypes = {
  //TODO props: PropTypes.string.isRequired
};

export default ArticleExcerpt;
