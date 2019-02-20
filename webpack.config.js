// webpack.config.js

const path = require('path');
const webpack = require('webpack');

module.exports = {
    entry: [
        './assets/js/index.js',
        './assets/sass/main.scss',
    ],
    output: {
        path: path.resolve(__dirname, 'public'),
        publicPath: '/static',
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
            },
            {
                test: /\.scss$/,
                use: [
                    {
                        loader: 'file-loader',
                        options: {
                            name: '[name].css',
                        }
                    },
                    {
                        loader: 'extract-loader'
                    },
                    {
                        loader: 'css-loader'
                    },
                    {
                        loader: 'sass-loader'
                    },
                ]
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