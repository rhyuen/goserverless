import React, { Component } from "react";
import { hot } from "react-hot-loader";
import axios from "axios";

class App extends Component {
  state = {
    serverdata: "",
    name: "",
    password: "",
    hashedResponse: "Nothing so far."
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

  handleSubmit = e => {
    e.preventDefault();
    const { name, password } = this.state;
    const payload = { name, password };

    axios
      .post("/api/signup", payload)
      .then(res => {
        console.log(res.data);
        this.setState(ps => {
          return {
            serverdata: ps.serverdata,
            name: "",
            password: "",
            hashedResponse: res.data.password
          };
        });
      })
      .catch(e => {
        console.log(e);
        this.setState(ps => {
          return {
            name: "",
            password: ""
          };
        });
      });
  };

  handleFormChange = e => {
    const { name, value } = e.target;
    this.setState({
      [name]: value
    });
  };

  render() {
    return (
      <div className="root">
        <h1>Dummy Data</h1>
        <h1>{this.state.serverdata.title}</h1>
        <h1>{this.state.serverdata.author}</h1>
        <h2>Response thus far: {this.state.hashedResponse}</h2>
        <div>
          <form onSubmit={this.handleSubmit}>
            <input
              type="text"
              name="name"
              value={this.state.name}
              onChange={this.handleFormChange}
              placeholder="name"
            />
            <br />
            <input
              type="password"
              name="password"
              value={this.state.password}
              onChange={this.handleFormChange}
              placeholder="password"
            />
            <br />
            <input type="submit" value="Login" />
          </form>
        </div>
      </div>
    );
  }
}

export default hot(module)(App);
