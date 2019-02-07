const path = require("path");
const HTMLWebpackPlugin = require("html-webpack-plugin");
const CopyWebpackPlugin = require("copy-webpack-plugin");

module.exports = {
    entry: {
        "index": path.join(__dirname, "frontend/Index.jsx")
    },
    output: {
        path: path.join(__dirname, "dist"),
        filename: "bundle.js"
    },
    module: {
        rules: [{
            test: /\.jsx$/,
            exclude: /node_modules/,
            loader: "babel-loader",
        }, {
            test: /\.json$/,
            use: {
                loader: "json-loader"
            }
        }, {
            test: /\.(png|jpg|gif)$/,
            use: {
                loader: "file-loader"
            }

        }]
    },
    plugins: [
        new HTMLWebpackPlugin({
            template: path.join(__dirname, "frontend/templates/index.html"),
            filename: "index.html",
            inject: "body",
            title: "go serverless thing"
        }),
        new CopyWebpackPlugin([{
            from: "./frontend/images/*.ico",
            flatten: true,
            test: /\.(ico)$/,
            ignore: ["*.jsx"],
            toType: "file"

        }], {
            debug: "debug",
            copyUnmodified: true
        })
    ]
};