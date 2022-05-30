/*
  configure file of key=value parse library.

  supported decode data to struct by reflact library,
  you can add "kv" tag in struct to ensure the correct keys.

  you can define a custom Marshaler and Unmarshaler interface
  to process data by yourself.

  supported value types:
    interface{}
    map[string]interface{}
    struct by tag kv
    map[string]<all of basic types>
  basic types:
    int, int8, int16, int32, int64
    uint, uint8, uint16, uint32, uint64
    float32, float64
    bool
    string
*/
package kvconf
