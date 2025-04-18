import { useEffect, useState } from "react";
import { ProcessAttempt } from "../api/gameService";
import { useCharacter } from "../context/useCharacter";

const GamePage: React.FC = () => {
  const { character } = useCharacter();
  const [inputName, setInputName] = useState<string>("");
  const [selectedAnime, setSelectedAnime] = useState<string>("");
  interface AttemptResult {
    success: boolean;
    message: string;
  }

  const [attemptResult, setAttemptResult] = useState<AttemptResult | null>(null);

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
    } catch (error) {
      console.error("Error processing attempt:", error);
    }
  };

  return (
    <div>
      <h1>Game Page</h1>
        <input
            type="text"
            value={inputName}
            onChange={(e) => setInputName(e.target.value)}
            placeholder="Type the character name"
        />
        <button onClick={handleSearch}>Search</button>
        {attemptResult && (
        <div>
            <h3>Attempt Result:</h3>
            <pre>{JSON.stringify(attemptResult, null, 2)}</pre>
        </div>
        )}
    </div>
  );
};
export default GamePage;