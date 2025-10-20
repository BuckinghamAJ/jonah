import type { Component } from "solid-js";
import { Greet } from "../wailsjs/go/main/App";

import logo from "./logo.svg";
import styles from "./App.module.css";
import Navbar from "./components/Navbar";
import { createResource } from "solid-js";
import Footer from "./components/Footer";
import Main from "./components/Main";

const App: Component = () => {
  const [greetAdam] = createResource(() => Greet("Adam"));

  return (
    <div class="min-h-screen">
      <Navbar />
      <Main />
      <Footer />
    </div>
  );
};

export default App;
