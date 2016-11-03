*input_list* it is a list with numbers.
Write a function which receive *input_list* and return a list *output* where

``` python
output[idx] = reduce(lambda x, y: x*y, input_list) / input_list

```
The algorithm shouldn't use the division.
