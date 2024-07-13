# wt

`wt` (webtext) is a command-line tool that fetches the content of a given URL, removes all HTML, CSS, and JavaScript, and outputs the raw human-readable text. This can be useful for various text processing tasks where only the main content of a webpage is needed.

## Features

- Fetches content from a given URL.
- Strips away HTML tags, CSS, and JavaScript.
- Outputs raw text for easy piping to other command-line tools.

## Installation

### Prerequisites

- Go 1.16+ installed on your machine.

### Build from Source

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/wt.git
    cd wt
    ```

2. Build the binary:

    ```sh
    go build -o wt main.go
    ```

3. Move the binary to a directory in your `$PATH`, for example:

    ```sh
    mv wt /usr/local/bin/
    ```

## Usage

```sh
wt <url>
```

### Examples

1. **Basic Usage**: Fetch the raw text from a webpage.

    ```sh
    wt https://example.com
    ```

2. **Piping to Another Command**: Count the number of words on a webpage.

    ```sh
    wt https://example.com | wc -w
    ```

3. **Saving Output to a File**: Save the raw text from a webpage to a file.

    ```sh
    wt https://example.com > output.txt
    ```

4. **Grep Through Webpage Content**: Search for a specific term in the raw text from a webpage.

    ```sh
    wt https://example.com | grep "specific term"
    ```

5. **Filter Out Specific Lines**: Filter out lines that contain a specific word.

    ```sh
    wt https://example.com | grep -v "word_to_filter_out"
    ```

6. **Chain Multiple Commands**: Chain multiple commands to process the webpage content.

    ```sh
    wt https://example.com | grep "interesting content" | sort | uniq
    ```

### Advanced Examples with AI Tooling

1. **Summarize Webpage Content using `sumy`**:

    ```sh
    wt https://example.com | sumy lex-rank --length=3 --text=- 
    ```

2. **Sentiment Analysis with `sentiment-cli`**:

    ```sh
    wt https://example.com | sentiment-cli
    ```

3. **Keyword Extraction with `rake-nltk`**:

    ```sh
    wt https://example.com | python3 -c "import sys; from rake_nltk import Rake; r = Rake(); r.extract_keywords_from_text(sys.stdin.read()); print(r.get_ranked_phrases())"
    ```

4. **Named Entity Recognition with `spacy`**:

    ```sh
    wt https://example.com | python3 -c "import sys; import spacy; nlp = spacy.load('en_core_web_sm'); doc = nlp(sys.stdin.read()); for ent in doc.ents: print(ent.text, ent.label_)"
    ```

5. **Topic Modeling with `gensim`**:

    ```sh
    wt https://example.com | python3 -c "import sys; from gensim import corpora, models; from gensim.parsing.preprocessing import preprocess_string; texts = [preprocess_string(sys.stdin.read())]; dictionary = corpora.Dictionary(texts); corpus = [dictionary.doc2bow(text) for text in texts]; ldamodel = models.LdaModel(corpus, num_topics=5, id2word = dictionary, passes=15); for idx, topic in ldamodel.print_topics(-1): print('Topic: {} \nWords: {}'.format(idx, topic))"
    ```

6. **Translate Webpage Content using `translate-shell`**:

    ```sh
    wt https://example.com | trans -b :fr
    ```

7. **Spell Check using `aspell`**:

    ```sh
    wt https://example.com | aspell list
    ```

8. **Word Frequency Analysis with `awk`**:

    ```sh
    wt https://example.com | tr -s '[:space:]' '\n' | sort | uniq -c | sort -nr | head -20
    ```

9. **Language Detection with `langid`**:

    ```sh
    wt https://example.com | python3 -c "import sys; import langid; print(langid.classify(sys.stdin.read()))"
    ```

10. **Summarize Webpage Content with `GPT-3`**:

    ```sh
    wt https://example.com | python3 -c "import sys; import openai; openai.api_key = 'your_api_key'; response = openai.Completion.create(engine='davinci', prompt=sys.stdin.read(), max_tokens=100); print(response.choices[0].text.strip())"
    ```

## License

This project is licensed under the GNU General Public License v2.0. See the [LICENSE](LICENSE) file for details.
```