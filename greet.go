package greet

import "strings"

func Greet(who ...string) string {
  if who == nil {
    return "Hello World!"
  } else if len(who) == 1 {
    return "Hello " + who[0] + "!"
  } else if len(who) == 2 {
    return "Hello " + strings.Join(who, " and ") + "!"
  } else {
    allButLast := strings.Join(who[0:len(who)-1], ", ")
    return "Hello " + allButLast + ", and " + who[len(who)-1] + "!"
  }
}

