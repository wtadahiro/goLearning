package main


  func main(){
    fizzbuzz()
    //println(fizzbuzz2(100))
  }

  func fizzbuzz(){
    for i :=1;  i < 100; i++{
      if i%15 == 0{
        print("fizzbuzz")
      }else if i%3 == 0{
        print("fizz")
      }else if i%5 == 0{
        print("buzz")
      }else{
        print(i)
      }
      print(" ")
    }
    println("finish")
  }



func fizzbuzz2(num int) string{
  for i := 1; i <= num; i++{

    fizz := i % 3
    buzz := i % 5

    switch{
      case fizz == 0 && buzz == 0:
        print("fizzbuzz")
      case fizz == 0:
        print("fizz")
      case buzz == 0:
        print("buzz")
      default :
        print(i)
    }

    print(" ")
  }
  return "finish"

}
