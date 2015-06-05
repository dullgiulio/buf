# bytes.Buffer Alternative Experiments

## NoSQL

Joke implementation with goroutines writing only when data is read. Doesn't really implement the io.Writer semantics (with a buffer too small it just fails).

The resulting read might not get all the writes in order, either.

## Blocks

Implementation using a slice of fixed size arrays. Works as intended.
