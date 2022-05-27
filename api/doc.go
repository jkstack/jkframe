/*
  http handler library, Context is the main struct of this library.


  default response data format like below

      {"code": 0, "payload": ...}

  parse request arguments:
    Context.X(type): means the argument by name is required,
                     if not passed the argument it will panic with MissingParam error,
                     if passed argument type is wrong it will panic with BadParam error
    Context.O(type): means the argument by name is optional,
                     if not passed the argument it will return default value,
                     if passed argument type is wrong it will return the zero value of its type

  response data:
    Context.OK: the api handler processed ok and response data like below
                 {"code": 0, "payload": ...}
    Context.ERR: the api handler processed failed and response data like below
                 {"code": code, "msg": error message}
    Context.Body: response data in byte slice
    Context.BodyFrom: response data from io.Reader

  http response:
    Context.HTTPNotFound: response http_code=404
    Context.HTTPServiceUnavailable: response http_code=503
    Context.HTTPTimeout: response http_code=408
    Context.HTTPConflict: response http_code=409
    Context.HTTPForbidden: response http_code=403

  common errors:
    Context.NotFound: panic with NotFound error
    Context.Timeout: panic with Timeout error
*/
package api
