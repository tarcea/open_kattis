const readline = require('readline');
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

rl.on('line', (line) => {
  const hours = line
    .slice(0, 2)
    .split('')
    .map((el) => {
      let digits = Number(el).toString(2);
      while (digits.length < 4) {
        digits = '0' + digits;
      }
      return digits;
    });
  const minutes = line
    .slice(2)
    .split('')
    .map((el) => {
      let digits = Number(el).toString(2);
      while (digits.length < 4) {
        digits = '0' + digits;
      }
      return digits;
    });
  const r = [...hours, ...minutes];
  let one = [];
  let two = [];
  let three = [];
  let four = [];
  r.forEach((el, i) => {
    one.push(el[0]);
    two.push(el[1]);
    three.push(el[2]);
    four.push(el[3]);
  });
  one.splice(2, 0, ' ');
  two.splice(2, 0, ' ');
  three.splice(2, 0, ' ');
  four.splice(2, 0, ' ');

  console.log(one.join(' ').replace(/0/gi, '.').replace(/1/gi, '*'));
  console.log(two.join(' ').replace(/0/gi, '.').replace(/1/gi, '*'));
  console.log(three.join(' ').replace(/0/gi, '.').replace(/1/gi, '*'));
  console.log(four.join(' ').replace(/0/gi, '.').replace(/1/gi, '*'));
  rl.close();
});
