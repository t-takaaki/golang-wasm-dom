const express = require('express');
express.static.mime.define({'application/wasm': ['wasm']});

const app = express();
app.use(express.static('public'));
app.listen(3000, () => console.log('Listening on port 3000...'));
