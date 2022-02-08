const webpack = require('webpack');
const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

const dev = process.env.APP_ENV !== 'production';

module.exports = {
  mode: dev ? 'development' : 'production',
  devtool: 'source-map',
  entry: {
    app: './src/index',
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
  },
  optimization: {
    noEmitOnErrors: true,
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env': { APP_ENV: JSON.stringify(process.env.APP_ENV) },
    }),
    new MiniCssExtractPlugin({ filename: 'styles.css' }),
    new HtmlWebpackPlugin({
      title: 'ProgrammingCourseMarket',
    }),
  ],
  resolve: {
    modules: [path.resolve(__dirname, 'node_modules')],
    extensions: ['.js', '.ts', '.tsx'],
  },
  module: {
    rules: [
      {
        test: [/\.tsx$/, /\.ts$/],
        exclude: [
          /node_modules/,
        ],
        use: [
          { loader: 'babel-loader' },
          {
            loader: '@linaria/webpack-loader',
            options: { sourceMap: dev },
          },
          'ts-loader',
        ],
      },
      {
        test: /\.css$/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
          },
          {
            loader: 'css-loader',
            options: { sourceMap: dev },
          },
        ],
      },
      {
        test: /\.(jpg|png|gif|woff|woff2|eot|ttf|svg)$/,
        use: [{ loader: 'file-loader' }],
      },
    ],
  },
  devServer: {
    port: 3000,
    static: {
      directory: path.resolve(__dirname, 'static'),
    },
    historyApiFallback: true,
  },
};
