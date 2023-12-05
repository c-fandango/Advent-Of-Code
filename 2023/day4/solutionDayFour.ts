// Not proud of this one

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

function solutionOne(cards: string[]) {
  return cards.reduce(function (scoreTotal: number, card: string) {
    const gameSplt: string = card.split(":")[1];
    const [winCardStr, ourCardStr] = gameSplt
      .split("|")
      .map((x) => x.split(/\s+/).filter((y) => y));

    const winCard = winCardStr.map((x) => parseInt(x));
    const ourCard = ourCardStr.map((x) => parseInt(x));

    let score: number = winCard.reduce(function (score: number, num: number) {
      return ourCard.includes(num) ? score * 2 : score;
    }, 0.5);

    return score == 0.5 ? scoreTotal : score + scoreTotal;
  }, 0);
}

function solutionTwo(cards: string[]) {
  const cardCounts: number[] = new Array(cards.length).fill(1);

  cards.forEach(function (card: string, id: number) {
    const gameSplt: string = card.split(":")[1];
    const [winCardStr, ourCardStr] = gameSplt
      .split("|")
      .map((x) => x.split(/\s+/).filter((y) => y));

    const winCard = winCardStr.map((x) => parseInt(x));
    const ourCard = ourCardStr.map((x) => parseInt(x));

    let score: number = winCard.reduce(function (score: number, num: number) {
      return ourCard.includes(num) ? score + 1 : score;
    }, 0);

    for (let i = id + 1; i < id + score + 1 && i < cards.length; i++) {
      cardCounts[i] += cardCounts[id];
    }
  });

  return cardCounts.reduce((sum, x) => sum + x, 0);
}

function runSolutions(data: string[]) {
  const resultOne: number = solutionOne(data);
  console.log("result", resultOne);
  const resultTwo: number = solutionTwo(data);
  console.log("result", resultTwo);
}

getData("data.txt", runSolutions);
