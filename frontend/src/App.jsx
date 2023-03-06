import { useState } from "react";
import reactLogo from "./assets/react.svg";
import "./App.css";

function App() {
  //const [count, setCount] = useState(0);

  return (
    <div className="App">
      <button
        onClick={(e) => {
          e.preventDefault();
          fetch("http://localhost:3000/cookie", {
            // withCredentials: true,
            credentials: "include", //para guardar cookie en localhost
            headers: {
              "Content-Type": "application/json",
            },
          })
            .then((res) => {
              console.log(res);
              return res.json();
            })
            .then((res) => console.log(res));
        }}
      >
        Cookie
      </button>
      <button
        onClick={(e) => {
          e.preventDefault();
          fetch("http://localhost:3000/api/login", {
            method: "POST",
            credentials: "include", //para guardar cookie en localhost
            headers: {
              "Content-Type": "application/json",
            },
          })
            .then((res) => {
              console.log(res);
              return res.json();
            })
            .then((res) => console.log(res));
        }}
      >
        Login
      </button>
    </div>
  );
}

export default App;
