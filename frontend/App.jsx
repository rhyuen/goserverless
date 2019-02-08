import React, { Component } from "react";
import { hot } from "react-hot-loader";
import axios from "axios";

class App extends Component {
  state = {
    serverdata: ""
  };

  componentDidMount() {
    axios
      .get("/api/index")
      .then(data => {
        this.setState({ serverdata: data });
      })
      .catch(e => {
        console.log(e);
      });
  }

  render() {
    return (
      <div className="root">
        hi, root, latest version<h1>{this.state.serverdata}</h1>
      </div>
    );
  }
}

export default hot(module)(App);
