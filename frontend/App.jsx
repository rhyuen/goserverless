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
      .then(res => {
        this.setState({ serverdata: res.data });
      })
      .catch(e => {
        console.log(e);
      });
  }

  render() {
    return (
      <div className="root">
        <h1>Dummy Data</h1>
        <h1>{this.state.serverdata.title}</h1>
        <h1>{this.state.serverdata.author}</h1>
      </div>
    );
  }
}

export default hot(module)(App);
