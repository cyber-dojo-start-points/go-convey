
lambda { |stdout,stderr,status|
  output = stdout + stderr
  # go names the package after the module, so the reason it could not run is
  # the bracket it puts after that name rather than any path. [build failed]
  # is the compiler refusing a source file, [setup failed] is go/parser
  # refusing a _test.go file before the compiler ever sees it.
  return :amber if /\[build failed\]/.match(output)
  return :amber if /\[setup failed\]/.match(output)

  # A test file holding no test function still prints ok and exits zero, and
  # says so only in this warning. Without it a learner who renames their only
  # test function gets a green light.
  return :amber if /no tests to run/.match(output)

  # SkipConvey leaves a section unrun, and goconvey says so on its count line.
  # A suite that skipped its way to zero assertions proved nothing.
  return :amber if /\(one or more sections skipped\)/.match(output)

  return :red   if /FAIL/.match(output)
  return :green if /PASS/.match(output)
  return :amber
}
