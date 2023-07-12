### Naive attempt at a compression algorithm

Motivation came from watching Silicon Valley and thinking "How hard could that be?"

This prgram takes a string of text and converts it down to its binary representation. Example: `1100010011101111`

We take that binary representation and convert it to an alternating count, meaning we count a string of ones, add that count to our compressed string, then we take the following zeros, count those, and add that count to our compressed string, etc.

In this example, we take our binary representation `1100010011101111`
And we can count `2` ones, `3` zeros, `1` one, ...

and we end up with this: `2312314`

We strore that along with the first bit (in this case, `1`) and we know that alternative digits represent a list of the next bit.

The bug: We can't have more than 10 of a single bit in a row :) - I'll fix that when I can come up with a way to delimit each number without adding to much space.
