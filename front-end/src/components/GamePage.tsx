import { useCharacter } from "../context/useCharacter";

const GamePage: React.FC = () => {
  const { character } = useCharacter();

  return (
    <div>
      <h1>Game Page</h1>
      {character ? (
        <pre>{JSON.stringify(character, null, 2)}</pre>
      ) : (
        <p>Carregando personagem...</p>
      )}
    </div>
  );
};
export default GamePage;