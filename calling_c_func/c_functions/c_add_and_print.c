#include <stdio.h>

int add_and_print(int a, int b) {
    int sum = a + b;
    printf("The sum of %d and %d is %d\n", a, b, sum);
    return sum;
}