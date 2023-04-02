import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {

  return (
    <div className="App">
      <button onClick={async () => {
        const response  = await fetch('/users')
        const data = await response.json()
        console.log(data)
      }}> Obtener datos</button>
    </div>
  )
}

export default App
