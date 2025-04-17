import { createContext, ReactNode, useState } from "react";
import { NarutoCharacter } from "../types/Characters";

export type Character = NarutoCharacter; // Você pode substituir por NarutoCharacter ou OnePieceCharacter

interface CharacterContextType {
  character: Character | null;
  setCharacter: (char: Character) => void;
}

const CharacterContext = createContext<CharacterContextType | undefined>(undefined);

export const CharacterProvider = ({ children }: { children: ReactNode }) => {
  const [character, setCharacter] = useState<Character | null>(null);

  return (
    <CharacterContext.Provider value={{ character, setCharacter }}>
      {children}
    </CharacterContext.Provider>
  );
};

export default CharacterContext;
