/*
  extended yaml library and supported #include syntax.

  you can decode and encode with #include <dir> format in yaml file,
  if the dir is relative, it will change the dir to relative from this file.

  each file that is expanded will add n space char before each line,
  the n count is the number before #include in each #include line.

  you can nested 10 layers by #include tag.

  the Render function will expand each #include tag, and replace each relative
  path to absolute path, the returned string is the result of the final yaml data string.
*/
package yaml
