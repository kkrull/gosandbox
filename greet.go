package greet

func Greet(who ...string) string {
  if who == nil {
    return "Hello World!"
  } else {
    return "Hello " + who[0] + "!"
  }
}

