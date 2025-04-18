import { useEffect, useState } from "react";
import { ProcessAttempt } from "../api/gameService";
import { useCharacter } from "../context/useCharacter";

const GamePage: React.FC = () => {
  const { character } = useCharacter();
  const [inputName, setInputName] = useState<string>("");
  const [selectedAnime, setSelectedAnime] = useState<string>("");
  const [attemptResult, setAttemptResult] = useState<any>(null);

  useEffect(() => {
    const savedAnime = localStorage.getItem("selectedAnime");
    if (savedAnime) {
      setSelectedAnime(savedAnime);
    }
  }, []);

  const handleSearch = async () => {
    if (!character) return;

    try {
      const response = await ProcessAttempt({
        name: inputName,
        anime: selectedAnime,
        currentCharacter: character,
      });
      setAttemptResult(response);      
      console.log("Resultado da tentativa:", response);
    } catch (error) {
      console.error("Erro ao tentar personagem:", error);
    }
  };

  return (
    <div>
      <h1>Game Page</h1>
        <input
            type="text"
            value={inputName}
            onChange={(e) => setInputName(e.target.value)}
            placeholder="Digite o nome do personagem"
        />
        <button onClick={handleSearch}>Search</button>
        {attemptResult && (
        <div>
            <h3>Resultado da tentativa:</h3>
            <pre>{JSON.stringify(attemptResult, null, 2)}</pre>
        </div>
        )}
    </div>
  );
};
export default GamePage;