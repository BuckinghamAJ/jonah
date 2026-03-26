import type { Component } from "solid-js";

import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import Main from "../layouts/Main";

const Home: Component = (props) => {
  return (
    <div class="min-h-screen bg-slate-900 text-white overflow-y-auto">
      <Navbar />
      {props.children}
      <Footer />
    </div>
  );
};

export default Home;
