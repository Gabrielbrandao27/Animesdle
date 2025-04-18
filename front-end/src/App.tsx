import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { startGame } from "./api/gameService";
import "./App.css";
import { useCharacter } from "./context/useCharacter";
import animedleLogo from "/naruto.svg";

function App() {
  const navigate = useNavigate();
  const { setCharacter } = useCharacter();

  const animeOptions = ["Naruto", "One Piece"];
  const [selectedAnime, setSelectedAnime] = useState(animeOptions[0]);

  const handleStartGame = async () => {
    try {
      const data = await startGame(selectedAnime);
      setCharacter(data);
      localStorage.setItem("selectedAnime", selectedAnime);
      navigate("/game");
    } catch (error) {
      console.error("Error starting game:", error);
    }
  };

  return (
    <>
      <div>
        <img src={animedleLogo} className="logo" alt="Animedle logo" />
      </div>
      <h1>Animesdle</h1>
      <p>Wordle like game but for Anime!</p>

      {/* Dropdown com as opções de anime */}
      <label htmlFor="anime-select">Choose an anime:</label>
      <select
        id="anime-select"
        value={selectedAnime}
        onChange={(e) => setSelectedAnime(e.target.value)}
      >
        {animeOptions.map((anime) => (
          <option key={anime} value={anime}>
            {anime}
          </option>
        ))}
      </select>

      <button className="start-game-button" onClick={handleStartGame}>
        Start Game
      </button>
    </>
  );
}

export default App;
