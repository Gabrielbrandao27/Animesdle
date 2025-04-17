import { createContext, ReactNode, useEffect, useState } from "react";
import { NarutoCharacter } from "../types/Characters";

export type Character = NarutoCharacter;

interface CharacterContextType {
  character: Character | null;
  setCharacter: (char: Character) => void;
}

const CharacterContext = createContext<CharacterContextType | undefined>(undefined);

export const CharacterProvider = ({ children }: { children: ReactNode }) => {
  const [character, setCharacter] = useState<Character | null>(() => {
    const storedCharacter = localStorage.getItem("character");
    return storedCharacter ? JSON.parse(storedCharacter) : null;
  });

  useEffect(() => {
    if (character) {
      localStorage.setItem("character", JSON.stringify(character));
    } else {
      localStorage.removeItem("character");
    }
  }, [character]);

  return (
    <CharacterContext.Provider value={{ character, setCharacter }}>
      {children}
    </CharacterContext.Provider>
  );
};

export default CharacterContext;
