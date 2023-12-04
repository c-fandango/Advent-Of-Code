// Not proud of this one

const fs = require("fs");

type solutionFunc = (_: string[]) => void;

interface StringArray {
  [index: string]: number[];
}

function padGrid(grid: string[], char: string) {
  let newGrid: string[] = [];
  const rowPad: string = char.repeat(grid[0].length + 2);
  newGrid.push(rowPad);

  for (const row of grid) {
    const newRow: string = char + row + char;
    newGrid.push(newRow);
  }
  newGrid.push(rowPad);
  return newGrid;
}

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

function isSymbol(char: string) {
  const spaceChar: string = ".";
  const digits: string = "1234567890";
  return !digits.includes(char) && !spaceChar.includes(char);
}

function hash(num1: number, num2: number) {
  return String(num1) + "x" + String(num2);
}

function solutionOne(grid: string[]) {
  const spaceChar: string = ".";
  const digits: string = "1234567890";

  let resultSum: number = 0;

  const paddedGrid: string[] = padGrid(grid, ".");

  for (let i = 1; i < paddedGrid.length - 1; i++) {
    let thisNum: string = "";
    let isValid: boolean = false;

    for (let j = 1; j < paddedGrid[0].length - 1; j++) {
      const thisChar: string = paddedGrid[i][j];

      if (digits.includes(thisChar)) {
        thisNum += thisChar;

        if (isSymbol(paddedGrid[i - 1][j - 1])) {
          isValid = true;
        } else if (isSymbol(paddedGrid[i - 1][j])) {
          isValid = true;
        } else if (isSymbol(paddedGrid[i - 1][j + 1])) {
          isValid = true;
        } else if (isSymbol(paddedGrid[i + 1][j - 1])) {
          isValid = true;
        } else if (isSymbol(paddedGrid[i + 1][j])) {
          isValid = true;
        } else if (isSymbol(paddedGrid[i + 1][j + 1])) {
          isValid = true;
        } else if (isSymbol(paddedGrid[i][j - 1])) {
          isValid = true;
        }

        if (!digits.includes(paddedGrid[i][j + 1])) {
          if (!spaceChar.includes(paddedGrid[i][j + 1])) {
            isValid = true;
          }

          if (isValid) {
            resultSum += parseInt(thisNum);
          }
          isValid = false;
          thisNum = "";
        }
      }
    }
  }
  console.log("result", resultSum);
}

function solutionTwo(grid: string[]) {
  const spaceChar: string = ".";
  const digits: string = "1234567890";

  let stars: StringArray = {};

  let resultSum: number = 0;

  const paddedGrid: string[] = padGrid(grid, ".");

  for (let i = 1; i < paddedGrid.length - 1; i++) {
    let thisNum: string = "";
    let potentialGears: string[] = [];

    for (let j = 1; j < paddedGrid[0].length - 1; j++) {
      const thisChar: string = paddedGrid[i][j];

      if (digits.includes(thisChar)) {
        thisNum += thisChar;

        if (paddedGrid[i - 1][j - 1] == "*") {
          potentialGears.push(hash(i - 1, j - 1));
        }
        if (paddedGrid[i - 1][j] == "*") {
          potentialGears.push(hash(i - 1, j));
        }
        if (paddedGrid[i - 1][j + 1] == "*") {
          potentialGears.push(hash(i - 1, j + 1));
        }
        if (paddedGrid[i + 1][j - 1] == "*") {
          potentialGears.push(hash(i + 1, j - 1));
        }
        if (paddedGrid[i + 1][j] == "*") {
          potentialGears.push(hash(i + 1, j));
        }
        if (paddedGrid[i + 1][j + 1] == "*") {
          potentialGears.push(hash(i + 1, j + 1));
        }
        if (paddedGrid[i][j + 1] == "*") {
          potentialGears.push(hash(i, j + 1));
        }
        if (paddedGrid[i][j - 1] == "*") {
          potentialGears.push(hash(i, j - 1));
        }

        if (!digits.includes(paddedGrid[i][j + 1])) {
          let uniqGears = [...new Set(potentialGears)];
          for (const star of uniqGears) {
            if (!(star in stars)) {
              stars[star] = [];
            }
            stars[star].push(parseInt(thisNum));
          }
          thisNum = "";
          potentialGears = [];
        }
      }
    }
  }
  for (const star in stars) {
    if (stars[star].length == 2) {
      resultSum += stars[star][0] * stars[star][1];
    }
  }
  console.log("result", resultSum);
}

function runSolutions(data: string[]) {
  solutionOne(data);
  solutionTwo(data);
}

getData("data.txt", runSolutions);
