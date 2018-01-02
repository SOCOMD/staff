var path = require( 'path' );
var ExtractTextPlugin = require( "extract-text-webpack-plugin" );
var copyWebpackPlugin = require( "copy-webpack-plugin" );
var webpack = require("webpack");
module.exports = {
	devtool: 'source-map',
	entry: [ './src/index' ],
	output: {
		path: path.resolve( './dist' ),
		filename: 'app.js'
	},
	resolve: {
		extensions: [ '.js', '.ts', '.tsx' ],
	},
	module: {
		loaders: [ {
				test: /\.tsx?$/,
				exclude: /node_modules/,
				loaders: [ 'ts-loader' ]
			},
			{
				test: /\.js$/,
				exclude: /node_modules/,
				loaders: [ 'babel-loader' ]
			},
			{
				test: /\.css$/,
				loaders: [ 'style-loader', 'css-loader' ]

			}
		]
	},
	plugins: [
		new copyWebpackPlugin( [ { from: 'src/assets/' } ] ),
		new webpack.EnvironmentPlugin({
			webgrpc_host:"http://192.168.0.15:8081"
		})
	]
};