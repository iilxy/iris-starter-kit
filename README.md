# iris-starter-kit

This project contains a quick starter kit for **Facebook React** Single Page Apps with **Golang** server side render via goja javascript engine, implemented in pure Golang and also with a set of useful features for rapid development of efficient applications.

## What it contains?

* server side render via [iris](https://github.com/kataras/iris)
* api requests between your react application and server side application directly  via [fetch polyfill](https://github.com/olebedev/gojax/tree/master/fetch)
* title, Open Graph and other domain-specific meta tags render for each page at the server and at the client
* server side redirect
* embedding static files into artefact via bindata
* highest performance [iris](https://github.com/kataras/iris) framework
* advanced cli via [cli](https://github.com/urfave/cli)
* routing via [react-router](https://github.com/reactjs/react-router)
* ES6 & JSX via [babel-loader](https://github.com/babel/babel-loader) with minimal runtime dependency footprint
* [redux](https://rackt.org/redux/) as state container
* [redux-devtools](https://github.com/gaearon/redux-devtools)
* css styles without global namespace via PostCSS, [css-loader](https://github.com/webpack/css-loader) & css-modules
* separate css file to avoid FOUC
* hot reloading via [react-transform](https://github.com/gaearon/babel-plugin-react-transform) & [HMR](http://webpack.github.io/docs/hot-module-replacement.html)
* webpack bundle builder
* eslint and golint rules for Makefile

## Workflow dependencies

* [golang](https://golang.org/)
* [node.js](https://nodejs.org/) with [yarn](https://yarnpkg.com)

Note that works even on Windows.

## Project structure

### The server's entry point

```bash
$ tree server
server
├── api.go
├── app.go
├── bindata.go <-- this file is gitignored, it will appear at compile time
├── conf.go
├── data
│   └── templates
│       └── react.html
├── main.go <-- main function declared here
├── react.go
└── utils.go
```

The `./server/` is flat golang package.

### The client's entry point

It's simple React application

```bash
$ tree client
client
├── actions.js
├── components
│   ├── app
│   │   ├── favicon.ico
│   │   ├── index.js
│   │   └── styles.css
│   ├── homepage
│   │   ├── index.js
│   │   └── styles.css
│   ├── not-found
│   │   ├── index.js
│   │   └── styles.css
│   └── usage
│       ├── index.js
│       └── styles.css
├── css
│   ├── funcs.js
│   ├── global.css
│   ├── index.js
│   └── vars.js
├── index.js <-- main function declared here
├── reducers.js
├── router
│   ├── index.js
│   ├── routes.js
│   └── toString.js
└── store.js
```

The client app will be compiled into `server/data/static/build/`.  Then it will be embedded into go package via _go-bindata_. After that the package will be compiled into binary.

**Convention**: javascript app should declare [_main_](https://github.com/kataras/iris-starter-kit/blob/master/client/index.js#L4) function right in the global namespace. It will used to render the app at the server side.

## Install

Clone the repo:

```bash
$ git clone https://github.com/kataras/iris-starter-kit.git $GOPATH/src/github.com/<username>/<project>
$ cd $GOPATH/src/github.com/<username>/<project>
```

Install dependencies:

```bash
$ npm install
$ webpack
$ go get github.com/kataras/iris/v12@v12.2.0-alpha
$ go get -u github.com/go-bindata/go-bindata/...
```

## Run development

Start dev server:

```bash
$ cd server
$ go-bindata -fs ./data/...
$ go build
$ server run # ./server.exe run
```

that's it. Open [http://localhost:5001/](http://localhost:5001/)(if you use default port) at your browser. Now you ready to start coding your awesome project.

## Build

Install dependencies and type `NODE_ENV=production make build`. This rule is producing webpack build and regular golang build after that. Result you can find at `$GOPATH/bin`. Note that the binary will be named **as the current project directory**.

## License

As https://github.com/olebedev/go-starter-kit says; MIT
