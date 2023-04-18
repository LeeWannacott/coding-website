// ternary example
// 1. Write a function that takes a number and returns a string
// 2. If the number is divisible by 3, return "fizz"
// 3. If the number is divisible by 5, return "buzz"
// 4. If the number is divisible by 3 and 5, return "fizzbuzz"
// 5. If the number is not divisible by 3 or 5, return the number

// question: refactor the function below to use ternary operators
function fizzbuzz1(num){
  if (num % 3 === 0 && num % 5 === 0){
    return "fizzbuzz";
  } else if (num % 3 === 0){
    return "fizz";
  } else if (num % 5 === 0){
    return "buzz";
  } else {
    return num;
  }
}

// answer
function fizzbuzz(num) {
  return num % 3 === 0 && num % 5 === 0 ? "fizzbuzz" : num % 3 === 0 ? "fizz" : num % 5 === 0 ? "buzz" : num;
}

console.log("funciton")
