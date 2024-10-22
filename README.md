# Word Finder

A program written in Go to find the shortest word that contains a phrase you input, and it copies the found word to your clipboard.

## How it Works

1. **Input a Phrase:**
   The user enters a phrase in the CLI. If the phrase matches part of any word in a predefined wordset (`dict.txt`), the program will find and return the shortest word containing that phrase.

2. **Randomized Word Search:**
   The wordset is shuffled before every search, ensuring that results don’t always start from the same words (e.g., avoiding alphabetical bias).

3. **Cycle Through Words:**
   If a word has already been selected in a previous run, it is marked as "used" and will not be selected again. When pressing Enter without typing a new phrase, the program will reuse the previous input, allowing you to cycle through results.

4. **Clipboard Integration:**
   Once a word is found, it is copied to your clipboard for easy pasting.

5. **Exit Command:**
   Type `exit` to quit the program at any time.

## Example

1. **Input:**
   You enter the phrase `cat`.

   **Output:**
   The shortest word containing `cat` is printed and copied to your clipboard.

2. **Reusing Input:**
   If you press Enter without typing a phrase, it will reuse the previous phrase (`cat`) and continue searching for another word.

3. **Exiting the Program:**
   Type `exit` to stop the program.

### Installation

1. **Install Dependencies:**
   This program uses the `clipboard` package. Install it by running:

   ```bash
   go get github.com/atotto/clipboard
   ```

2. **Build the Program:**
   You can compile the program using:

   ```bash
   go build -o wordfinder
   ```

3. **Run the Program:**
   After compiling, run it with:

   ```bash
   ./wordfinder
   ```

Place your word list in a file called `dict.txt` in the same directory as the program, with each word on a new line. Its important that all the words in the `dict.txt` file is in uppercase. Or you can change the code to change the case.
