/*
  log library supported loglevel and logrotate.

  loglevel:
    Debug: 1‰ probability to write log with format [DEBUG]...
    Info: write info log with format [INFO]...
    Warning: write warning log with format [WARNING]...
    Error: write error log with stacktrace and format [ERROR]...

  logrotate:
    date: it will rename the log file to <name>_<yyyymmdd>.log in the next day.
    size: it will rename the log file to <name>.log.<n> when log size is more than limit size.
*/
package logging
