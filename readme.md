# wt - Web Text Extractor

**wt** is a command-line tool written in Go that extracts meaningful text content from web pages. It fetches a given URL, processes the HTML content, and outputs clean text by excluding unwanted elements like navigation menus, footers, headers, and advertisements. The tool is designed with sensible defaults but allows customization through command-line options.

## Features

- Extracts the main content from web pages using the [go-readability](https://github.com/go-shiori/go-readability) library.
- Filters out unwanted elements based on HTML tags and attributes.
- Configurable minimum text length threshold.
- Customizable include and exclude tags.
- Designed for easy integration with other tools via piping.
- Can be used to preprocess text for AI applications like Fabric AI.

## Installation

Ensure you have [Go](https://golang.org/dl/) installed (version 1.17 or later).

```bash
# Clone the repository
git clone https://github.com/copyleftdev/wt.git

# Navigate to the project directory
cd wt

# Download dependencies
go get ./...

# Build the application
go build -o wt ./cmd
```

## Basic Usage

```bash
./wt [options] <URL>
```

### Options:

- `-include`: Comma-separated list of HTML tags to include. Defaults to common content tags if not specified.
- `-exclude`: Comma-separated list of HTML tags to exclude. Defaults to non-content tags if not specified.
- `-minlength`: Minimum length of text blocks to include (default: 50).

## Examples

### Extracting Text with Default Settings

Extract text from a web page using default settings:

```bash
./wt https://example.com
```

### Specifying Tags to Include

Include only specific tags in the extraction:

```bash
./wt -include "p,h1,h2" https://example.com
```

### Excluding Specific Tags

Exclude certain tags from the extraction:

```bash
./wt -exclude "div,span" https://example.com
```

### Setting Minimum Text Length

Include only text blocks longer than a specified length:

```bash
./wt -minlength 20 https://example.com
```

### Combining Include, Exclude, and Minimum Length Options

Fine-tune the extraction by combining include, exclude, and minimum length options:

```bash
./wt -include "p,h1,h2,h3" -exclude "nav,footer,header" -minlength 10 https://example.com
```

### Piping Output to Another Program

Pipe the extracted text to another command-line tool or script:

```bash
./wt https://example.com | another_program
```

### Piping Output to Fabric AI

Use **wt** to preprocess text before feeding it into Fabric AI:

```bash
./wt https://example.com | fabric-ai process
```

## Use Cases

- **Content Extraction**: Extract articles or blog posts from web pages without ads or navigation links.
- **Text Analysis**: Preprocess web content for natural language processing tasks.
- **Data Mining**: Gather textual data from websites for analysis or machine learning.
- **Integration with AI Tools**: Pipe extracted text into AI applications like Fabric AI for further processing.
- **Web Scraping**: Automate the collection of textual content from multiple web pages.

## How It Works

1. **Fetches the Web Page**: Uses `http.Get` to retrieve the HTML content from the provided URL.
2. **Extracts Main Content**: Utilizes the `go-readability` library to parse and extract the main article content.
3. **Parses HTML Content**: Processes the extracted HTML to build a parse tree.
4. **Filters Content**:
   - **Tag Filtering**: Includes or excludes content based on specified HTML tags.
   - **Attribute Filtering**: Skips elements with unwanted `class`, `id`, or `role` attributes.
   - **Content Length Threshold**: Ignores text blocks shorter than a specified length (default is 50 characters).
5. **Outputs Clean Text**: Prints the extracted and filtered text to the console or pipes it to another program.

## Configuration

### Default Tags

The tool uses sensible defaults for include and exclude tags but allows customization:

- **Default Include Tags**:

  ```
  article, section, p, h1, h2, h3, h4, h5, h6, li, blockquote, pre, code, td, th
  ```

- **Default Exclude Tags**:

  ```
  script, style, nav, footer, header, aside, form, noscript, iframe, input, button, select, label, option, link, meta, figure, figcaption, dialog, menu, svg, canvas, video, audio, embed, object, and more.
  ```

### Minimum Text Length

You can set the `-minlength` option to control the minimum length of text blocks to include in the output. The default is `50` characters, but this can be adjusted based on your requirements.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests on the [GitHub repository](https://github.com/copyleftdev/wt).

## License

This project is licensed under the MIT License.

## Contact

For questions or support, please contact [copyleftdev](mailto:dj@codetestcode.io).
