import { useState, useEffect } from "react";
import reactLogo from "./assets/react.svg";
import "./App.css";
import axios from "axios";

function App() {
  const [count, setCount] = useState(0);
  const [name, setName] = useState("");
  const [queryName, setQueryName] = useState("");
  const [status, setStatus] = useState(0);
  const [found, setFound] = useState(false);
  const words = ["Fun", "Cool", "Fast", "Nice"];
  const [blank, setBlank] = useState("");

  const addUser = async () => {
    setStatus(1);
    let response = await axios.post("/api/users", {
      name: name,
    });
    console.log(response);
    if (response.status > 200 && response.status < 300) {
      setStatus(200);
    }
  };

  const getUser = async () => {
    setFound(false);
    let response = await axios.get(`/api/user/${queryName}`);
    console.log(response);
    if (response.status === 200) {
      setFound(true);
    }
  };

  useEffect(() => {
    setBlank(words[count % words.length]);
    // every 3 seconds change the word to a random one
    const interval = setInterval(() => {
      setBlank(words[Math.floor(Math.random() * words.length)]);
    }, 1000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="App">
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src="/vite.svg" className="logo" alt="Vite logo" />
        </a>
        <a href="https://reactjs.org" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>
        Vite is <span>{blank}</span>
      </h1>
      <div className="card">
        <button
          onClick={() => setCount((count) => count + 1)}
          style={{ marginBottom: "2rem" }}
        >
          Count {count}
        </button>
        <span style={{ display: "flex", flexDirection: "column" }}>
          <input
            type="text"
            name="name"
            onChange={(e) => setName(e.target.value)}
          />
          <br></br>
          <button type="submit" onClick={addUser}>
            Add User
          </button>
          <p>
            {status === 200
              ? "User Added To Database"
              : "User Not Added To Database"}
          </p>
        </span>
        <span style={{ display: "flex", flexDirection: "column" }}>
          <input
            type="text"
            name="name"
            onChange={(e) => setQueryName(e.target.value)}
          />
          <br></br>
          <button type="submit" onClick={getUser}>
            Get User
          </button>
          <p>
            {found === true ? "User Found In Database" : "User Does Not Exist"}
          </p>
        </span>
      </div>
    </div>
  );
}

export default App;
