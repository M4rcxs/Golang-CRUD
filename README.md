<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
  </head>
  <body>
    <h1>Golang CRUD</h1>
    <p>The <code>middleware</code> package is responsible for handling requests and responses between a server and a PostgreSQL database.</p>
    <h2>Routes</h2>
    <table>
      <thead>
        <tr>
          <th>Route</th>
          <th>HTTP Method</th>
          <th>Description</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>/users</td>
          <td>GET</td>
          <td>Returns all users from the database.</td>
        </tr>
        <tr>
          <td>/user/{id}</td>
          <td>GET</td>
          <td>Returns a user with the specified ID from the database.</td>
        </tr>
        <tr>
          <td>/user</td>
          <td>POST</td>
          <td>Creates a new user in the database.</td>
        </tr>
        <tr>
          <td>/user/{id}</td>
          <td>PUT</td>
          <td>Updates a user with the specified ID in the database.</td>
        </tr>
        <tr>
          <td>/user/{id}</td>
          <td>DELETE</td>
          <td>Deletes a user with the specified ID from the database.</td>
        </tr>
      </tbody>
    </table>
    <h2>Functions</h2>
    <ul>
      <li><code>createConnection()</code>: establishes a connection to the PostgreSQL database.</li>
      <li><code>CreateUser(w http.ResponseWriter, r *http.Request)</code>: creates a new user in the database.</li>
      <li><code>GetUser(w http.ResponseWriter, r *http.Request)</code>: returns a user with the specified ID from the database.</li>
      <li><code>GetAllUser(w http.ResponseWriter, r *http.Request)</code>: returns all users from the database.</li>
      <li><code>UpdateUser(w http.ResponseWriter, r *http.Request)</code>: updates a user with the specified ID in the database.</li>
      <li><code>DeleteUser(w http.ResponseWriter, r *http.Request)</code>: deletes a user with the specified ID from the database.</li>
      <li><code>insertUser(user models.User) int64</code>: inserts a new user into the database.</li>
      <li><code>getUser(id int64) (models.User, error)</code>: returns a user with the specified ID from the database.</li>
      <li><code>getAllUsers() ([]models.User, error)</code>: returns all users from the database.</li>
      <li><code>updateUser(id int64, user models.User) int64</code>: updates a user with the specified ID in the database.</li>
      <li><code>deleteUser(id int64) int64</code>: deletes a user with the specified ID from the database.</li>
    </ul>
    <h2>Dependencies</h2>
    <ul>
      <li><a href="https://github.com/gorilla/mux">gorilla/mux</a></li>
      <li><a href="github.com/joho/godotenv">joho/godotenv</a></li>
      <li><a href="github.com/lib/pq">lib/pq</a></li>
    </ul>
    </html>
