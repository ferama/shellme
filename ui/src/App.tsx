import React from 'react'
import { MainContainerWithRouter } from './MainContainer'
import './App.css'
import {
  // BrowserRouter as Router,
  HashRouter as Router
} from "react-router-dom";

const App = () => {
  return (
    <div className="App">
      <Router>
        <MainContainerWithRouter />
      </Router>
    </div>
  )
}

export default App
