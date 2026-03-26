import type { Component } from "solid-js";

import logo from "./logo.svg";
import styles from "./App.module.css";
import { HashRouter, Route } from "@solidjs/router";
import Home from "./routes/Home";

import "@fontsource/inter";
import Chapter from "./routes/Chapter";
import Main from "./layouts/Main";

const App: Component = () => {
  return (
    <HashRouter root={Home}>
      <Route path="/" component={Main} />
      <Route path="/read/:bookId/:chapterId" component={Chapter} />
    </HashRouter>
  );
};

export default App;
