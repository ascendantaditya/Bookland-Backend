Here’s a corrected version of your **README.md** file, addressing any issues and ensuring that the image URLs are properly linked:

---

# Bookland-Backend

![Bookland Logo](https://github.com/ascendantaditya/Bookland-Backend/blob/main/Image/ok.jpg)

This is the backend for **Bookland**, a bookstore management system. The backend is built using **Go** with the **Gorm** ORM for database handling and **Gorilla Mux** for routing.

The **Bookland-Backend** uses a RESTful API structure, connecting a MySQL database with the frontend via the following components:
- **Go** for API logic.
- **Gorilla Mux** for routing.
- **Gorm** for managing database interactions.

## Installation

To install the necessary dependencies, run the following commands:

1. Install Gorm (Go ORM library):
   ```bash
   go get "github.com/jinzhu/gorm"
   ```

2. Install the MySQL dialect for Gorm:
   ```bash
   go get "github.com/jinzhu/gorm/dialects/mysql"
   ```

3. Install Gorilla Mux for routing:
   ```bash
   go get "github.com/gorilla/mux"
   ```

## Usage

### Clone the repository

![Clone Repo](https://github.com/ascendantaditya/Bookland-Backend/blob/main/Image/clone.png)

1. Clone the repository:
   ```bash
   git clone https://github.com/ascendantaditya/Bookland-Backend.git
   ```

2. Navigate to the project directory:
   ```bash
   cd Bookland-Backend
   ```

3. Ensure the Go modules are initialized (if not already):
   ```bash
   go mod init Bookland-Backend
   ```

4. Install the required dependencies:
   ```bash
   go mod tidy
   ```

5. Run the application:
   ```bash
   go run main.go
   ```

## Libraries

The following Go libraries are used in this project:

1. **Gorm**: An ORM (Object-Relational Mapping) library for Go, which simplifies database handling.
   - [Gorm Documentation](https://gorm.io/docs/)

2. **MySQL Dialect for Gorm**: Allows Gorm to work with MySQL databases.
   - [MySQL Dialect for Gorm Documentation](https://gorm.io/docs/connecting_to_the_database.html#MySQL)

3. **Gorilla Mux**: A powerful HTTP router and URL matcher for building web applications.
   - [Gorilla Mux Documentation](https://github.com/gorilla/mux)

## License

![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Fixes:
1. Corrected image links to ensure they point to the correct GitHub paths.
2. Maintained consistency in file structure and instructions. 

This version should now work correctly in your repository.
