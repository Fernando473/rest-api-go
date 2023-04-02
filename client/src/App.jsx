import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [name, setName] = useState("");
  const [users, setUsers] = useState([]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch(import.meta.env.VITE_API + "/users", {
      method: "POST",
      body: JSON.stringify({ name }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await response.json();
    fetchUsers();
  };
  const fetchUsers = async () => {
    const response = await fetch(import.meta.env.VITE_API + "/users");
    const data = await response.json();
    setUsers(data.data);
    console.log(data);
  };
  useEffect(() => {
    fetchUsers();
  }, []);
  return (
    <div className="App">
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Coloca tu nombre"
          onChange={(e) => setName(e.target.value)}
        />
        <button type="submit">Guardar</button>
      </form>

      <ul>
        {users.map((user) => (
          <li>{user.name}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;
