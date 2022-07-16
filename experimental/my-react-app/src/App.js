import logo from "./logo.svg";
import "./App.css";

import React, { useState } from "react";


const App = () => {
  const defaultUrl = "https://api.github.com/zen";

  const [value, setValue] = React.useState(defaultUrl);
  const [output, setOutput] = React.useState("no data yet");

  const handleChange = (e) => setValue(e.target.value);

  const do_a_get = (e) => {
    console.log("doing a get request of", value);

    fetch(value)
      .then((response) => {
        return response.text();
      })
      .then((a) => {
        setOutput(a);
        console.log(a);
      });
  };

  return (
    <div className="App">
      <br />
      <br />
      <br />

      <textarea
        type="text"
        style={{ width: "95%", border: "solid 1px orange"}}
        value={value}
        onChange={handleChange}
      />

      <p>{value}</p>

      <br />
      <br />
      <br />
      <br />
      <br />
      <hr />
      <button 
        style={{ width: "200px", height: "30px", padding: "20px"}}
        onClick={do_a_get}>Do a get request
      </button>

      <br />
      <br />

      <p>{output}</p>
      <hr />
      <br />
      <br />
    </div>
  );
};

export default App;
