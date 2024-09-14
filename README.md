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

January requires **Go version** `1.18` or **higher** to run. If you need to install or upgrade Go, visit the [official Go download page](https://go.dev/dl/). 

To start setting up your project download the **January-CLI** tool
from here [January-CLI](https://github.com/akshanshgusain/january-cli)
or,

**Homebrew installation**: coming soon!


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


## 👀 Examples

Listed below are some of the common use-cases. If you want to see more code examples, please visit our [Build With repository](https://github.com/gofiber/recipes)

### 📖 **Create APIs**

1. Connect to your **Postgres** database by supplying these values in the `.env` file

    ```.dotenv
    DATABASE_TYPE=postgres
    DATABASE_HOST=localhost
    DATABASE_PORT=5432
    DATABASE_USER=postgres_use
    DATABASE_PASS=postgres_password
    DATABASE_NAME=postgres_db
    DATABASE_SSL_MODE=disable
    ```
2. Write your routes in the `routes.go`:
    ```Go
   func (a *application) routes() *chi.Mux {
        // GET
        a.get("/api/your-get-route", a.Handlers.GetHandler)
        // POST
        a.post("/api/your-post-route", a.Handlers.PostHandler)
        // PUT
        a.put("/api/your-put-route", a.Handlers.PutHandler)
        // PATCH
        a.patch("/api/your-patch-route", a.Handlers.PatchHandler)
        // DELETE
        a.delete("/api/your-delete-route", a.Handlers.DeleteHandler)
   }
    ```
   3. Create a **Handler** and write your business login:
       Create a handler using the **january-cli**: 
       ```bash
      ./january-cli make handler <your-handler-name>
       ```
      The above command will create a handler named: `your-handler-name.go` in the **handlers** folder.
       ```Go
      func (h *Handlers) GetHandler(w http.ResponseWriter, r *http.Request) {
               // Get path parameter
               userID := chi.URLParam(r, "pathParamVar")

               // Get query parameters
               name := r.URL.Query().Get("queryParamVar")
            
               // Business Logic goes here:
                ...
      
                // declare return response struct
               var resp struct {
                       Error   bool   `json:"error"`
                       Message string `json:"message"`
                       Value   string `json:"value"`
                   }
				   
                resp.Error = false
                resp.Message = "Success"
                resp.Value = fromCache.(string)
       
                _ = h.App.WriteJson(w, http.StatusOk, resp)
      } 
       ```