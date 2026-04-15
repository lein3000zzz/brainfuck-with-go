printString("Fibonacci:\n");

var a = 0;
var b = 1;
var i = 0;
var n = 10;
var next = 0;

while (i < n) {
    printNumber(a);
    printString(" ");
    next = a + b;
    a = b;
    b = next;
    i = i + 1;
}
printString("\n asd \n");