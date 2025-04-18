import { useEffect, useState } from "react";
import { ProcessAttempt } from "../api/gameService";
import { useCharacter } from "../context/useCharacter";
import "./GamePage.css";

const orderedKeys = [
  "Name",
  "Species",
  "PlaceOrigin",
  "IntroArc",
  "Affiliation",
  "ChakraType",
  "KekkeiGenkai",
  "JutsuAffinity",
  "SpecialAttribute",
];

interface AttemptResult {
  success: boolean;
  message: string;
  [key: string]: string | { Value: string; Status: string } | boolean;
}

const GamePage: React.FC = () => {
  const { character } = useCharacter();
  const [inputName, setInputName] = useState<string>("");
  const [selectedAnime, setSelectedAnime] = useState<string>("");
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
    <div className="page-container">
      <h1>Game Page</h1>

      <div className="input-wrapper">
        <input
          type="text"
          value={inputName}
          onChange={(e) => setInputName(e.target.value)}
          placeholder="Type the character name"
          className="styled-input"
        />
        <button onClick={handleSearch} className="styled-button">
          Search
        </button>
      </div>

      {attemptResult && (
        <div className="result-wrapper">
          <h3>Attempt Result:</h3>
          <div className="result-grid">
            {orderedKeys.map((key) => {
              const entry = attemptResult[key];
              if (
                typeof entry === "object" &&
                entry !== null &&
                "Value" in entry &&
                "Status" in entry
              ) {
                const { Value, Status } = entry;
                return (
                  <div key={key} className={`box ${Status.toLowerCase()}`}>
                    <div className="category-title">{key}</div>
                    <div>{Value}</div>
                    <div className="status-text">({Status})</div>
                  </div>
                );
              }
              return null;
            })}
          </div>
        </div>
      )}
    </div>
  );
};

export default GamePage;
