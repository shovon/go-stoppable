# Stoppable

Package `stoppable` provides a helper struct to be extended into for the purpose of identifying when loops end.

## Usage

```
package main

import "github.com/shovon/go-stoppable"

type SomeStoppable struct {
	Stoppable
}

func (s *SomeStoppable) Loop() {
  s.Stoppable = stoppable.NewStoppable()
  go func() {
    <-s.OnStopped()
  }()
}

func main() {
  s := SomeStoppable{stoppable.NewStoppable}
  s.Loop()

  <-s.OnStopped()

  s.Stoppable = stoppable.NewStoppable

  s.Loop()

  <-s.OnStopped()
}
```
