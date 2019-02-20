// webpack.config.js

const path = require('path');
const webpack = require('webpack');

module.exports = {
    entry: './assets/js/index.js',
    output: {
        path: path.resolve(__dirname, 'public/js'),
        publicPath: '/static/js',
        filename: 'main.js'
    },
    devtool: 'inline-source-map',
    module: {
        rules: [
            {
                test: /\.tag$/,
                exclude: /node_modules/,
                loader: 'riot-tag-loader',
                query: {
                    type: 'es6', // transpile the riot tags using babel
                    hot: true
                }
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader: 'babel-loader'
            }
        ]
    },
    devServer: {
        proxy: {
            '**': {
                target: 'http://localhost:8044',
            }
        }
    }
}