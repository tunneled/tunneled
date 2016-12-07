package main

import(
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
  "time"
)

const(
  terminalWidth = 80
)

func colorize(text string, color string) string {
  colors := make(map[string]int)
  colors["purple"] = 35
  colors["cyan"] = 36
  colors["grey"] = 90
  colors["yellow"] = 33
  colors["red"] = 31

  colorValue := colors[color]
  return fmt.Sprintf("\033[%dm%s\033[0m", colorValue, text)
}

func clearScreen() {
  fmt.Printf("\033[H\033[2J")
}

func printDelimiter() {
  delimiter := colorize(strings.Repeat("*", terminalWidth), "grey")
  fmt.Printf("%s\n", delimiter)
}

func printNewline() {
  fmt.Println("")
}

func printRequestInformation(request *http.Request) {
  address := request.RemoteAddr
  method := request.Method
  uri := request.RequestURI

  leftSide := fmt.Sprintf("%s - %s %s", address, method, uri)
  leftWidth := len(leftSide)

  rightSide := time.Now().Format("2006-01-02 15:04:05 -0700")
  rightWidth := len(rightSide)

  spaces := strings.Repeat(" ", terminalWidth - (leftWidth + rightWidth))

  output := fmt.Sprintf("%s%s%s\n", leftSide, spaces, rightSide)
  fmt.Printf(colorize(output, "yellow"))
}

func printRequestHeaders(request *http.Request) {
  for key, values := range request.Header {
    for _, value := range values {
      output := fmt.Sprintf("%s: %s\n", key, value)
      fmt.Printf(colorize(output, "purple"))
    }
  }
}

func printRequestBody(request *http.Request) {
  var output string
  body, err := ioutil.ReadAll(request.Body)

  if err != nil {
    fmt.Printf("Could not print request body")
  } else {
    var bodyBuffer bytes.Buffer
    err := json.Indent(&bodyBuffer, []byte(body), "", "  ")

    if err != nil {
      output = fmt.Sprintf("%s\n", body)
    } else {
      formattedBody := bodyBuffer.String()
      output = fmt.Sprintf("%s\n", formattedBody)
    }

    fmt.Printf(colorize(output, "cyan"))
  }
}


func handler(writer http.ResponseWriter, request *http.Request) {
  printDelimiter()
  printRequestInformation(request)
  printNewline()
  printRequestHeaders(request)
  printNewline()
  printRequestBody(request)
  printDelimiter()
  printNewline()
  printNewline()
}

func main() {
  clearScreen()
  http.HandleFunc("/", handler)
  http.ListenAndServe("localhost:8000", nil)
}
