# Go SMTP Mail Sender

This project is a Go application that sends emergency alert emails using Google's SMTP server. It utilizes HTML templates for constructing the email body and can be run in a Docker container.

## Features

- Sends emails via Google SMTP
- Uses HTML templates for customizable email content
- Containerized with Docker for easy deployment
- Reads SMTP configuration from environment variables

## Prerequisites

- Go (1.21 or later)
- Docker
- A Google account with SMTP access (you may need to create an app-specific password if using 2-Step Verification).

## Project Structure
.
├── Dockerfile
├── .env
├── go.mod
├── go.sum
├── templates
|    └── sample_template.html
├── models
|    └── sample_model.go
├── controller
|    └── sample_controller.go
├── conusmer
|    └── sample_consumer.go


- **Dockerfile**: Defines the container for the Go application.
- **.env**: (Optional) Contains environment variables for SMTP configuration.
- **go.mod/go.sum**: Go modules for dependency management.
- **templates/**: Directory containing HTML email templates.
- **models/**: Directory containing struct of the mail.
- **controller/**: Directory containing the functional logic.
- **conusmer/**: Directory containing the kafka consumer logic.


## Setting Up

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Manikandan-Parasuraman/smtp-mail-sender.git
   cd smtp-mail-sender ```

2. **Create a** .env **file** (optional) in the project root with your SMTP configuration:

    ```bash
    SMTP_HOST=smtp.gmail.com
    SMTP_PORT=587
    SMTP_USER=your-email@gmail.com
    SMTP_PASS=your-app-specific-password
    ```

Alternatively, you can hardcode these values in the Dockerfile.

3. **Update the HTML template** in templates/emergency_alert.html to customize your email body.

## Building the Docker Image
Run the following command to build the Docker image:
    ``` bash
    docker build -t go-smtp-app .
    ```
## Running the Application
You can run the application using Docker. If you have created a `.env` file, use the following command:
    ```bash
    docker run --env-file .env go-smtp-app
    ```
If you have hardcoded the SMTP settings in the Dockerfile, run:
    ```bash
    docker run go-smtp-app
    ```
## Sending an Email
The application will send an email to the recipient specified in the `main.go` file. Make sure to replace `"recipient@example.com"` with the actual recipient's email address in the code.

## Troubleshooting
1. Ensure that you have allowed "Less secure apps" or have set up an app-specific password for your Google account.
2. Check for any errors in the console when running the Docker container for more details.

## License
This project is licensed under the GNU General Public License v3.0 (GPL). See the [LICENSE](https://www.gnu.org/licenses/gpl-3.0-standalone.html) file for more information.

## Author
Manikandan Parasuraman

