import { useEffect, useState } from "react";
import { startGame } from "../api/gameService";
import { NarutoCharacter } from "../types/Characters";

const GamePage: React.FC = () => {
    const [character, setCharacter] = useState<NarutoCharacter | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const data = await startGame("Naruto");
                console.log("Personagem recebido:", data);
                setCharacter(data);
            } catch (error) {
                console.error("Error fetching data:", error);
            }
        };
        
        fetchData();
    }, []);

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