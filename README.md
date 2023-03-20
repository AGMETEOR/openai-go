# OpenAI Go SDK Example

This code defines a Go package called "openai" that provides a client for accessing the OpenAI API. The package contains an "Api" struct that serves as an entry point for accessing the different API endpoints.

## Installation

To use this example, you will first need to install the OpenAI Go SDK. You can do this using the following command:

```bash
go get github.com/AGMETEOR/openai-go/openai
```

## Usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/AGMETEOR/openai-go/openai"
)

func main() {
	c := openai.NewClient(nil)
	req := &openai.CompletionRequest{
		Model:       "text-davinci-003",
		Prompt:      "Say this is a test",
		MaxTokens:   7,
		Temperature: 0,
	}
	completion, _, err := c.Completions.CreateCompletion(context.Background(), req)
	if err != nil {
		return
	}

	fmt.Printf("COMPLETION %+v", completion)
}

```

## License
This example program is licensed under the MIT License. See the `LICENSE` file for more information.