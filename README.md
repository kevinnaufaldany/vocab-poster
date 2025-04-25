# PinVocab

PinVocab is a web application designed to help users learn English vocabulary by creating a downloadable poster. Users can add English words along with their Indonesian translations, get word suggestions while typing, fetch random words, and download their vocabulary list in two formats: PNG or PDF. The application consists of a frontend and a backend, deployed on Railway.

- **Frontend URL**: [https://frontend-production-8243.up.railway.app](https://frontend-production-8243.up.railway.app)
- **Backend URL**: [https://vocab-poster-production.up.railway.app](https://vocab-poster-production.up.railway.app)

## Features
- Add English words and their Indonesian translations to a list.
- Get real-time word suggestions while typing an English word (powered by Datamuse API).
- Translate English words to Indonesian (powered by MyMemory API).
- Generate a random English word from various topics (e.g., animal, food, color).
- Download your vocabulary list as a poster in two formats: PNG or PDF (A4-sized).

## Usage Guide

### Prerequisites
- A modern web browser (e.g., Chrome, Firefox, or Safari).
- An internet connection to access the deployed application and external APIs.

### Step-by-Step Usage
Follow these steps to use PinVocab:

1. **Access the Application**  
   Open your web browser and navigate to the following URL:  
   [https://frontend-production-8243.up.railway.app](https://frontend-production-8243.up.railway.app)

2. **Add a Vocabulary Word**  
   - In the "InputVocabInggris" field, type an English word (e.g., "hap").
   - As you type, suggestions will appear below the input field (e.g., "happy", "happen"). Click on a suggestion to select it, or continue typing your word.
   - The "InputVocabIndo" field will automatically display the Indonesian translation (e.g., "senang"). If no translation is found, you can manually type the translation.
   - Click the "Add" button to add the word pair to your list.

3. **Get a Random Word**  
   - Click the "Get Random Word" button to fetch a random English word (e.g., "dog").
   - The word will automatically be added to your list along with its Indonesian translation (e.g., "anjing").

4. **View Your Vocabulary List**  
   - Your added words will appear in a list below the input fields.
   - Each entry shows the English word and its Indonesian translation.

5. **Download Your Poster**  
   - Once you have added at least one word to your list, you can download your vocabulary list in two formats:
     - **Download as PDF**: Click the "Download as PDF" button to generate a PDF file named `PinVocabPoster.pdf`. This will be an A4-sized poster ready for printing.
     - **Download as PNG**: Click the "Download as PNG" button to generate a PNG image named `PinVocabPoster.png`. This can be used for sharing or viewing on digital devices.
   - Open the downloaded file to view your vocabulary list in the chosen format.

## Project Structure
- **Frontend**: Located in the `frontend` folder, built with HTML, JavaScript, React, and Tailwind CSS (via CDN). It handles the user interface and communicates with the backend.
- **Backend**: Located in the `backend` folder, built with Go using the Gorilla Mux router. It provides API endpoints for word suggestions, translations, and random words.

## Backend Details
The backend (`main.go`) is a REST API that handles HTTP GET requests to the following endpoints:
- `/suggest?prefix=<prefix>`: Provides word suggestions based on a prefix using the Datamuse API.
- `/translate?word=<word>`: Translates an English word to Indonesian using the MyMemory API.
- `/random`: Returns a random English word based on predefined topics (e.g., animal, food, color) using the Datamuse API.

### External APIs
- **Datamuse API** (`https://api.datamuse.com`): Used for word suggestions and random word generation.
- **MyMemory API** (`https://api.mymemory.translated.net`): Used for translating English words to Indonesian.

### Notes on Backend
- **CORS**: The backend includes a CORS middleware to allow cross-origin requests from the frontend (`Access-Control-Allow-Origin: *`).
- **Port**: The server listens on a port specified by the `PORT` environment variable, with a fallback to `8080`.
- **Limitations**:
  - The backend relies on external APIs, which may have rate limits or downtime.
  - Errors from API requests are not detailed in the response, which can make debugging harder.
  - The application does not implement authentication, making it publicly accessible.

## Deployment
The application is deployed on Railway:
- **Frontend**: `https://frontend-production-8243.up.railway.app`
- **Backend**: `https://vocab-poster-production.up.railway.app`

### Deployment Notes
- The frontend is a static site served directly from the `frontend` folder.
- The backend is a Go application that requires a Railway service with the `PORT` environment variable set.
- Both services are linked to the GitHub repository (`kevinnaufaldy/vocab-poster`) and automatically redeploy on new commits.

## Running Locally
To run the application locally:

### Prerequisites
- Go (for the backend).
- A web browser (for the frontend).
- An internet connection (for API requests).

### Steps
1. **Clone the Repository**  
   ```bash
   git clone https://github.com/kevinnaufaldy/vocab-poster.git
   cd vocab-poster
   ```

2. **Run the Backend**  
   - Navigate to the `backend` folder:
     ```bash
     cd backend
     ```
   - Install dependencies:
     ```bash
     go mod tidy
     ```
   - Run the backend:
     ```bash
     go run main.go
     ```
   - The backend will run on `http://localhost:8080`.

3. **Run the Frontend**  
   - Open the `frontend/index.html` file in a web browser.
   - Update the backend URL in `index.html` to `http://localhost:8080`:
     ```javascript
     const response = await fetch(`http://localhost:8080/suggest?prefix=${value}`);
     const response = await fetch(`http://localhost:8080/translate?word=${word}`);
     const response = await fetch(`http://localhost:8080/random`);
     ```
   - Access the frontend by opening `index.html` in your browser (e.g., `file:///path/to/vocab-poster/frontend/index.html`).

## Contributing
If you want to contribute to PinVocab, feel free to fork this repository, make changes, and submit a pull request.

## License
This project is licensed under the MIT License.

## Notes
- The application uses external APIs (Datamuse and MyMemory), which may have rate limits. If you experience issues, consider switching to a local dictionary (as implemented in a later version of the backend).
- The vocabulary list is not saved after you close the browser. Download the PDF or PNG to keep your list.
- For personal use, you can share the frontend URL with friends or family, but be aware that the application is publicly accessible.