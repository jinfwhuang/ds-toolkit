console.error(Object.keys(require('chalk')));
console.error(Object.keys(require('colors')));
console.error(Object.keys(require('color')));

console.log(require.resolve('color'));
console.log(module);

const { increment } = require('@exphz/hello-lib');
console.log(increment(5));

// sleep and allow exec into the container to debug
console.log("sleeping...");
setTimeout(() => console.log("waiting a long time"), 300000);

