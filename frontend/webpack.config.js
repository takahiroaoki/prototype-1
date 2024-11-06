const path = require('path')
const PugPlugin = require('pug-plugin')
const { CleanWebpackPlugin } = require('clean-webpack-plugin')

const baseDir = __dirname
const srcDir = path.resolve(baseDir, 'src')
const buildDir = path.resolve(baseDir, 'build')

module.exports = {
    context: srcDir,
    output: {
        publicPath: '/',
        path: buildDir,
    },
    plugins: [
        new CleanWebpackPlugin(),
        new PugPlugin({
            pretty: true,
            entry: {
                'templates/main/index': './pages/main/index.pug'
            },
            css: {
                filename: 'static/css/[name].[contenthash:8].css',
            },
        }),
    ],
    resolve: {
        extensions: ['.pug', '.scss'],
        alias: {
            '@src': srcDir,
        }
    },
    module: {
        rules: [
            {
                test: /\.scss$/,
                use: ['css-loader', 'sass-loader'],
            },
        ],
    },
}