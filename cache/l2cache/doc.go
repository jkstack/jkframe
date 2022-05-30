/*
  local cache of memory and disk library.

  when writing data size is more than limit size in New function,
  it will write all of the data to temporary file in New function seted directory.

  Read function is streaming by offset.
*/
package l2cache
