<p align="center">
  <a href=""><img src="https://github.com/akshanshgusain/january/blob/master/media/logo_dark.png?raw=true" alt="January"></a>
</p>
<p align="center">
    <em>January is Batteries-Included Go Framework inspired by Django. Designed to ease things up for fast development.</em>
</p>
<h2 align="center">
    January Starter Application
</h2>
<h2 align="center">
  ⚡️Quick start: <a href="https://github.com/akshanshgusain/january">January Framework</a>
</h2>
<p align="center">
<a href="">
    <img src="https://img.shields.io/badge/january-docs-blue" alt="January Docs">
</a>
<a href="https://pypi.org/project/fastapi" target="_blank">
    <img src="https://img.shields.io/github/go-mod/go-version/akshanshgusain/january" alt="Supported Go versions">
</a>
</p>


---

## 🌝 Getting Started
The app comes bundled with a Makefile(currently only support macOS). Run the web app running:
```bash
make start


Building January...
January built!
Starting January...
January started!
INFO	2024/09/14 10:53:20 load session called
INFO	2024/09/14 10:53:20 Starting January server at http://127.0.0.1:9095/
INFO	2024/09/14 10:53:20 Quit the server with control+c
```
Visit http://localhost:9095 in your browser to the Home page. You can run the `make stop` command to stop the web server.

```bash
make stop


Stopping January...
Stopped January!
```


## 🏢️ Project Structure

```console
your-repository-name/
│
├── data/
│   └── models.go
│
├── database/
│   └── docker-compsoe.yaml
│
├── handlers/
│   ├── handlers.go
│   └── handlerHelper.go
│
├── middleware/
│   └── middleware.go
│
├── migartions/
│
├── public/
│   ├── ico/
│   └── images/
│
├── views/
│   ├── layouts/
│   └── home.jet
│
├── .gitignore
├── Makefile
├── go.mod
├── init.january.go
├── januaryAppHelper.go
├── main.go
└── README.md
└── routes.go
```