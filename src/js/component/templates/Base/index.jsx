import React from "react";
import PropTypes from "prop-types";
import styles from "./index.sass";
import { Header } from "../../organisms/";

class Base extends React.Component {
  render() {
    //TODO write component detail
    return (
      <div className={styles.root}>
        <Header />
        <section className={styles.content}>
          <div className={styles.body}>aiueo,kakikukeko,sasisuseso</div>
          <div className={styles.nav}>asdf</div>
        </section>
      </div>
    );
  }
}

Base.propTypes = {
  //TODO props: PropTypes.string.isRequired
};

export default Base;
