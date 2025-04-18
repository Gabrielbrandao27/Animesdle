# Animesdle 🥷🍜
Go Learning Project!

Based on a webpage game that I like to play with my friends.

Web Application similar to a "Wordle", but for Animes!

# Steps for running the Game:

## Prerequisites:
* 💻 Frontend (React + Vite)
    - Node.js v18 or higher
    - npm or yarn – for package managing..

* 🔙 Backend (Golang)
    - Go v1.20 or higher

* 🗃️ Docker
    - https://docs.docker.com/get-started/get-docker/

## Tutorial for running locally:
1. `docker pull gabrielbrandao2711/animesdle-mysql` pulls the imagem with the character tables and data
2. `docker run -p 3306:3306 gabrielbrandao2711/animesdle-mysql` runs the image locally on port 3306
3. `cd back-end/cmd`
4. `go run main.go` runs the back-end
5. `cd ../front-end`
6. `npm run dev` runs the front-end
