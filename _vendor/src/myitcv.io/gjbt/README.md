### `gjbt`

`gjbt` is a simple (temporary) wrapper for GopherJS to run tests in Chrome as opposed to NodeJS. It should be considered
to be a direct replacement for `gopherjs test`.

Running your tests in Chrome has a number of benefits:

* it is almost certainly the VM in which your code will ultimately run
* you have full access to the DOM

Example:

```
$ gjbt myitcv.io/gjbt
ok      myitcv.io/gjbt  0.273s
PASS
```

### Requirements

(For now) A small wrapper is required around `TestMain` in each package to be tested. It is sufficient to copy
`init_test.go`, modifying the test package name, to your package.

### DOM Access

See [the tests in `myitcv.io/react`](https://github.com/myitcv/react/blob/master/a_elem_test.go) for examples of DOM
access.
