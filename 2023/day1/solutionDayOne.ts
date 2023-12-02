const fs = require("fs");

type filterFunc = (_: string[]) => string[];
type solutionFunc = (_: string[]) => void;
interface StringArray {
  [index: string]: string;
}

function getData(filePath: string, cb: solutionFunc) {
  fs.readFile(filePath, function (err: string, data: any) {
    if (!err) {
      let lines: string[];
      lines = data.toString().split("\n");
      lines = lines.filter((n) => n);
      return cb(lines);
    } else {
      console.log("error");
      throw err;
    }
  });
}

function filterNums(input: string[]) {
  let numChars: StringArray = {
    one: "1",
    two: "2",
    three: "3",
    four: "4",
    five: "5",
    six: "6",
    seven: "7",
    eight: "8",
    nine: "9",
  };

  let output: string[];

  output = input.map(function (line) {
    for (const letter in numChars) {
      line = line.replaceAll(letter, letter + numChars[letter] + letter);
    }
    return line;
  });
  return output;
}

function runSolutions(data: string[]) {
  solution(data);
  solution(data, filterNums);
}

function solution(calData: string[], filter?: filterFunc) {
  const digits = "123456789";
  let result = 0;
  if (filter) {
    calData = filter(calData);
  }

  for (const line of calData) {
    let values = "";

    for (const char of line) {
      if (digits.includes(char)) {
        values += char;
      }
    }
    let numStr = values[0] + values[values.length - 1];
    result += parseInt(numStr);
  }
  console.log("result", result);
}

getData("data.txt", runSolutions);
