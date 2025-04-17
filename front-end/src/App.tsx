import { useNavigate } from 'react-router-dom';
import './App.css';
import animedleLogo from '/naruto.svg';

function App() {
  const navigate = useNavigate();

  const handleStartGame = () => {
    navigate('/start-game');
  };

  return (
    <>
      <div>
        <img src={animedleLogo} className="logo" alt="Animedle logo" />
      </div>
      <h1>Animesdle</h1>
      <p>Wordle like game but for Anime!</p>
      <button className="start-game-button" onClick={handleStartGame}>
        Start Game
      </button>
    </>
  )
}

export default App
