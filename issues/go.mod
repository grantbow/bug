module issues

replace github.com/driusan/bug/bugs => ../../../grantbow/bug/issues // fork

require github.com/driusan/bug/scm v0.0.0

replace github.com/driusan/bug/scm => ../../../grantbow/bug/scm // fork

require github.com/driusan/bug/bugapp v0.0.0

replace github.com/driusan/bug/bugapp => ../../../grantbow/bug/fitapp // fork

require github.com/ghodss/yaml v1.0.0

go 1.13
