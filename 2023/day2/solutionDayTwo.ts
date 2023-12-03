const fs = require("fs");

type solutionFunc = (_: string[]) => void;

function getData(filePath: string, cb: solutionFunc) {
  fs.readFile(filePath, function (err: string, data: any) {
    if (!err) {
      let lines: string[];
      lines = data
        .toString()
        .split("\n")
        .filter((n: any) => n);

      return cb(lines);
    } else {
      console.log("error");
      throw err;
    }
  });
}

function solutionOne(gameLines: string[]) {
  const redMax: number = 12;
  const greenMax: number = 13;
  const blueMax: number = 14;

  let gameSum: number = 0;

  for (const gameLine of gameLines) {
    let [gameId, game]: string[] = gameLine.split(":");
    let turns: string[] = game.split(";");
    let possible: boolean = true;

    for (const turn of turns) {
      let redMatch: RegExpMatchArray | null = turn.match(/([0-9]+) red/);
      let greenMatch: RegExpMatchArray | null = turn.match(/([0-9]+) green/);
      let blueMatch: RegExpMatchArray | null = turn.match(/([0-9]+) blue/);

      if (redMatch && parseInt(redMatch[1]) > redMax) {
        possible = false;
        break;
      } else if (greenMatch && parseInt(greenMatch[1]) > greenMax) {
        possible = false;
        break;
      } else if (blueMatch && parseInt(blueMatch[1]) > blueMax) {
        possible = false;
        break;
      }
    }

    if (possible) {
      let id: RegExpMatchArray | null = gameId.match(/[0-9]+/);
      if (id) {
        gameSum += parseInt(id[0]);
      }
    }
  }
  console.log("result", gameSum);
}

function solutionTwo(gameLines: string[]) {
  let gameSum: number = 0;

  for (const gameLine of gameLines) {
    let [gameId, game]: string[] = gameLine.split(":");
    let turns: string[] = game.split(";");
    let redMax: number = 0;
    let greenMax: number = 0;
    let blueMax: number = 0;

    for (const turn of turns) {
      let redMatch: RegExpMatchArray | null = turn.match(/([0-9]+) red/);
      let greenMatch: RegExpMatchArray | null = turn.match(/([0-9]+) green/);
      let blueMatch: RegExpMatchArray | null = turn.match(/([0-9]+) blue/);

      if (redMatch && parseInt(redMatch[1]) > redMax) {
        redMax = parseInt(redMatch[1]);
      }
      if (greenMatch && parseInt(greenMatch[1]) > greenMax) {
        greenMax = parseInt(greenMatch[1]);
      }
      if (blueMatch && parseInt(blueMatch[1]) > blueMax) {
        blueMax = parseInt(blueMatch[1]);
      }
    }

    gameSum += redMax * greenMax * blueMax;
  }
  console.log("result", gameSum);
}

function runSolutions(data: string[]) {
  solutionOne(data);
  solutionTwo(data);
}

getData("data.txt", runSolutions);
