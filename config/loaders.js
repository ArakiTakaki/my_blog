const isSourceMaps = {
  development: true,
  production: false
};
console.log(process.env.MODE);
const isSourceMap = isSourceMaps[process.env.MODE] || false;

module.exports = [
  {
    test: /\.(js|jsx)$/,
    exclude: /node_modules/,
    use: ["babel-loader"]
  },
  {
    enforce: "pre",
    test: /\.(js|jsx)$/,
    exclude: /node_modules/,
    loader: "eslint-loader"
  },
  {
    test: /\.pug$/,
    use: [{ loader: "pug-loader" }]
  },
  {
    test: /\.(sass|scss)$/,
    use: [
      { loader: "style-loader" },
      {
        loader: "css-loader",
        options: {
          sourceMap: isSourceMap,
          importLoader: 2,
          modules: true,
          localIdentName: "[path]__[name]__[hash:base64:5]"
        }
      },
      {
        loader: "postcss-loader",
        options: {
          sourceMap: isSourceMap,
          plugins: [require("autoprefixer")({ grid: true })]
        }
      },
      {
        loader: "sass-loader",
        options: {
          sourceMap: isSourceMap
        }
      }
    ]
  },
  {
    test: /\.svg$/,
    use: [
      { loader: "babel-loader" },
      {
        loader: "react-svg-loader",
        options: {
          jsx: true
        }
      }
    ]
  }
];
