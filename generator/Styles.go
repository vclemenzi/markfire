package generator

func GetHtmlStyle() string {
	return ` body {
  font-family: Arial, sans-serif;
  line-height: 1.6;
  margin: 0;
  padding: 20px;
  max-width: 100%;
}

h1, h2, h3, h4, h5, h6 {
  color: #333;
}

a {
  color: #007bff;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}

ul, ol {
  padding-left: 20px;
}

strong {
  font-weight: bold;
}

em {
  font-style: italic;
}

blockquote {
  margin: 0;
  padding-left: 20px;
  border-left: 2px solid #ccc;
  color: #666;
}

pre {
  background-color: #f4f4f4;
  border: 1px solid #ccc;
  padding: 10px;
  overflow-x: auto;
  margin: 20px 0;
}

code {
  font-family: Consolas, monospace;
  background-color: #f8f9fa;
  padding: 2px 5px;
  border-radius: 3px;
}

img {
  max-width: 100%;
  height: auto;
  display: block;
  margin: 20px auto;
}

@media (min-width: 768px) {
  body {
    max-width: 800px;
    margin: 0 auto;
  }
}`
}
