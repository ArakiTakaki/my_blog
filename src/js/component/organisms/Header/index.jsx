import React from "react";
import PropTypes from "prop-types";
import styles from "./index.sass";

class Header extends React.Component {
  render() {
    //TODO write component detail
    return (
      <div className={styles.root}>
        <div className={styles.title}>
          <h1>ORE BLOG</h1>
        </div>
        <div>三</div>
      </div>
    );
  }
}

Header.propTypes = {
  //TODO props: PropTypes.string.isRequired
};

export default Header;
