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


## 💡 Inspiration
When teams at my organization began transitioning from Python, Node.js, and PHP to Go, we encountered a new challenge: as each developer utilized different libraries, standards, and frameworks, interoperability between teams became increasingly problematic. The January Web Framework was conceived with a "batteries-included" approach, enabling teams to swiftly integrate into the Go ecosystem while adhering to consistent standards across the organization.


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
This command creates a new project directory named **your-repository-name**.
After creating the project move the CLI to the project directory :

```bash
mv ./january-cli ./your-repository-name
```
and cd into the project directory:
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


## 👀 Example Usage

Listed below are some of the common use-cases. If you want to see more code examples, please visit our [Build With January repository](https://github.com/gofiber/recipes)

### 📖 **Generate Migration Files**

Create migration file using the **january-cli**. The following command will create two migration files: **<migration-name>.up.sql** and **<migration-name>.down.sql** in the 
```bash
./january-cli make migration <migartion-name>
```

### 📖 **Run Migrations**

Run migrations with the **january-cli**. The following command will run the migration:
1. Migrate Up: Run the latest up migration:
```bash
./january-cli migrate up
```
2. Migrate Down: Run the latest down migration:
```bash
./january-cli migrate down
```
3. Migrate Rest: Run all the down migrations and then run all the up migrations:
```bash
./january-cli migrate reset
```

### 📖 **Generate Model**

Create a model with **january-cli**. The following command will create a `model-name.go` file in the data directory: 
```bash
./january-cli make model <model-name>
```
By default the Models are generate with [Upper](https://upper.io/v4/) DAL to access the Database. After generating the `modelname.go` it needs to me added to the **Model struct** in the `data/models.go`.
```Go
var db *sql.DB
var upper db2.Session

type Models struct {
	// any models inserted here (and in the new function)
	// are easily accessible throughout the entire application
	ModelNames  ModelName   // <---- your Model(s)
}

func New(databasePool *sql.DB) Models {
	db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":

		upper, _ = mysql.New(databasePool)
	case "postgres", "postgresql":

		upper, _ = postgresql.New(databasePool)
	default:
		// do nothing
	}

	return Models{
        ModelNames:  ModelName{},
	}
}
```
Now, your models are ready to use. The Models come with pre-build CRUD methods:
```Go
// ModelName struct
type ModelName struct {
    ID        int       `db:"id,omitempty"`
    CreatedAt time.Time `db:"created_at"`
    UpdatedAt time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *ModelName) Table() string {
    return "modelnames"
}

// GetAll gets all records from the database, using upper
func (t *ModelName) GetAll(condition up.Cond) ([]*ModelName, error) {
    collection := upper.Collection(t.Table())
    var all []*ModelName

    res := collection.Find(condition)
    err := res.All(&all)
    if err != nil {
        return nil, err
    }

    return all, err
}

// Get gets one record from the database, by id, using upper
func (t *ModelName) Get(id int) (*ModelName, error) {
    var one ModelName
    collection := upper.Collection(t.Table())

    res := collection.Find(up.Cond{"id": id})
    err := res.One(&one)
    if err != nil {
        return nil, err
    }
    return &one, nil
}

// Update updates a record in the database, using upper
func (t *ModelName) Update(m ModelName) error {
    m.UpdatedAt = time.Now()
    collection := upper.Collection(t.Table())
    res := collection.Find(m.ID)
    err := res.Update(&m)
    if err != nil {
        return err
    }
    return nil
}

// Delete deletes a record from the database by id, using upper
func (t *ModelName) Delete(id int) error {
    collection := upper.Collection(t.Table())
    res := collection.Find(id)
    err := res.Delete()
    if err != nil {
        return err
    }
    return nil
}

// Insert inserts a model into the database, using upper
func (t *ModelName) Insert(m ModelName) (int, error) {
    m.CreatedAt = time.Now()
    m.UpdatedAt = time.Now()
    collection := upper.Collection(t.Table())
    res, err := collection.Insert(m)
    if err != nil {
        return 0, err
    }

    id := getInsertID(res.ID())

    return id, nil
}

// Builder is an example of using upper's sql builder
func (t *ModelName) Builder(id int) ([]*ModelName, error) {
    collection := upper.Collection(t.Table())

    var result []*ModelName

    err := collection.Session().
        SQL().
        SelectFrom(t.Table()).
        Where("id > ?", id).
        OrderBy("id").
        All(&result)
    if err != nil {
        return nil, err
    }
    return result, nil
}
```


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
   The above command will create a handler named: `your-handler-name.go` in the **handlers** directory.
    ```Go
   func (h *Handlers) GetHandler(w http.ResponseWriter, r *http.Request) {
            // Get path parameter
            pathParamVar := chi.URLParam(r, "pathParamVar")

            // Get query parameters
            queryParamVar := r.URL.Query().Get("queryParamVar")
            
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
   
### 📆 Broken Features
1. I will be removing the Jet Templating support, the standard library support for templating is decent.

### 📆 Upcoming Features
1. MariaDB support
2. Live-reloads 
3. Automatic Swagger Documentation
4. File system support
5. Support for Websockets, GraphQL and gRPC