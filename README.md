<p align="center">
  <a href=""><img src="https://github.com/akshanshgusain/january/blob/master/media/logo_dark.png?raw=true" alt="January"></a>
</p>
<p align="center">
    <em>January is Batteries-Included Go Framework inspired by Django. Designed to ease things up for fast development.</em>
</p>

<p align="center">
<a href="https://github.com/fastapi/fastapi/actions?query=workflow%3ATest+event%3Apush+branch%3Amaster" target="_blank">
    <img src="https://github.com/fastapi/fastapi/workflows/Test/badge.svg?event=push&branch=master" alt="Test">
</a>
<a href="https://coverage-badge.samuelcolvin.workers.dev/redirect/fastapi/fastapi" target="_blank">
    <img src="https://coverage-badge.samuelcolvin.workers.dev/fastapi/fastapi.svg" alt="Coverage">
</a>
<a href="">
    <img src="https://img.shields.io/badge/january-docs-blue" alt="January Docs">
</a>
<a href="https://pypi.org/project/fastapi" target="_blank">
    <img src="https://img.shields.io/github/go-mod/go-version/akshanshgusain/january" alt="Supported Go versions">
</a>
</p>

---

## 🎯 Features

* Robust routing
* Serve static files
* Easy access to multiple databases
* Complete user authentication
* Database migrations
* Session support
* Template engines
* Generate handlers
* Middleware support
* Form Validation
* CSRF protection
* Encryption
* Multiple Caching backends

---

## ⚙️ Installation

January requires **Go version** `1.18` or **higher** to run. To start setting up your project download the **January-CLI** tool
from here [January-CLI](https://github.com/akshanshgusain/january-cli)


## ⚡️ Quickstart

Here is  a basic example to create a simple web app with **January-CLI**:
```bash
./january-cli new github.com/your_username/your-repository_name
```
This command creates a new project folder named **your-repository-name**.
After creating the project move the CLI to the project folder :

```bash
mv ./january-cli ./your-repository-name
```
and cd into the project folder:
```bash
cd your-repository-name/
```
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