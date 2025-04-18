import axiosInstance from "./axiosConfig.ts";

export const startGame = async (anime: string) => {
  try {
    const response = await axiosInstance.post(
      `/start-game?anime=${encodeURIComponent(anime)}`
    );
    return response.data;
  } catch (error) {
    console.error("Error starting game:", error);
    throw error;
  }
};

export const ProcessAttempt = async <T>(data: {
  name: string;
  anime: string;
  currentCharacter: T;
}) => {
  try {
    const payload = {
      name: data.name,
      anime: data.anime,
      currentCharacter: JSON.stringify(data.currentCharacter),
    };
    const response = await axiosInstance.post("/attempt", payload, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    console.error("Error processing attempt:", error);
    throw error;
  }
};
