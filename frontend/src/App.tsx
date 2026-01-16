import type { Component } from "solid-js";
import { Greet } from "../wailsjs/go/main/App";

import logo from "./logo.svg";
import styles from "./App.module.css";
import { HashRouter, Route } from "@solidjs/router";
import Home from "./routes/Home"
const App: Component = () => {

  return (
    <HashRouter root={Home}>
      <Route path="/" component={Home} />

    </HashRouter>

  );
};

export default App;
