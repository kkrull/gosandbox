package greet_with_ginkgo

import "strings"

func Greet(who ...string) string {
  switch {
  case who == nil:
    return "Hello World!"
  case len(who) == 1:
    return "Hello " + who[0] + "!"
  case len(who) == 2:
    return "Hello " + strings.Join(who, " and ") + "!"
  default:
    allButLast := strings.Join(who[0:len(who)-1], ", ")
    return "Hello " + allButLast + ", and " + who[len(who)-1] + "!"
  }
}
