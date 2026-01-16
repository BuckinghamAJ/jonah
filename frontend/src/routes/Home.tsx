import type { Component } from "solid-js";
import { Greet } from "../wailsjs/go/main/App";

import Navbar from "../components/Navbar";
import { createResource } from "solid-js";
import Footer from "../components/Footer";
import Main from "../components/Main";

const Home: Component = () => {
  // Testing
  return (
    <div class="min-h-screen bg-gray-900 text-white">
      <Navbar />
      <Main />
      <Footer />
    </div>
  );
};

export default Home;
