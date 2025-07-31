import './App.css';
import { useEffect, useState } from 'react';

function App() {
  const [todos, setTodos] = useState([])
  const [title, setTitle] = useState("")

  const fetchTodos = async()=>{
    try {
      const res = await fetch("http://localhost:8080/api/todos", {
        method : "GET"
      })
      const data = await res.json()
      setTodos(data)
    } catch (error) {
      console.error(error)
      setTodos([])
    }
  }

  const addTodos = async()=>{
    if (!title.trim()) return;
    try {
        const res = await fetch("http://localhost:8080/api/todos", {
          method : "POST",
          headers : {"Content-Type" : "application/json"},
          body : JSON.stringify({title})
        })
      setTitle("")
      fetchTodos()
    } catch (error) {
      console.log(error)
    }
  }

  const deleteTodos = async(id)=>{
    try {
        await fetch(`http://localhost:8080/api/todos/${id}`, {
          method : "DELETE",
        })
        fetchTodos()
    } catch (error) {
      console.log(error)
    }
  }

  useEffect(() => {
    fetchTodos()
  }, [])


  return (
    <div style={{ padding: 20 }}>
        <h1>TO-DO LIST</h1>
        <input value={title} onChange={(e) => setTitle(e.target.value)} placeholder='ENTER TASK HERE'></input>

        <button onClick={addTodos}>Add Task</button>

        <ul>
          {Array.isArray(todos) ? (
            todos.length > 0 ? (
              todos.map((todo, index) => (
                <li key={todo.id ?? index}>
                  {todo.title ?? "Untitled"}
                  <button onClick={()=>deleteTodos(todo.id)}>X</button>
                </li>
              ))
            ) : (
                <h3>Kosong</h3>
            )
          ) : (
            <li>
              Belum pernah diisi
            </li>
          )
        }
        </ul>
      </div>
  );
}

export default App;
