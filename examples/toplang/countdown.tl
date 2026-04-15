// Countdown from 5 to 1, printing each digit
// ASCII '0' = 48, so digit N = 48 + N

var i = 5;
var zero = 48;
var nl = 10;

while (i > 0) {
    var ch = i + zero;
    print(ch);
    print(nl);
    i = i - 1;
}
