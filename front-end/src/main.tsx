import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import App from './App.tsx'
import GamePage from './components/GamePage.tsx'
import { CharacterProvider } from './context/CharacterContext'
import './index.css'


ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <CharacterProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="/game" element={<GamePage />} />
        </Routes>
      </BrowserRouter>
    </CharacterProvider>
  </React.StrictMode>,
)

