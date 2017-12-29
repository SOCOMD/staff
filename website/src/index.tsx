import { h, render } from 'preact';
//import HelloWorld from "./helloworld"
require('file-loader?name=[name].[ext]!../index.html');
//render(<HelloWorld name="World" />, document.querySelector('#app'));

import App from './components/app'
render(<App />, document.querySelector('#app'))