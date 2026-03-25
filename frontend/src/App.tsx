import type { Component } from "solid-js";

import logo from "./logo.svg";
import styles from "./App.module.css";
import { HashRouter, Route } from "@solidjs/router";
import Home from "./routes/Home";

import "@fontsource/inter";

const App: Component = () => {
  return (
    <HashRouter root={Home}>
      <Route path="/" component={Home} />
    </HashRouter>
  );
};

export default App;
