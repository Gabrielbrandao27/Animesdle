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
