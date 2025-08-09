# Expense Tracker API

## 📝 Project Description

This is an API for an expense tracker application. This API allows users to create, read, update, and delete expenses. Users are able to sign up and log in to the application. Each user have their own set of expenses.

## ✨ Features

Here are the features that are implemented in the Expense Tracker API:

- Sign up as a new user.
- Generate and validate JWTs for handling authentication and user session.
- List and filter past expenses using the following filters:
  - Past week
  - Past month
  - Last 3 months
  - Custom (to specify a start and end date of your choosing).
- Add a new expense
- Remove existing expenses
- Update existing expenses

**Constraints**

You can use any programming language and framework of your choice. You can use a database of your choice to store the data. You can use any ORM or database library to interact with the database.

Here are some constraints that you should follow:

You’ll be using JWT (JSON Web Token) to protect the endpoints and to identify the requester.
For the different expense categories, you can use the following list (feel free to decide how to implement this as part of your data model):

- Groceries
- Leisure
- Electronics
- Utilities
- Clothing
- Health
- Others

## 🛠️ Core Technologies Used

- Go (Golang): The primary programming language.

- Standard Library Packages:

  - fmt: For formatted input and output.

  - os: For operating system interactions, including file system operations.

  - strings: For string manipulation (e.g., trimming, splitting, case conversion).

  - encoding/json: For marshaling (saving) and unmarshaling (loading) Go structs to/from JSON.

## 🚀 Installation

To get expense tracker up and running on your local machine, follow these steps:

1. Ensure Go is Installed:
   Make sure you have Go installed (version 1.18 or higher is recommended). You can download it from go.dev/dl/.
   Verify your installation:

```
go version
```

2. Clone the Repository (or create project manually):
   If you're starting from scratch as part of a learning exercise, you'd create the project structure manually as described in the task instructions. If this were a real repository:

```
git clone https://github.com/philipoyelegbin/expense-tracker
cd expense-tracker
```

3. Initialize Go Module (if not already done):

```
go mod init github.com/philipoyelegbin/expense-tracker.git      # Only if you created the project manually
```

4. Build the Executable:
   This command compiles your Go source code into a single executable binary.

```
go build -o expense-tracker
```

This will create an executable file named expense-tracker in your project's root directory.

## 💡 Usage

Once built, you can run the CLI commands from your terminal.

**General Usage**

```
./expense-tracker     # Prompt you interactively to select an action to perform
```

## 📂 Project Structure

```
expense-tracker/
    ├── main.go # Main entry point and CLI command handling
    └── controller/ # Directory for defined logic
        ├── user.go # Defines the User input validation and file persistence logic
        ├── note.go # Defines the Note input validation and file persistence logic
    └── model/ # Directory for defined model
        ├── model.go # Defines the User and Note data model
    └── routes/ # Directory for data storage
        └── users.json # JSON file where all users are saved
        └── notes.json # JSON file where all notes are saved
```

## 💾 Data Persistence

All data are stored on a mysql database.

## 🤝 Contributing

Contributions are welcome! If you'd like to contribute, please:

1. Fork the repository.

2. Create a new branch (git checkout -b feature/your-feature-name).

3. Make your changes.

4. Commit your changes (git commit -m 'feat: Add new feature').

5. Push to the branch (git push origin feature/your-feature-name).

6. Open a Pull Request.
