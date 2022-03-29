function increment(n) {
  console.log("hello-lib function call increment");
  console.log(require.resolve('smallest'));
  return n + 1;
}

exports.increment = increment;
