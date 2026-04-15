// Echo: reads a character from input and prints it back
// Then prints the character + 1 (next ASCII character)

var ch = 0;
read(ch);
print(ch);

print(10); // newline

var limit = ch + 10;
while (ch < limit) {
    print(ch);
    ch = ch + 1;
}

print(10);
print(10);

var next = ch + 1;
print(next);
print(10); // newline
