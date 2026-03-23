import type { Component } from "solid-js";

import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import Main from "../layouts/Main";

const Home: Component = () => {
  return (
    <div class="min-h-screen bg-gray-900 text-white">
      <Navbar />
      <Main />
      <Footer />
    </div>
  );
};

export default Home;
