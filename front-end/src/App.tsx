import { useNavigate } from "react-router-dom";
import { startGame } from "./api/gameService";
import "./App.css";
import { useCharacter } from "./context/useCharacter";
import animedleLogo from "/naruto.svg";

function App() {
  const navigate = useNavigate();
  const { setCharacter } = useCharacter();

  const handleStartGame = async () => {
    try {
      const data = await startGame("Naruto"); // ou "One Piece", futuramente podemos colocar dropdown
      setCharacter(data);
      navigate("/game");
    } catch (error) {
      console.error("Erro ao iniciar jogo:", error);
    }
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
  );
}

export default App;
